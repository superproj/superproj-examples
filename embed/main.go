package main

import (
	"embed"
	"net/http"
)

//go:embed assets
var assets embed.FS

func main() {
	fs := http.FileServer(http.FS(assets))
	http.ListenAndServe(":8080", fs)
}
