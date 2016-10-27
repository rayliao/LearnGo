package main

import (
	"net/http"

	"github.com/rayliao/go-learning/controllers"
)

func init() {
	for _, route := range controllers.Routes {
		http.HandleFunc(route.Pattern, route.Handler)
	}
}

func main() {
	http.ListenAndServe(config["port"].(string), nil)
}
