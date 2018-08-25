package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

// TODO include model info and stats
type Message struct {
	Name string
}

type Event struct {
	Section string `form:"section" json:"section"`
	Name    string `form:"name" json:"name"`
	Value   string `form:"value" json:"value"`
}

// Events from the web application
func postEvent(c *gin.Context) {
	var e Event
	err := c.ShouldBind(&e)
	if err != nil {
		fmt.Println(err)
		return
	}
	client_event <- e
}

// TODO build a pipeline defined by a YAML file in the assets
// load into the database for editing
func eventBus(exit chan bool) {
	for {
		select {
		case <-exit:
			return
		// track changed files
		case fi := <-files:
			fmt.Println("build this file ->", fi)
			build(fi)
		case inc := <-incoming:
			fmt.Println("posted incoming file ->", inc)
			m := mc.Find(inc)
			fmt.Println(m)
			menu <- m
		case r := <-render_update:
			fmt.Println("RENDER UPDATE", r)
			m := mc.Find(r.Name)
			fmt.Println("update me ", m)
			m.View = &r
			m.Img = "/pic/" + m.Name + ".png"
			mc.Put(m)
//			err := mc.Put(m)
//			if err == nil {
//				menu <- m
//			}
		case ce := <-client_event:
			fmt.Println("INCOMING EVENT ", ce)
			if ce.Section == "pin" {
				fmt.Println("PIN ME")
				m := mc.Find(ce.Name)
				m.Pinned = !m.Pinned
				mc.Put(m)
			}
		}
	}
}

// Base event runner
// need to convert to individual go routines per user with a broadcast channel
// change into a model/issue/stuff events
func event(c *gin.Context) {
	// fill up the list with the current store listing
	//fmt.Println(mc.List())
	//li, err := mc.List()
	//if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//for _, i := range li {
	//		menu <- i
	//	}
	ticker := time.NewTicker(20 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	c.Stream(func(w io.Writer) bool {
		select {
		// add to menu but don't load
		case msg := <-menu:
			fmt.Println("message -> |", msg, "|")
			data, _ := json.Marshal(msg)
			fmt.Println(string(data))
			c.SSEvent("menu", string(data))
		// add to menu and load
		//case msg := <-incoming:
		//	fmt.Println("message:", msg)
		//	mess := Message{Name: msg}
		//	data, _ := json.Marshal(mess)
		//	c.SSEvent("update", string(data))
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
