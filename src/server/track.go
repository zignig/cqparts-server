package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// Watch for file changes
func watch(name string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}
	filelist, err := ioutil.ReadDir(name)
	if err != nil {
		fmt.Println("LIST FAIL")
		return
	}
	for _, f := range filelist {
		n := f.Name()
		if strings.HasSuffix(n, ".py") {
			//fmt.Println(n)
			watcher.Add(name + n)
			// TODO add to the model store
		}
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("event", event)
				if (event.Op&fsnotify.Remove == fsnotify.Remove) || (event.Op&fsnotify.Write == fsnotify.Write) {
					fmt.Println("removed file add again: ", event.Name)
					watcher.Add(event.Name)
					files <- event.Name
				}
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
			fmt.Println("next")
		}
	}()
}
