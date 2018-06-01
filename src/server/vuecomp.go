package main

// view component loader
import (
	"bytes"
	"errors"
	"fmt"
	"log"

	"golang.org/x/net/html"
)

func GetVue(name string, data []byte) (node *html.Node, err error) {
	fmt.Println("ARRRHG")
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "template" {
			b = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("missing <template>")
}

func ProcVues() {
	pages, err := AssetDir("asset/vue")
	if err != nil {
		return
	}
	for _, j := range pages {
		fmt.Println(j)
		data, _ := Asset("asset/vue/" + j)
		n, err := GetVue(j, data)
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
	}
}
