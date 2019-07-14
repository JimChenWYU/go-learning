package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// 以下就是示例代码，会出现在文档对应函数的实示例代码下
// Output: 这个标记用来在文档中标记出示例函数运行后期望的输出

func ExampleSendJson() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(rw, r)
	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR: ", err)
	}

	fmt.Println(u)

	// Output:
	// {Bill Bill@email.com}
}
