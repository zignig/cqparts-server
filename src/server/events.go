package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type Message struct {
	Name string
}

func event(c *gin.Context) {
	//list := wh.store.ListActive()
	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	c.Stream(func(w io.Writer) bool {
		select {
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
