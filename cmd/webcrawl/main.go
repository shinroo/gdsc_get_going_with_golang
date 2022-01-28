package main

import (
	"github.com/shinroo/gdsc_get_going_with_golang/pkg/webcrawl"
)

func main() {
	webcrawl.Crawl("https://golang.org/", 4, webcrawl.Example)
}
