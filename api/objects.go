package api

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Handler struct {
	path           string
	handler        func(w http.ResponseWriter, r UrlRequest)
	method         string
	pathSliced     []string
	pathVars       []string
	pathWithoutVars []string
	pathVarIndexes pathParamsIndexes
}

type Handlers []Handler

// Add's a handler struct to a slice of handlers spliting the path in the character '/'
func (h *Handlers) Add(path, method string,fn func(w http.ResponseWriter, r UrlRequest)) {

	if strings.Index(path, "/") == 0 {
		path = path[1:]
	}

	if len(path) > 1 && strings.LastIndex(path, "/") == (len(path) - 1) {
		path = path[0:len(path)-1]
	}

	if path == "" {
		*h = append(*h, Handler{
			path: path,
			handler: fn,
			method: method,
		})
	}

	// Slices the path string when it find's the char '/'
	pathSliced := strings.Split(path, "/")

	// Returns 3 slices each of of them containing: a slice containing the path parameters excluding the
	// variable placeholders; the index of each variable placeholder in the pathSliced slice; a slice
	// containing the name of each variable found in the pathSliced slice ( the value between the characters '{}' )
	pathWithoutVars, pathVarIndexes, pathVars := splitVarsFromStaticPathParameters(pathSliced)


	*h = append(*h, Handler{
		path,
		fn,
		method,
		pathSliced,
		pathVars,
		pathWithoutVars,
		pathVarIndexes,
	})
}

type VarsHandler struct {
	H    *Handler
	Vars map[string]string
}

type UrlRequest struct {
	R    *http.Request
	Vars map[string]string
}

type Error struct {
	msg    string
	status int
}

func (e Error) Error() string {
	return fmt.Sprintf("%s", e.msg)
}

type Data struct {
	status  int
	message string
	data    interface{}
}

type pathParamsIndexes []int

// Returns the index of the slice position of a value.
// This function will return the value of -1 if the slice of int's doesn't contain the value that we used to call it
func (e pathParamsIndexes) Contains(val int) int {
	for i, v := range e {
		if v == val {
			return i
		}
	}
	return -1
}

// GetByMethodAndType returns the handler that matches the url and the http request method being sent by the user,
// the path string may contain variables identified by a start character and an end character,
// this characters are '{' and '}' respectively.
// This method returns an 'Error' object in case of not finding any handler that matches our path url or the an
// VarsHandler object that contains the handler that matches the url and the http request method in the property 'H
// ' and a map of type key->string and value->string containing the variables that the matching url should contain.
// For instance if we call this method with a path value '/28/contacts/1' and we have a Handler whose path value is
// '/{id}/contacts/{contactId} this method will return a VarsHandler object with property 'H' having the matching
// hanlder object and the property 'Vars' a map with {"id": 28, "contactId": 1}
func (h Handlers) GetByMethodAndType(path, method string) (*VarsHandler, *Error) {

	if strings.Index(path, "/") == 0 {
		path = path[1:]
	}

	if len(path) > 1 && strings.LastIndex(path, "/") == (len(path) - 1) {
		path = path[0:len(path)-1]
	}

	// split's the path string in a slice ,
	// if the path string contains a value of 'users/list' the slicedPath should contain a slice of {"users", "list"}
	slicedPath := strings.Split(path, "/")

	// Loops through all Handler objects
	for _, handler := range h {

		// If the handler property path value is equal to the value of the path arg and the handler method value is
		// equal to the method arg,
		// means that we already found the handler that we were looking for and we should return that handler inside
		// of a VarsHandler struct.
		if handler.path == path && handler.method == method {
			return &VarsHandler{H: &handler}, nil
		}

		// If the value of path property inside our handler doesn't contain a single character equal to '{' means
		// that this handler contains a simple and straight forward path and we should skip the next check
		if len(handler.pathVars) == 0 {
			continue
		}

		// Check's the handler path,
		// splitting the handler path into a slice using the substring '/' and comparing it's length with the length
		// of the slicedPath arg. If the length is the same starts by removing the variables from the sliced
		// handler Path parameters and the value off the matching slice position of the slicedPath  slice.
		// Then it will extract the variables values identified in the handler path string from the slicedPath slice
		// and returns the actualPath without the variable values that matched the position of the handler variable
		// placeholders and a map containing a set of key/value pairs of the variable names and values extracted
		// from the originalPathString and from the handler path string.
		actualCleanPaths, handlerCleanPaths, vars, ok := verifyPathMatchingAndReturnPathParameterVariables(handler,
			slicedPath)

		// The verifyPathMatchingAndReturnPathParameterVariables should return a non ok value if the slicedPath len
		// doesn't match to len of the slice that will be generated after splitting the handler path value when the
		// char '/' is found in it, if the len's doen't match we should proceed to testing the next handler.
		if !ok {
			continue
		}

		// Compares both slice objects and if they are deeply equal and the method is the same we should return this
		// handler.
		if compareSlices(actualCleanPaths, handlerCleanPaths) && handler.method == method {
			return &VarsHandler{H: &handler, Vars: vars}, nil
		}
	}

	return nil, &Error{ErrNotFound, http.StatusNotFound}
}

