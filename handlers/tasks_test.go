package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	testBody := []byte(`{"name:"task1", "done":"no"}`)

	testr, err := http.NewRequest("POST", "/task", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatalf("Create failed with err: %v", err)
	}
	testr.Header.Set("Content-Type", "application/json; charset=utf-8")

	wr := httptest.NewRecorder()

	http.HandlerFunc(Create).ServeHTTP(wr, testr)

	for i := range TasksList {
		if TasksList[i].Name == "task1" {
			return
		}
	}
	err = errors.New("Could not find task in list")
	t.Fatalf("List of tasks: %v; /n err: %v", TasksList, err)
}
