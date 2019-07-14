package main

import (
	"errors"
	"fmt"
	"log"
)

// 自定义错误结构，可以包含低层级的错误信息，还可以自定义保存其他有用信息
type FirstError struct {
	Raw string
	Err error
}

type SecondError struct {
	Raw string
	Err error
}

func (f *FirstError) Error() string {
	return f.Raw + " - " + f.Err.Error()
}

func (s *SecondError) Error() string {
	return s.Raw + " - " + s.Err.Error()
}

func main() {
	err := throwErr(2)
	// 类型断言
	if e, ok := err.(*FirstError); ok {
		fmt.Println(e)
	} else {
		fmt.Println("It is not firstError.")
	}

	// 类型判断（type-switch）
	switch e := err.(type) {
	case *FirstError:
		log.Println(e)
	case *SecondError:
		log.Println(e)
	}

	// 通过 fmt.Errorf 创建错误对象
	// 查看源码其实就是通过 Sprinf 拼装字符串
	// 再通过 errors.New 获得一个错误类型
	err = fmt.Errorf("error")
	log.Println(err)
}

func throwErr(i int) error {
	switch i {
	case 1:
		return &FirstError{
			Raw: "FirstError",
			Err: errors.New("Lower Error Message."),
		}
	case 2:
		return &SecondError{
			Raw: "SecondError",
			Err: errors.New("Lower Error Message."),
		}
	}
	return nil
}
