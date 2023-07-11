package main

import (
	"io/ioutil"
	"net/http"
)

type indexHandler struct {
	content string
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Hello", "www")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func main() {
	http.Handle("/", &indexHandler{content: "hello world!"})
	http.ListenAndServe(":8001", nil)
}
