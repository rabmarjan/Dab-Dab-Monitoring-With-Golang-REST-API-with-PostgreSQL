package main

import (
	"fmt"
	"log"
	"net/http"
	"workspace/goweb/routers"
)

func Hell(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested:%s\n", r.FormValue("post"))
}

func main() {
	r := routers.GetRoute()
	r.HandleFunc("/", Hell)
	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Server run at port :8070")
	log.Fatal(http.ListenAndServe(":8070", r))

}
