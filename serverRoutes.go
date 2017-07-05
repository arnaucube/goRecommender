package main

import (
	"fmt"
	"net/http"
)

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"NewImage",
		"POST",
		"/image",
		NewImage,
	},
}

//ROUTES

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "send images to the /image path")
}
func NewImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "response")
}
