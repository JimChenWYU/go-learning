package main

import (
	"fmt"
)

// 切片比数组可以更加方便管理数据集合，它是围绕动态数组的概念构建的，可以按需自动增长和缩小。
// 切片是对象，是对底层数组的抽象，并提供了相应的操作方法。
// 切片有三个字段的数据结构，分别是指向底层数组的指针，切片访问的元素个数（长度），切片允许增长到的元素个数（容量）。
func main() {
	// 创建一个字符串切片
	// 其长度和容量都是5个元素
	slice1 := make([]string, 5)

	// 创建一个字符串切片
	// 其长度为3，容量为5，注意：创建切片时容量不能小于长度
	slice2 := make([]string, 3, 5)

	// 创建一个 string 类型切片
	// 其长度和容量都是5个元素
	slice3 := []string{"red", "blue", "white", "black", "green"}

	// 创建nil切片
	var slice4 []string

	// 创建空切片
	slice5 := make([]string, 0)
	slice6 := []string{}

	// 切片使用跟数组使用一样
	slice7 := []string{"red", "blue", "white", "black", "green"}
	slice7[0] = "yellow"

	// 使用切片来创建新的切片
	// 其长度为2，容量为4
	// 对底层数组容量是 k 的切片 slice[i:j] 来说，长度为j-i，容量为k-i
	// newSlice7 与 slice7 共享底层数组，所以任意一方修改都会影响另外一方
	newSlice7 := slice7[1:3]

	// append
	// 合并切片API
	slice8 := []string{"red", "blue", "white", "black", "green"}
	newSlice8 := slice8[1:3]

	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为60
	newSlice8 = append(newSlice8, "yellow") // newSlice8=["blue", "white", "yellow"], slice8=["red", "blue", "white", "yellow", "green"]

	// len() 获取切片长度
	// cap() 获取切片容量
	slice9 := make([]string, 5, 10)
	slice9_len := len(slice9)
	slice9_cap := cap(slice9)

	fmt.Printf("slice1=%s | slice2=%s | slice3=%s | slice4=%s | slice5=%s | slice6=%s | slice7=%s | slice8=%s \n",
		slice1, slice2, slice3, slice4, slice5, slice6, slice7, slice8)
	fmt.Printf("newSlice7=%s | newSlice8=%s \n", newSlice7, newSlice8)
	fmt.Printf("slice9, len=%d, cap=%d \n", slice9_len, slice9_cap)
}
