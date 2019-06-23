package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

// 使用值接收者实现一个方法
// 值接收者，调用方法会用这个值的一个副本来执行
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u user) changeName(name string) {
	u.name = name
}

// 使用指针接收者实现一个方法
// 指针接收者，调用方法时会用这个值的指针来执行，方法就会共享调用方法时接收者所指向的值
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.changeName("Jim")
	bill.notify()

	lisa := user{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()
}
