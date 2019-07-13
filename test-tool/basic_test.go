package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// Go 语言会以 _test.go 结尾的文件是测试文件，testing 包提供了从测试框架到报告测试的输出和状态的各种测试功能的支持
// 测试函数必须是公开的函数，以 Test 为开头，签名中必须接收一个指向 testing 的指针，而且不返回任何值。
func TestGetJson(t *testing.T) {
	url := "https://httpbin.org/anything"
	statusCode := 200

	// t.Fatal 方法不但报告这个单元测试已经失败，而且还会向测试输出一些消息，而后立刻停止这个测试函数的执行。
	t.Log("Send a request")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\".", url, statusCode)
		{
			r, err := http.Get(url)
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

			var m map[string]interface{}

			err = json.NewDecoder(r.Body).Decode(&m)
			if err != nil {
				t.Fatal("\t\tShould receive a json string from a response.", ballotX, err)
			}

			_, ok := m["headers"]
			if !ok {
				t.Errorf("\t\tShould receive headers from map.")
			}
			t.Log(m)
		}
	}
}
