package main

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"fmt"

	"archive/zip"
	"bytes"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Is the renderer running
var render_active bool

// 3d point structure
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

func NewRender(name string) (r *Render){
	r =  &Render{}
	r.Cam = XYZ{
		X:0.2,
		Y:0.2,
		Z:0.2,
	}
	return r
}

// TODO add error checking
// add json datafile and even the python code
// Hands a zip files of all files in a model
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

// TODO version multi store for each model
func receiveImage(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	files := form.File["objs"]
	for _, j := range files {
		f, err := j.Open()
		if err != nil {
			panic(err)
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		images.Save(j.Filename, data)
	}
}

// Take a render event from the web interface
// and pass id down to the render stream
func postrender(c *gin.Context) {
	var r Render
	err := c.ShouldBind(&r)
	fmt.Println(r, err)
	b, _ := json.Marshal(r)
	view.Save(r.Name, b)
	//if render_active {
	render_chan <- r
	render_update <- r
	//}
}

func renderAll(c *gin.Context) {
	data, err := mc.List()
	if err != nil {
		return
	}
	for _, i := range data {
		fmt.Println("rendering ", i.Name)
		render_chan <- *i.View
	}
}

// Get a picture out of the storage
func pic(c *gin.Context) {
	path := c.Params.ByName("name")
	data, err := images.Load(path)
	if err != nil {
		c.AbortWithStatus(404)
	}
	size := int64(len(data))
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	io.Copy(c.Writer, bytes.NewReader(data))
}

// Blender renderer attaches to here
// /src/blender/
// TODO pass down render sizes and handle a group of
// render daemons
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
