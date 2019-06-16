package main

import (
	"fmt"
)

func main() {

	fmt.Println("普通数组：")
	array()

	fmt.Println("多维数组：")
	multi()

	a1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("按值转递：")
	// 复制数组，会使用更多内存
	foo(a1)

	fmt.Println("指针转递：")
	// 复制数组的指针，更有效使用内存，性能也很好，不过需要意识到传递的是数组指针，在函数内部改变数组指向的值，会改变共享内存。
	bar(&a1)
}

func array() {
	// 声明数组，长度为5
	var a1 [5]int

	// 初始化赋值
	a1[0] = 1

	// 声明和初始化5个元素的数组
	a2 := [5]int{1, 2, 3, 4, 5}

	// 声明和初始化3个元素的数组，数组容量由初始化的值的数量决定
	a3 := [...]int{1, 2, 3}

	// 声明一个有5个元素的数组
	// 用具体值初始化索引为1和2的元素
	// 其余元素保持零值
	a4 := [5]int{1: 10, 2: 20}

	// 用整型指针初始化索引为0和1的数组元素
	a5 := [5]*int{0: new(int), 1: new(int)}

	// 为索引为0和1的元素赋值
	*a5[0] = 10
	*a5[1] = 20

	// 把a2数组赋值给a6数组，只有数组长度相同，每个元素类型相同才可以实现
	var a6 [5]int
	a6 = a2

	// 修改索引为1的数组元素
	a7 := [3]string{"a", "b", "c"}
	a7[1] = "d"

	// 打印
	fmt.Println(a1, a2, a3, a4, a5, a6, a7)
}

func multi() {
	a1 := [3][3]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}

	fmt.Println(a1)
}

func foo(array [5]int) {
	fmt.Println(array)
}

func bar(array *[5]int) {
	fmt.Println(array)
}
