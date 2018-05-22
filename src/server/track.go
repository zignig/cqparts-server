package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func watch(name string) {
	files := make(chan string, 10)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}
	watcher.Add(name)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("event", event)
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("removed file add again: ", event.Name)
					watcher.Add(name)
					files <- event.Name
				}
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
			fmt.Println("next")
		}
	}()
	go func() {
		for {
			select {
			case action := <-files:
				fmt.Println("RUN THIS")
				build(name, action)
			}
		}
	}()
}
