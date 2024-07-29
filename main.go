package main

import (
	"fmt"
	"net/http"
	"os"

	Bitcoin "Bitcoin/server"
)

func main() {
	if len(os.Args) != 1 {
		return
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // fs := http.FileServer(http.Dir("static"))
	// assets := http.FileServer(http.Dir("assets"))
	http.HandleFunc("/", Bitcoin.Home)
	// http.HandleFunc("/ascii-art", Bitcoin.Art)
	// http.Handle("/static/", http.StripPrefix("/assets/static/", fs))
	// http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	// http.HandleFunc("/about", Bitcoin.About)
	// http.HandleFunc("/instructions", Bitcoin.Instructions)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/home":
			Bitcoin.Home(w, r)
		default:
			Bitcoin.ErrorPage(w, http.StatusNotFound, "404 - Not Found")
		}
	})
	fmt.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func Serve(w http.ResponseWriter, r *http.Request) {
	imagePath := "./" + r.URL.Path
	_, err := os.Stat(imagePath)
	if err != nil {
		fmt.Fprintf(w, "Image not found")
	}

	content, _ := os.ReadFile(imagePath)
	fmt.Fprintf(w, "%s", string(content))
}
