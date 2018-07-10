package main

import (
	"html/template"
	"log"
	"strings"
	"syscall/js"
	"time"
)

var viewTmpl = `<h1>{{.Title}}</h1>
<p>{{.Text}}</p>
`

type view struct {
	Title string
	Text  string
}

func main() {
	log.SetFlags(log.Lshortfile)
	log.Printf("msg=%q\n", "Hello from Go!")

	document := js.Global().Get("document")
	app := document.Call("querySelector", "#app")
	log.Printf("app=%v\n", app)

	tmpl, err := template.New("viewTmpl").Parse(viewTmpl)
	if err != nil {
		log.Printf("err=%v\n", err)
		return
	}
	var buf strings.Builder
	err = tmpl.Execute(&buf, view{
		Title: "Go WASM",
		Text:  "Loaded!",
	})
	if err != nil {
		log.Printf("err=%v\n", err)
		return
	}
	app.Set("innerHTML", buf.String())

	time.Sleep(1 * time.Second)
	buf.Reset()
	err = tmpl.Execute(&buf, view{
		Title: "Go WASM",
		Text:  "Changed after 1 sec.",
	})
	if err != nil {
		log.Printf("err=%v\n", err)
		return
	}
	app.Set("innerHTML", buf.String())

	log.Printf("msg=%q\n", "done")
}
