package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// admin 类型主动实现接口
//func (a *admin) notify()  {
//	fmt.Printf("Sending admin email to %s<%s>\n", u.name, u.email)
//}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 直接访问内部类型的方法
	ad.user.notify()

	// 内部类型方法被提升
	// 需要注意的是，admin 类型也可以主动实现接口，这样内部类型的方法就不会被提升
	ad.notify()

	// admin 类型没有实现 notifier 接口，但其内部类型嵌入的 user 类型实现了接口
	// 因为内部类型的提升，使得内部类型实现的接口会自动提升到外部类型，也就是说外部类型也实现这个接口
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
