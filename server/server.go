package server

import (
	"fmt"
	"net/http"
	"text/template"

	functions "server/functions"
)

type ascii struct {
	AsciiArt string
	Error    string
}

func Home(w http.ResponseWriter, r *http.Request) {
	home := template.Must(template.ParseFiles("./template/index.html"))
	w.WriteHeader(http.StatusOK)

	home.Execute(w, nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	home := template.Must(template.ParseFiles("./template/about.html"))
	w.WriteHeader(http.StatusOK)

	home.Execute(w, nil)
}

func Instructions(w http.ResponseWriter, r *http.Request) {
	home := template.Must(template.ParseFiles("./template/instructions.html"))
	w.WriteHeader(http.StatusOK)

	home.Execute(w, nil)
}

func ErrorPage(w http.ResponseWriter, statusCode int, message string) {
	tmpl := template.Must(template.ParseFiles("template/error.html"))
	w.WriteHeader(statusCode)
	data := ascii{Error: message}
	tmpl.Execute(w, data)
}

func Art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorPage(w, http.StatusBadRequest, "400 - Bad Request")
		return
	}

	r.ParseForm()
	text := r.FormValue("input")
	banner := r.FormValue("bannerfile")
	ascii_Art, err := functions.Input(text, banner)
	fmt.Println(err)

	if err != nil {
		if err.Error() == "file not found" {

			// w.WriteHeader(http.StatusNotFound)
			ErrorPage(w, http.StatusNotFound, "404 - Not Found")
			return
		}
		if err.Error() == ("non ascii Character") {
			ErrorPage(w, http.StatusInternalServerError, "400 - Bad Request")
			return
		}

		ErrorPage(w, http.StatusInternalServerError, "500 - Internal Server Error")
		return
	}

	// fmt.Println(ascii_Art)
	data := ascii{AsciiArt: ascii_Art}

	home := template.Must(template.ParseFiles("template/index.html"))
	err2 := home.Execute(w, data)
	if err2 != nil {
		fmt.Println(err2)
	}
}
