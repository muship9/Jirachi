package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once sync.Once
	fileName string
	temple   *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter , r *http.Request)  {
	t.once.Do(func() {
		t.temple = template.Must(template.ParseFiles(filepath.Join("templates" , t.fileName)))
	})
	t.temple.Execute(w, nil)
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{fileName: "index.html"})
	http.Handle("/room", r)
	go r.run()
	if err := http.ListenAndServe(":8080", nil);
	err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
