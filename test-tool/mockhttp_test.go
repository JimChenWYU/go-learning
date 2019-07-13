package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var JSON = `
{
  "args": {}, 
  "data": "", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3", 
    "Accept-Encoding": "gzip, deflate, br", 
    "Accept-Language": "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7", 
    "Dnt": "1", 
    "Host": "httpbin.org", 
    "Purpose": "prefetch", 
    "Upgrade-Insecure-Requests": "1", 
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
  }, 
  "json": null, 
  "method": "GET", 
  "origin": "110.64.73.232, 110.64.73.232", 
  "url": "https://httpbin.org/anything"
}
`

func TestGetJson3(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	defer server.Close()

	t.Log("Begin testing...")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\".", server.URL, statusCode)
		{
			r, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer r.Body.Close()

			if r.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, r.StatusCode)
			}

			result, _ := ioutil.ReadAll(r.Body)
			t.Log(bytes.NewBuffer(result).String())
		}
	}
}

// 声明的 mockServer 函数，返回一个指向 httptest.Server 类型的指针。
// 这个 httptest.Server 的值是整个模仿（mock）服务关键
// http.HandlerFunc 类型是一个适配器，允许常规函数作为 HTTP 的处理函数使用。如果函数 f 具有合适的签名，
// http.HandlerFunc(f) 就是一个处理 HTTP 请求的 Handler 对象，内部通过调用 f 处理请求
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, JSON)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}
