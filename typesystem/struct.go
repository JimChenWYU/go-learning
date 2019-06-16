package main

import (
	"fmt"
)

// 声明一个结构类型
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

// 使用其他类型结构类型作为字段声明
type admin struct {
	person user
	level  string
}

func main() {
	// 声明 user 类型变量
	var bill user

	// 声明并初始化
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	// 不使用字段名声明并初始化（不推荐）
	// 通常一行，顺序很重要，要按字段定义顺序
	jim := user{"Jim", "jim@email.com", 123, true}

	manager := admin{
		person: user{
			name:       "Jim",
			email:      "jim@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}

	fmt.Println(bill, lisa, jim, manager)
}
