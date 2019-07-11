package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	slide struct {
		Title string `json:"title"`
		Type  string `json:"type"`
	}

	// 映射返回结果
	// 使用标签的方式试图将 JSON 文档和结构类型的字段一一映射起来，
	// 如果不存在标签，编码和解码过程会试图以大小写无关的方式，直接使用字段的名字进行匹配。
	result struct {
		Slideshow struct {
			Author string  `json:"author"`
			Date   string  `json:"date"`
			Slides []slide `json:"slides"`
			Title  string  `json:"title"`
		} `json:"slideshow"`
	}

	person struct {
		Name    string `json:"name"`
		Title   string `json:"title"`
		Contact struct {
			Home string `json:"home"`
			Cell string `json:"cell"`
		} `json:"contact"`
	}

	user struct {
		Name string `json:"name"`
	}

	admin struct {
		User  user   `json:"user"`
		Level string `json:"level"`
	}
)

func main() {
	decodeIOReader()
	decodeString()
	encodeMap()
	encodeStruct()
}

func decodeIOReader() {
	uri := "https://httpbin.org/json"
	r, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	defer r.Body.Close()

	// 将 JSON 响应解码到结构类型
	var rt result

	err = json.NewDecoder(r.Body).Decode(&rt)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(rt)
}

func decodeString() {
	var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "413.333.3333",
		"cell": "415.555.5555"
	}
}`

	var p person
	var m map[string]interface{}

	err := json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	err = json.Unmarshal([]byte(JSON), &m)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(p)
	fmt.Println("Name: ", m["name"])
	fmt.Println("title: ", m["title"])
	fmt.Println("Contact")
	fmt.Println("Home: ", m["contact"].(map[string]interface{})["home"])
	fmt.Println("Cell: ", m["contact"].(map[string]interface{})["cell"])
}

func encodeMap() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(data))
}

func encodeStruct() {

	// 结构体里包含一个匿名结构
	p := &person{
		Name:  "Gopher",
		Title: "programmer",
	}
	p.Contact.Cell = "415.333.3333"
	p.Contact.Home = "415.555.5555"

	data1, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Println("ERROR: ", err)
	}

	fmt.Println(string(data1))

	// 嵌入结构
	admin := &admin{
		User: user{
			Name: "john",
		},
		Level: "super",
	}

	data2, err := json.MarshalIndent(admin, "", "  ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(data2))
}
