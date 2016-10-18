package ch5

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	// 如果当前节点是a标签
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// 循环遍历子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// 递归
		links = visit(links, c)
	}
	return links
}

// FindLinks func
func FindLinks() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		n, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println(n)
		doc, err := html.Parse(os.Stdout)
		fmt.Println("output doc files*********")
		fmt.Println(doc)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		// 	os.Exit(1)
		// }
		// for _, link := range visit(nil, doc) {
		// 	fmt.Println(link)
		// }
	}
}
