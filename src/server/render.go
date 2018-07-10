package main

import (
	"encoding/json"
	"io"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
)

func postrender(c *gin.Context) {
	var name string
	name, err := c.GetPostForm("name")
	if err == false {
		name = "ERROR"
	}
	render_chan <- name
}

func render(c *gin.Context) {
	// fill up the list with the current store listing
	fmt.Println("RENDER ATTACHED")
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-render_chan:
			mess := Message{Name: msg}
			data, _ := json.Marshal(mess)
			c.SSEvent("render", string(data))
		case <-ticker.C:
			c.SSEvent("tick", "")
		}
		return true
	})
}
