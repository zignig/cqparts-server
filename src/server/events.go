package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO include model info and stats
type Message struct {
	Name string
}

func event(c *gin.Context) {
	// fill up the list with the current store listing
	fmt.Println(models.List())
	for _, i := range models.List() {
		menu <- i
	}
	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	c.Stream(func(w io.Writer) bool {
		select {
		// add to menu but don't load
		case msg := <-menu:
			mess := Message{Name: msg}
			data, _ := json.Marshal(mess)
			c.SSEvent("menu", string(data))
		// add to menu and load
		case msg := <-incoming:
			fmt.Println("message:", msg)
			mess := Message{Name: msg}
			data, _ := json.Marshal(mess)
			c.SSEvent("update", string(data))
		case <-ticker.C:
			c.SSEvent("tick", "")
		}
		return true
	})
}
