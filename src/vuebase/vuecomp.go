package main

// view component loader
import (
	"bytes"
	"errors"
	"io"
	"log"
	"strings"

	"text/template"

	"golang.org/x/net/html"
)

type VueParse struct {
	Name     string
	Template string
	Script   string
	Css      string
}

var scriptTmpl string = ` var {{ .name }} = Vue.component('{{ .name }}',
{ template : '#{{.name}}',
{{ .script }} 
}); `

func GetVue(name string, data []byte) (component *VueParse, err error) {
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	// just the name without the .vue
	trimmedName := strings.Split(name, ".")[0]
	component = &VueParse{Name: trimmedName}
	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "template" {
			b = n
			b.Attr = append(b.Attr, html.Attribute{Key: "id", Val: trimmedName})
			component.Template = renderNode(b)
		}
		if n.Type == html.ElementNode && n.Data == "style" {
			style := renderNode(n.FirstChild)
			component.Css = style
		}
		if n.Type == html.ElementNode && n.Data == "script" {
			script := renderNode(n.FirstChild)
			// trim the export off so it default componeints work
			start := strings.Index(script, "{")
			finish := strings.LastIndex(script, "}")
			trimmedScript := script[start+2 : finish-2]
			// template the script section
			tmpl := template.New("scr")
			tmpl, err := tmpl.Parse(scriptTmpl)
			if err != nil {
				panic(err)
			}
			values := map[string]interface{}{"name": trimmedName, "script": trimmedScript}
			var tpl bytes.Buffer
			if err := tmpl.Execute(&tpl, values); err != nil {
				panic(err)
			}
			component.Script = html.UnescapeString(tpl.String())
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return component, nil
	return nil, errors.New("missing <template>")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func ProcVues(tmpl *template.Template) (script string, css string) {
	pages, err := AssetDir("asset/vue")
	if err != nil {
		return
	}
	vues := make([]*VueParse, 0)
	for _, j := range pages {
		data, _ := Asset("asset/vue/" + j)
		n, err := GetVue(j, data)
		if err != nil {
			panic(err)
		}
		vues = append(vues, n)
	}
	// get templates
	html_template := ""
	for _, j := range vues {
		html_template += j.Template
		html_template += "\n\n"
	}
	_, err = tmpl.New("vuetmpl").Delims("[[", "]]").Parse(html_template)
	if err != nil {
		panic(err)
	}
	// Css
	css_string := ""
	for _, j := range vues {
		css_string += j.Script
		css_string += "\n\n"
	}
	css = css_string
	// Script
	script_string := ""
	for _, j := range vues {
		script_string += j.Css
		script_string += "\n\n"
	}
	script = script_string
	return
}
