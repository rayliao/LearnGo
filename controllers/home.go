package controllers

import (
	"net/http"
	"os"

	"github.com/rayliao/go-learning/utils"
)

// View view
type View struct {
	name   string
	modal  interface{}
	action func(w http.ResponseWriter, r *http.Request)
}

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	os.Setenv("my-api", "xoxo")
	t := utils.LoadTemplate("views/index.html")
	t.Execute(w, os.Environ())
}

// About handler
func About(w http.ResponseWriter, r *http.Request) {
	t := utils.LoadTemplate("views/about.html")
	t.Execute(w, os.Getenv("api"))
}
