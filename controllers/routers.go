package controllers

import "net/http"

type router struct {
	Pattern string
	Handler func(w http.ResponseWriter, r *http.Request)
}

// Routes all routes
var Routes = []router{
	{
		Pattern: "/",
		Handler: Home,
	},
	{
		Pattern: "/about",
		Handler: About,
	},
}
