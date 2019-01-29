//go:generate go-bindata --debug asset/...
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"flag"

	"github.com/gin-gonic/gin"
)

// TODO wrap this into a scruct ( per user )
// TODO add cookies for session control
var incoming chan string
var issue chan string
var files chan string
var client_event chan Event
var model_chan chan Model
var menu chan *Model
var render_chan chan Render
var render_update chan Render
var css string
var script string
var store Storage  // model file storage
var models Storage // list of model names
var images Storage // list of model image
var view Storage   // list of model views
var mc *ModelCollection

// TODO add authentication and session per users
func main() {
	// Wrap all these into a struct
	bolty := NewBBoltStore("data")
	store = bolty.NewBucket("names")
	models = bolty.NewBucket("models")
	images = bolty.NewBucket("images")
	view = bolty.NewBucket("view")
	mc = NewModelCollection("default", bolty)

	incoming = make(chan string, 100)
	files = make(chan string, 100)
	menu = make(chan *Model, 100)
	issue = make(chan string, 100)
	render_chan = make(chan Render, 100)
	render_update = make(chan Render, 100)
	model_chan = make(chan Model, 100)
	client_event = make(chan Event, 100)

	// wrap me

	fileToWatch := flag.String("d", "./", "folder to watch")
	module := flag.String("m", "./", "name of the module")
	port := flag.Int("p", 8080, "port to listen on")

	flag.Parse()
	// base web server
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
	r.GET("/layout", func(c *gin.Context) {
		val := gin.H{}
		c.HTML(http.StatusOK, "layout.tmpl", val)
	})

	r.GET("/vue/comp.js", CompJS)
	r.GET("/vue/comp.css", CompCSS)

	r.GET("/static/*name", Static)
	r.GET("/model/*name", model)
	r.GET("/events", event)
	r.GET("/list", list)

	// Event stuff
	r.POST("/postevent", postEvent)
	r.GET("/status", status)
	r.POST("/notify", notify)
	r.POST("/upload", upload)

	r.POST("/snapshot", snapshot)

	// blender render endpoints
	r.GET("/render", render)
	// spool all the models to be rendered
	r.GET("/renderall", renderAll)
	r.POST("/postrender", postrender)
	r.GET("/zipped/:name", zipped)
	r.GET("/pic/:name", pic)
	r.POST("/image", receiveImage)

	// web server
	go r.Run(":" + strconv.Itoa(*port))
	// file watcher
	fmt.Println("watching :", *fileToWatch)
	go watch(*fileToWatch, *module)
	done := make(chan bool)
	go eventBus(done)
	//fmt.Println(models.List())
	// fmt.Println(store.List())
	//fmt.Println(images.List())
	//	fmt.Println(bolty.names.List())
	<-done
}

// TODO save threejs snapshot into the database
func snapshot(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(data)
}

// Cqparts display aliveness test
func status(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

// Upload from cqparts display
// gltf files
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

func list(c *gin.Context) {
	data, err := mc.List()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, data)
}

// push from cqparts display info blob
func notify(c *gin.Context) {
	name, err := c.GetPostForm("name")
	fmt.Println(name, err)
	incoming <- name
}

// vuejs composited scripts
func CompJS(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/javascript")
	io.Copy(c.Writer, strings.NewReader(script))
}

// vuejs composited css
func CompCSS(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/css")
	io.Copy(c.Writer, strings.NewReader(css))
}

// get out.glft for mode name
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

// serve the static content
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

// load the and build the templates from the assets
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
