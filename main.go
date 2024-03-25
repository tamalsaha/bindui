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

		info := CRDsInfo{
			SessionID:   "sid",
			ClusterName: "my-cluster",
			CRDs: []CRD{
				{
					GVK: schema.GroupVersionKind{
						Group:   "kubedb.com",
						Version: "v1alpha2",
						Kind:    "MongoDB",
					},
					Scoped: "Namespaced",
					Bound:  false,
					Icon:   "https://cdn.appscode.com/k8s/icons/kubedb.com/mongodbs.svg",
				},
				{
					GVK: schema.GroupVersionKind{
						Group:   "kubedb.com",
						Version: "v1alpha2",
						Kind:    "MySQL",
					},
					Scoped: "Namespaced",
					Bound:  true,
					Icon:   "https://cdn.appscode.com/k8s/icons/kubedb.com/mysqls.svg",
				},
				{
					GVK: schema.GroupVersionKind{
						Group:   "kubedb.com",
						Version: "v1alpha2",
						Kind:    "Postgres",
					},
					Scoped: "Namespaced",
					Bound:  false,
					Icon:   "https://cdn.appscode.com/k8s/icons/kubedb.com/postgreses.svg",
				},
			},
		}
		err = tmpl.ExecuteTemplate(w, "resources.gohtml", info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("POST /bind", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
