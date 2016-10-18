package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// Fetch func
func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		// b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%v", b)
		doc, err := html.Parse(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(doc)
	}
}
