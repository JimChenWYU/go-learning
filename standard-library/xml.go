package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type (
	image struct {
		Url string `xml:"url"`
	}

	article struct {
		Title  string  `xml:"title"`
		Images []image `xml:"images"`
		Author struct {
			Name string `xml:"name"`
		} `xml:"author"`
	}

	good map[string]string
)

// 针对 map 需要转化为 XML 需要自行实现 MarshalXML 接口做 XML 转化处理
func (g good) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{start}

	for key, value := range g {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	decodeIOReader()
	decodeString()
	decodeStruct()
	decodeMap()
}

func decodeIOReader() {
	file, err := os.OpenFile("standard-library/test.xml", os.O_RDONLY, 0666)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	var at article

	err = xml.NewDecoder(file).Decode(&at)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(at)
}

func decodeString() {
	var XML = `
<article>
	<title>真的很好</title>
	<images>
		<url>https://images.unsplash.com/photo-1560981288-6bb2f13da3df?ixlib=rb-1.2.1&amp;q=80&amp;fm=jpg&amp;crop=entropy&amp;cs=tinysrgb&amp;w=1080&amp;fit=max</url>
	</images>
	<images>
		<url>https://images.unsplash.com/photo-1560981288-6bb2f13da3df?ixlib=rb-1.2.1&amp;q=80&amp;fm=jpg&amp;crop=entropy&amp;cs=tinysrgb&amp;w=1080&amp;fit=max</url>
	</images>
	<images>
		<url>https://images.unsplash.com/photo-1560981288-6bb2f13da3df?ixlib=rb-1.2.1&amp;q=80&amp;fm=jpg&amp;crop=entropy&amp;cs=tinysrgb&amp;w=1080&amp;fit=max</url>
	</images>
	<author>
		<name>John</name>
	</author>
</article>
`

	var a article

	err := xml.Unmarshal([]byte(XML), &a)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(a)
	fmt.Println("Title: ", a.Title)
	fmt.Println("Images0: ", a.Images[0].Url)
	fmt.Println("Author")
	fmt.Println("Name: ", a.Author.Name)
	//fmt.Println("Name: ", m["author"].(map[string]interface{})["name"])
}

func decodeStruct() {
	article := &article{
		Title: "HelloWorld",
		Images: []image{
			{
				Url: "https://images.unsplash.com/photo-1560981288-6bb2f13da3df?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=1080&fit=max",
			},
			{
				Url: "https://images.unsplash.com/photo-1561548146-1cd0500152ad?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=1080&fit=max",
			},
			{
				Url: "https://images.unsplash.com/photo-1560836056-bc8c61d1d8b6?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=1080&fit=max",
			},
		},
		Author: struct {
			Name string `xml:"name"`
		}{Name: string("John")},
	}

	data, err := xml.MarshalIndent(article, "", "  ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(data))
}

// 针对 map 需要转化为 XML 需要自行实现 MarshalXML 接口做 XML 转化处理
func decodeMap() {

	m := &good{
		"name":  "Apple",
		"price": "10.5",
	}

	data, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(data))
}
