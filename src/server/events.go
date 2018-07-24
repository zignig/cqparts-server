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

type Event struct {
	Section string `form:"section" json:"section"`
	Name    string `form:"name json:"name"`
}

// Events from the web application
func postEvent(c *gin.Context) {
	var e Event
	err := c.ShouldBind(&e)
	fmt.Println(e, err)
}

// TODO build a pipeline defined by a YAML file in the assets
// load into the database for editing
func eventBus(exit chan bool) {
	for {
		select {
		case <-exit:
			return
		}
	}
}

// Base event runner
// need to convert to individual go routines per user with a broadcast channel
// change into a model/issue/stuff events
func event(c *gin.Context) {
	// fill up the list with the current store listing
	fmt.Println(models.List())
	for _, i := range models.List() {
		menu <- i
	}
	ticker := time.NewTicker(20 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	c.Stream(func(w io.Writer) bool {
		select {
		// add to menu but don't load
		case msg := <-menu:
			fmt.Println("message -> |", msg, "|")
			mess := Model{
				Name: msg,
				Img:  "/pic/" + msg + ".png",
			}
			data, _ := json.Marshal(mess)
			fmt.Println(string(data))
			c.SSEvent("menu", string(data))
		// add to menu and load
		case msg := <-incoming:
			fmt.Println("message:", msg)
			mess := Message{Name: msg}
			data, _ := json.Marshal(mess)
			c.SSEvent("update", string(data))
		// sending issues
		case msg := <-issue:
			mess := Message{Name: msg}
			data, _ := json.Marshal(mess)
			c.SSEvent("issue", string(data))
		// ticker
		case <-ticker.C:
			c.SSEvent("tick", "")
		}
		return true
	})
}
