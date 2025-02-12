package main

import (
	"fmt"
	"net/http"
	"os"

	Bitcoin "Bitcoin/server"
)


// type User struct {
//     Username string `json:"username"`
//     Password string `json:"password"`
// }

func main() {
	if len(os.Args) != 1 {
		return
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // fs := http.FileServer(http.Dir("static"))
	// assets := http.FileServer(http.Dir("assets"))
	http.HandleFunc("/home", Bitcoin.Home)
	http.HandleFunc("/checkout", Bitcoin.Checkout)
	http.HandleFunc("/login", Bitcoin.HandleLogin)
	http.HandleFunc("/registeration", Bitcoin.HandlerRegisterPAge)
	http.HandleFunc("/register", Bitcoin.HandleRegister)
	// http.HandleFunc("/validate", Bitcoin.HandleValidate)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/login":
			Bitcoin.Home(w, r)
		default:
			// Bitcoin.ErrorPage(w, http.StatusNotFound, "404 - Not Found")
		}
	})
	fmt.Println("Server is running on port :http://localhost:8080")
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
