package main

import (
	"fmt"
	"os/exec"
	"time"
)

func build(name string, file string) bool {
	n := time.Now()
	fmt.Println(n)
	cmd := exec.Command("/usr/bin/python2.7", file)
	output, err := cmd.CombinedOutput()
	t := time.Since(n)
	fmt.Println(t)
	fmt.Println(name, file)
	fmt.Println(string(output), err)
	if err == nil {
		//		incoming <- name
	}
	return false
}