func compareSlices(actual []string, expected []string) bool {
	if len(actual) == len(expected) && len(actual) == 0 {
		return true
	}

	return reflect.DeepEqual(actual, expected)
}


// Verifies if the number of path parameters present in the request path matches with the number of path parameters
// of an handler
func verifyPathMatchingAndReturnPathParameterVariables(handler Handler, originalSlicedPath []string) ([]string,
	[]string,
	map[string]string, bool) {

	// If the length of the original originalSlicedPath doesn't match with the length of handler.pathSliced means that
	// this handler is not the handler that we are looking for and we should return false in the last return value
	if len(originalSlicedPath) != len(handler.pathSliced) {
		return nil, nil, nil, false
	}

	originalPathWithoutValuesThatMatchVarIndexes := []string{}
	vars := map[string]string{}

	// Loops throught all path parameters of the original url
	for i, path := range originalSlicedPath {

		// Checks if the handler contains a variable in the same index that this path parameter is located,
		// and if it don't add's the path parameter to the originalPathWithoutValuesThatMatchVarIndexes slice.
		index := handler.pathVarIndexes.Contains(i)
		if index == -1 {
			originalPathWithoutValuesThatMatchVarIndexes = append(originalPathWithoutValuesThatMatchVarIndexes, path)
			continue
		}
		// If the path parameter index is equal to the index value of variable of the handler path property ,
		// stores the variable value in the vars map obtaining the variable name from the handler.pathVars[index] slice
		// using the previous obtained index.
		vars[handler.pathVars[index]] = path
	}

	// returns the originalPathWithoutValuesThatMatchVarIndexes slice, the handler.pathWithoutVars slice ,
	// the variables map containing the variables names and values and a status of true.
	return originalPathWithoutValuesThatMatchVarIndexes, handler.pathWithoutVars, vars, true
}

// This method removes variables from handler path strings returning the handler path parameters that don't contain
// variables, the variables removed and the slice indexes were these values were located.
func splitVarsFromStaticPathParameters(splitedPath []string) (handlerPathWithoutVars []string,
	handlerPathVarIndexes pathParamsIndexes, handlerPathVarNames []string) {

	// Loops through all slices of the resulting string split executed on the handlerPath content
	for i, path := range splitedPath {
		if len(path) == 0 {
			continue
		}

		// If the path slice start's and end's with '{' and '}' respectively,
		// means that this path parameter is a variable and we should store the index of the path parameter in a
		// slice and store the string that is between the same characters as a var name/id.
		if path[len(path)-1:] == "}" && path[0:1] == "{" {
			handlerPathVarIndexes = append(handlerPathVarIndexes, i)
			handlerPathVarNames = append(handlerPathVarNames, path[1:len(path)-1])
			continue
		}
		// If the path parameter is not a variable we shoudl append it to a slice that will contains all path
		// parameters that aren't variables.
		handlerPathWithoutVars = append(handlerPathWithoutVars, path)
	}
	return
}
