package test

import (
	"net/http"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestDowload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	res, err := http.Get(url)

	t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
	if err != nil {
		t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
	}
	defer res.Body.Close()
	t.Log("\t\tShould be able to make the Get call.",
		checkMark)

	if res.StatusCode == statusCode {
		t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
	} else {
		t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, res.StatusCode)
	}

}
