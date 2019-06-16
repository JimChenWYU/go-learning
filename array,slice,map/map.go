package main

import (
	"fmt"
)

// 映射是一个存储键值对的无序集合，其底层实现是散列表
// 映射在函数之间传递是按引用传递
func main() {
	// 创建一个映射
	colors := map[string]string{"Red": "#da1337", "Orange": "#FFD306"}

	// 使用映射
	colors["Red"] = "#da1337"

	// 创建 nil 映射
	// 需要注意的是 nil 映射不能用于存储键值对
	var map1 map[string]string
	fmt.Println(map1)

	// 映射取值并判断存在性
	// exists 返回的是 bool 值，存在为true，不存在为false
	// 映射不存在键，则值会返回 nil，一般情况也可作为判断依据，但特殊情况下，如果对应键的值为 nil 则不可以
	red, exists := colors["Red"]
	if exists {
		fmt.Println(red)
	}

	// 使用 range 迭代映射
	for key, value := range colors {
		fmt.Printf("key=%s, value=%s \n", key, value)
	}

	// delete 删除映射元素
	delete(colors, "Red")
	fmt.Println(colors)
}
