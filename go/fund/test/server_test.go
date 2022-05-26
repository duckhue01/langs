package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestMain(m *testing.M) {
	Routes()
}

func TestEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/json", nil)
	if err != nil {
		t.Fatal("\tShould be able to create a request.", err)
	}

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)
	t.Log(rw)
}
