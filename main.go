package main

import (
	"fmt"
	"net/http"
	"os"
	"Bitcoin/server"
	
)

func main() {
	if len(os.Args) != 1 {
		return
	}
	fs := http.FileServer(http.Dir("static"))
	// assets := http.FileServer(http.Dir("assets"))
	http.HandleFunc("/", Bitcoin.Home)
	// http.HandleFunc("/ascii-art", handler.Art)
	http.Handle("/static/", http.StripPrefix("/assets/static/", fs))
	// http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	// http.HandleFunc("/about", handler.About)
	// http.HandleFunc("/instructions", handler.Instructions)
	fmt.Println("Server is starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
