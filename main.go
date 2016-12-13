package main

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/login.html")
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		log.Println(t.Execute(w, token))
	} else {
		r.ParseForm()
		if len(r.Form["username"][0]) == 0 {
			fmt.Println("username is empty")
		}
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Println("It's not a number")
		}
		// or use reg
		// if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		// 	return false
		// }
		if getint > 100 {
			fmt.Println("Age is too big")
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func socket(ws *websocket.Conn) {
	var err error

	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	// http.ListenAndServe(config["port"].(string), nil)
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.Handle("/msg", websocket.Handler(socket))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
