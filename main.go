package main

import (
	"embed"
	"fmt"
	"html"
	"html/template"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"
	"net/http"
)

//go:embed *.gohtml
var fs embed.FS

type CRD struct {
	GVK    schema.GroupVersionKind
	Scoped string
	Bound  bool
	Icon   string
}

type CRDsInfo struct {
	SessionID, ClusterName string
	CRDs                   []CRD
}

func main() {
	http.HandleFunc("GET /bind", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(fs, "*.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(w, "resources.gohtml", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("POST /bind", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
