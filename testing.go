package contactsApi

import "testing"

func AssertStringEquals(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("strings don't match got %s want %s", got, want)
	}
}
