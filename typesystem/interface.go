package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	// http.Get 返回的是 http.Response 类型的指针，包含了一个名为 Body 的字段
	// Body 是 io.ReadCloser 接口类型的值
	r, err := http.Get("https://httpbin.org/anything")
	if err != nil {
		fmt.Println(err)
		return
	}

	// os.Stdout 实现了 io.Writer 接口
	// r.Body 实现了 io.Reader 接口
	// 故可以直接传递给 io.Copy 处理
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}

	define()
}

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", a.name, a.email)
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func define() {
	u := user{"Bill", "bill@email.com"}
	a := admin{"Lisa", "lisa@email.com"}

	// 因为是使用指针接收者实现接口，所以传入时需要传指针
	sendNotification(&u)
	sendNotification(&a)
}

func sendNotification(n notifier) {
	n.notify()
}
