package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
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

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
