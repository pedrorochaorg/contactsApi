package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlers_Add(t *testing.T) {

	t.Run("adding an handler with an empty path" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("", http.MethodGet, handlerFunc)

		assert.Equal(t, "", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathSliced, "PathSliced value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathWithoutVars, "PathWithoutvars value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes(nil), handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})


	t.Run("adding an handler with a path with '/'" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("/", http.MethodGet, handlerFunc)

		assert.Equal(t, "", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathSliced, "PathSliced value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathWithoutVars, "PathWithoutvars value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes(nil), handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})

	t.Run("adding an handler with a path with '/contacts'" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("/contacts", http.MethodGet, handlerFunc)

		assert.Equal(t, "contacts", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string{"contacts"}, handlers[0].pathSliced, "PathSliced value should be equal")
		assert.Equal(t, []string{"contacts"}, handlers[0].pathWithoutVars, "PathWithoutvars value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes(nil), handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})

	t.Run("adding an handler with a path with 'contacts'" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("contacts", http.MethodGet, handlerFunc)

		assert.Equal(t, "contacts", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string{"contacts"}, handlers[0].pathSliced, "PathSliced value should be equal")
		assert.Equal(t, []string{"contacts"}, handlers[0].pathWithoutVars, "PathWithoutvars value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes(nil), handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})

	t.Run("adding an handler with a path with 'contacts/'" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("contacts/", http.MethodGet, handlerFunc)

		assert.Equal(t, "contacts", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string{"contacts"}, handlers[0].pathSliced, "PathSliced value should be equal")
		assert.Equal(t, []string{"contacts"}, handlers[0].pathWithoutVars, "PathWithoutvars value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes(nil), handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})


	t.Run("adding an handler with a path with '{id}'" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("{id}", http.MethodGet, handlerFunc)

		assert.Equal(t, "{id}", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string{"{id}"}, handlers[0].pathSliced, "PathSliced value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathWithoutVars, "PathWithoutvars value should be equal")
		assert.Equal(t, []string{"id"}, handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes{0}, handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})

	t.Run("adding an handler with a path with '/{id}'" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("/{id}", http.MethodGet, handlerFunc)

		assert.Equal(t, "{id}", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string{"{id}"}, handlers[0].pathSliced, "PathSliced value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathWithoutVars, "PathWithoutvars value should be equal")
		assert.Equal(t, []string{"id"}, handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes{0}, handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})

	t.Run("adding an handler with a path with '/{id}/'" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("/{id}/", http.MethodGet, handlerFunc)

		assert.Equal(t, "{id}", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string{"{id}"}, handlers[0].pathSliced, "PathSliced value should be equal")
		assert.Equal(t, []string(nil), handlers[0].pathWithoutVars, "PathWithoutvars value should be equal")
		assert.Equal(t, []string{"id"}, handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes{0}, handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})

	t.Run("adding an handler with a path with '/{id}/contacts/'" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("/{id}/contacts/", http.MethodGet, handlerFunc)

		assert.Equal(t, "{id}/contacts", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string{"{id}", "contacts"}, handlers[0].pathSliced, "PathSliced value should be equal")
		assert.Equal(t, []string{"contacts"}, handlers[0].pathWithoutVars, "PathWithoutvars value should be equal")
		assert.Equal(t, []string{"id"}, handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes{0}, handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})

	t.Run("adding an handler with a path with '/sample/{id}/contacts/'" , func(t *testing.T) {
		handlers := Handlers{}


		handlerFunc := func(w http.ResponseWriter, r UrlRequest) {
		}

		handlers.Add("/sample/{id}/contacts/", http.MethodGet, handlerFunc)

		assert.Equal(t, "sample/{id}/contacts", handlers[0].path, "Path should be empty")
		assert.Equal(t, http.MethodGet, handlers[0].method, "Method's should be equal")
		assert.Equal(t, []string{"sample", "{id}", "contacts"}, handlers[0].pathSliced,
		"PathSliced value should be equal")
		assert.Equal(t, []string{"sample", "contacts"}, handlers[0].pathWithoutVars,
		"PathWithoutvars value should be equal")
		assert.Equal(t, []string{"id"}, handlers[0].pathVars, "PathVars value should be equal")
		assert.Equal(t, pathParamsIndexes{1}, handlers[0].pathVarIndexes, "PathVarIndexs value should be equal")

	})
}


func TestHandlers_GetByMethodAndType(t *testing.T) {
	handlers := Handlers{}

	handlers.Add("/sample/{id}/contacts/", http.MethodGet, func(w http.ResponseWriter, r UrlRequest) {

	})

	handlers.Add("", http.MethodGet, func(w http.ResponseWriter, r UrlRequest) {

	})
	handlers.Add("", http.MethodPost, func(w http.ResponseWriter, r UrlRequest) {

	})
	handlers.Add("/something", http.MethodGet, func(w http.ResponseWriter, r UrlRequest) {

	})

	t.Run("valid path with vars should return a valid handler" , func(t *testing.T) {
		handler, err := handlers.GetByMethodAndType("/sample/stuff/contacts", http.MethodGet)

		assert.NotNil(t, handler, "handler should not be nil")
		assert.Nil(t, err, "error should be nil")
	})

	t.Run("valid path with vars should return a valid handler" , func(t *testing.T) {
		handler, err := handlers.GetByMethodAndType("/sample/stuff/contacts/", http.MethodGet)

		assert.NotNil(t, handler, "handler should not be nil")
		assert.Nil(t, err, "error should be nil")
	})

	t.Run("valid path without vars should return a valid handler" , func(t *testing.T) {
		handler, err := handlers.GetByMethodAndType("/", http.MethodGet)

		assert.NotNil(t, handler, "handler should not be nil")
		assert.Nil(t, err, "error should be nil")
	})

	t.Run("valid path without vars and different method should return a different handler" , func(t *testing.T) {
		handler, err := handlers.GetByMethodAndType("/", http.MethodPost)

		assert.NotNil(t, handler, "handler should not be nil")
		assert.Nil(t, err, "error should be nil")

		getHandler, err := handlers.GetByMethodAndType("/", http.MethodGet)

		assert.NotNil(t, getHandler, "handler should not be nil")
		assert.Nil(t, err, "error should be nil")

		assert.NotEqual(t, handler, getHandler, "handlers shouldn't match")
	})

	t.Run("invalid path should return an error" , func(t *testing.T) {
		handler, err := handlers.GetByMethodAndType("/test", http.MethodPost)

		assert.NotNil(t, err, "err should not be nil")
		assert.Nil(t, handler, "handler should be nil")

		assert.Equal(t, err.status, http.StatusNotFound)
		assert.Equal(t, err.msg, ErrNotFound)

		assert.Equal(t, err.Error(), ErrNotFound)

	})
}