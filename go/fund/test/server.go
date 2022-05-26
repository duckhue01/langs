package test

import (
	"encoding/json"
	"net/http"
)

func main() {


	http.NewServeMux()

}

func Routes() {
	http.HandleFunc("/json", SendJSON)

}

// SendJSON returns a simple JSON document.
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanstudios.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
