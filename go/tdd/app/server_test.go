package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server = &PlayerServer{
	store: &InMemoryStore{},
}

func TestGetPlayer(t *testing.T) {

	testCases := []struct {
		Want  string
		Name  string
		Input struct {
			Name string
		}
	}{
		{
			Want: "duckhue01",
			Name: "get player named khue",
			Input: struct{ Name string }{
				Name: "duckhue01",
			},
		},
	}


	// for _, v := range testCases {
	// 	t.Run(v.Name, func(t *testing.T) {
	// 		res := httptest.NewRecorder()
	// 		req, _ := createGetMethod(fmt.Sprintf("/player/%s", v.Input.Name))
	// 		server.ServeHTTP(res, req)

	// 		got := res.Body.String()
	// 		want := v.Want

	// 		if want != got {
	// 			t.Errorf("got %q want %q", got, want)
	// 		}
	// 	})
	// }




	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res := httptest.NewRecorder()
			req, _ := createGetMethod(fmt.Sprintf("/player/%s", v.Input.Name))
			server.ServeHTTP(res, req)

			got := res.Body.String()
			want := v.Want

			if want != got {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}

func createGetMethod(url string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, url, nil)
}

// func createPostMethod(url string, body )
