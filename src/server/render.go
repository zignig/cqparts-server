package main

import (
	"encoding/json"
	"io"

	"fmt"

	"archive/zip"
	"bytes"
	"strconv"

	"github.com/gin-gonic/gin"
)

var render_active bool

type XYZ struct {
	X float64 `form:"x" json:"x" `
	Y float64 `form:"y" json:"y" `
	Z float64 `form:"z" json:"z" `
}

// post reder post format
type Render struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Cam    XYZ    `form:"cam" json:"cam" binding:"required"`
	Target XYZ    `form:"target" json:"target" binding:"required"`
}

func zipped(c *gin.Context) {
	path := c.Params.ByName("name")
	data, err := store.Multi(path)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	for name, data := range data {
		fmt.Println("adding " + name)
		f, err := w.Create(name[1:])
		if err != nil {
			panic(err)
		}
		_, err = f.Write(data)
		if err != nil {
			panic(err)
		}
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}
	size := int64(buf.Len())
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	c.Writer.Header().Set("Content-Type", "application/zip")
	io.Copy(c.Writer, buf)
}

func receiveImage(c *gin.Context) {

}

func postrender(c *gin.Context) {
	var r Render
	err := c.ShouldBind(&r)
	fmt.Println(r, err)
	//if render_active {
	render_chan <- r
	//}
}

func render(c *gin.Context) {
	// fill up the list with the current store listing
	fmt.Println("RENDER ATTACHED")
	// render_active = true
	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-render_chan:
			data, _ := json.Marshal(msg)
			c.SSEvent("render", string(data))
		}
		return true
	})
	// render_active = false
	fmt.Println("RENDERER DETACHED")
}
