package main

import (
	"crawder/single/frontend/controller"
	"net/http"
)

func main()  {
	http.Handle("/", http.FileServer(http.Dir("single/frontend/view")))
	http.Handle("/search", controller.CreateHandler("single/frontend/view/template.html"))
	err := http.ListenAndServe("localhost:8888", nil)
	if err != nil {
		panic(err)
	}
}
