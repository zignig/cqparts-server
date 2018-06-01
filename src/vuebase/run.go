//go:generate go-bindata --debug asset/...
package main

import (
	"fmt"
	"text/template"
)

func main() {
	fmt.Println("Process Vue Components")
	tmpl := template.New("")
	scripts, css := ProcVues(tmpl)
	fmt.Println(scripts)
	fmt.Println("-----------------")
	fmt.Println(css)
	fmt.Println("-----------------")
	fmt.Println(tmpl.DefinedTemplates())
}
