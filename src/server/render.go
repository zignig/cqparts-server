package main

import (
	"encoding/json"
	"io"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
)

// post reder post format
type Render struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func postrender(c *gin.Context) {
	var r Render
	err := c.ShouldBind(&r)
	fmt.Println(r, err)
	render_chan <- r.Name
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
