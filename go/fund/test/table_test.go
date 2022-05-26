package test

import (
	"net/http"
	"testing"
)

func TestDownloadTable(t *testing.T) {
	var tcs = []struct {
		url        string
		statusCode int
	}{
		{
			url:        "http://www.goinggo.net/feeds/posts/default?alt=rss",
			statusCode: http.StatusOK,
		},
		{
			url:        "http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			statusCode: http.StatusNotFound,
		},
	}

	t.Log("given the need to downloading difference content")

	for _, tc := range tcs {
		t.Logf("checking [%s] for status code [%d]", tc.url, tc.statusCode)
		res, err := http.Get(tc.url)
		if err != nil {
			t.Fatal("should be able to get the url", err)
		}
		t.Log("should be able to get url")

		defer res.Body.Close()

		t.Logf("should have a [%d] status code", tc.statusCode)
		if res.StatusCode != tc.statusCode {
			t.Errorf("got %d, but expected %d", res.StatusCode, tc.statusCode)

		}

	}

}
