//go:generate go-bindata --debug asset/...
package main

import "fmt"

func main() {
	fmt.Println("Process Vue Components")
	ProcVues()
}
