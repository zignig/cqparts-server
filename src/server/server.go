//go:generate go-bindata --debug asset/...
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	//"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"flag"

	"github.com/gin-gonic/gin"
)

var incoming chan string
var menu chan string
var issue chan string
var render_chan chan string
var css string
var script string
var store Storage  // model file storage
var models Storage // list of model names

func main() {
	models = NewMemStore("models")
	store = NewBBoltStore("hello")
	incoming = make(chan string, 100)
	menu = make(chan string, 100)
	issue = make(chan string, 100)
	render_chan = make(chan string, 100)

	fileToWatch := flag.String("d", "./", "folder to watch")
	flag.Parse()
	r := gin.Default()
	//r := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	script, css = ProcVues(t)
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *gin.Context) {
		val := gin.H{}
		c.HTML(http.StatusOK, "index.tmpl", val)
	})
	r.GET("/dev", func(c *gin.Context) {
		val := gin.H{}
		c.HTML(http.StatusOK, "dev.tmpl", val)
	})

	r.GET("/vue/comp.js", CompJS)
	r.GET("/vue/comp.css", CompCSS)

	r.GET("/static/*name", Static)
	r.GET("/model/*name", model)
	r.GET("/events", event)

	r.GET("/status", status)
	r.POST("/notify", notify)
	r.POST("/upload", upload)

	r.POST("/snapshot", snapshot)

	// blender render endpoints
	r.GET("/render", render)
	r.POST("/postrender", postrender)

	// web server
	go r.Run(":8080")
	// file watcher
	fmt.Println("watching :", *fileToWatch)
	go watch(*fileToWatch)
	done := make(chan bool)
	<-done
}

func snapshot(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(data)
}

func status(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

func upload(c *gin.Context) {
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
		store.Save("/"+j.Filename, data)
	}
}

func notify(c *gin.Context) {
	name, err := c.GetPostForm("name")
	fmt.Println(name, err)
	models.Save(name, []byte{'_'})
	incoming <- name
}

func CompJS(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/javascript")
	io.Copy(c.Writer, strings.NewReader(script))
}

func CompCSS(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/css")
	io.Copy(c.Writer, strings.NewReader(css))
}

func model(c *gin.Context) {
	path := c.Params.ByName("name")
	data, err := store.Load(path)
	if err != nil {
		c.AbortWithStatus(404)
	}
	size := int64(len(data))
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	io.Copy(c.Writer, bytes.NewReader(data))
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
		_, err = templ.New(j).Delims("[[", "]]").Parse(string(data))
		if err != nil {
			return nil, err
		}
	}
	return templ, nil
}
