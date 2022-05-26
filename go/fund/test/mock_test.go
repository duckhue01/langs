package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const mock = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

// mockServer returns a pointer to a server to handle the get call.
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, mock)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestHTTPServer(t *testing.T) {
	server := mockServer()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("Should be able to make the Get call.", err)
	}
	res.Body.Close()

}
