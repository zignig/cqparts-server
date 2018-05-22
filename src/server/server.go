//go:generate go-bindata --debug asset/...
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"

	"flag"
	"github.com/gin-gonic/gin"
)

var incoming chan string

func main() {
	fileToWatch := flag.String("file", "testing.py", "file to watch")
	flag.Parse()
	r := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	//r.GET("/model")
	r.GET("/static/*name", Static)
	r.StaticFS("/model", http.Dir("/tmp/cqpss"))
	r.GET("/events", event)
	r.POST("/notify", notify)
	// web server
	incoming = make(chan string, 100)
	go r.Run(":8080")
	// file watcher
	fmt.Println("watching :", *fileToWatch)
	go watch(*fileToWatch)
	done := make(chan bool)
	<-done
}

func notify(c *gin.Context) {
	name, err := c.GetPostForm("name")
	fmt.Println(name, err)
	incoming <- name
}

func Static(c *gin.Context) {
	path := c.Params.ByName("name")
	data, err := Asset("asset" + path)
	if err != nil {
		c.AbortWithStatus(404)
	}
	if strings.HasSuffix(path, ".css") {
		c.Writer.Header().Set("Content-Type", "text/css")
	}
	if strings.HasSuffix(path, ".js") {
		c.Writer.Header().Set("Content-Type", "text/javascript")
	}
	size := int64(len(data))
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	io.Copy(c.Writer, bytes.NewReader(data))
}

func loadTemplate() (templ *template.Template, e error) {
	pages, err := AssetDir("asset/html")
	if err != nil {
		return
	}
	templ = template.New("")
	for _, j := range pages {
		fmt.Println(j)
		data, _ := Asset("asset/html/" + j)
		_, err = templ.New(j).Parse(string(data))
		if err != nil {
			return nil, err
		}
	}
	return templ, nil
}
