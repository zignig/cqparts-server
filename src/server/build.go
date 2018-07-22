package main

import (
	"fmt"
	"os/exec"
	"time"
)

// run python on a model file
// if there is an error
// turn issue output into formatted array
func build(name string, file string) bool {
	n := time.Now()
	cmd := exec.Command("/usr/bin/python2.7", file)
	output, err := cmd.CombinedOutput()
	t := time.Since(n)
	fmt.Println(t)
	if err != nil {
		issue <- string(output)
		fmt.Println(string(output), err)
	}
	return false
}
