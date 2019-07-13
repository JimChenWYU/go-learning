package main

import (
	"net/http"
	"testing"
)

// 表组测试就是一组基础单元测试，通过多个测试用例进行测试
func TestGetJson2(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			url:        "https://httpbin.org/status/200",
			statusCode: http.StatusOK,
		},
		{
			url:        "https://httpbin.org/status/200",
			statusCode: http.StatusForbidden,
		},
		{
			url:        "https://httpbin.org/status/404",
			statusCode: http.StatusNotFound,
		},
	}

	t.Log("Begin testing...")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				r, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.", ballotX, err)
				}
				t.Log("\t\tShould be able to Get the url", checkMark)

				defer r.Body.Close()

				if r.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status. %v", u.statusCode, ballotX)
				}
			}
		}
	}
}
