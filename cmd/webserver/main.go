package main

import (
	"net/http"
)

func main() {
	// start fileserver and create handler
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
}
