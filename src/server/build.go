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
	if err == nil {
		//		incoming <- name
		fmt.Println(string(output), err)
	}
	return false
}
