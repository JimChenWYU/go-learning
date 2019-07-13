package handlers_test

import (
	"encoding/json"
	"go-learning/test-tool/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	handlers.Routes()
}

func TestSendJson(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.", err)
		}

		t.Log("\tShould be able to create a request.", checkMark)

		// 创建一个 http.ResponseRecorder 值
		rw := httptest.NewRecorder()

		// 调用服务默认的多路选择器（mux）的 ServeHTTP 方法，调用这个方法是模拟外部客户端对服务端点的请求
		http.DefaultServeMux.ServeHTTP(rw, req)
		if rw.Code != http.StatusOK {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", err)
		}
		t.Log("\tShould decode the response.", checkMark)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name.", checkMark)
		} else {
			t.Error("\tShould have a Name.", ballotX, u.Name)
		}

		if u.Email == "Bill@ardanstudios.com" {
			t.Log("\tShould have an Email.", checkMark)
		} else {
			t.Error("\tShould have a Email.", ballotX, u.Email)
		}
	}
}
