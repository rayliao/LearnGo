package main

import (
	"log"
	"net/http"

	"fmt"
	"strings"
	// "github.com/rayliao/go-learning/controllers"
)

func init() {
	// for _, route := range controllers.Routes {
	// 	http.HandleFunc(route.Pattern, route.Handler)
	// }
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数
	fmt.Println(r.Form)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Rayliao")
}

func main() {
	// http.ListenAndServe(config["port"].(string), nil)
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
