package Bitcoin

import (
	// "fmt"
	"net/http"
	"text/template"

	// functions "server/functions"
)

type ascii struct {
	AsciiArt string
	Error    string
}

func Home(w http.ResponseWriter, r *http.Request) {
	home := template.Must(template.ParseFiles("./web/templates/login.html"))
	w.WriteHeader(http.StatusOK)

	home.Execute(w, nil)
}


func Checkout(w http.ResponseWriter, r *http.Request) {
	home := template.Must(template.ParseFiles("./web/templates/checkout.html"))
	w.WriteHeader(http.StatusOK)

	home.Execute(w, nil)
}




func ErrorPage(w http.ResponseWriter, statusCode int, message string) {
	tmpl := template.Must(template.ParseFiles("./web/template/error.html"))
	w.WriteHeader(statusCode)
	data := ascii{Error: message}
	tmpl.Execute(w, data)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = registerUser("users.json", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User registered successfully")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	authenticated, err := authenticateUser("users.json", user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !authenticated {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Login successful")
}


// func Art(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		ErrorPage(w, http.StatusBadRequest, "400 - Bad Request")
// 		return
// 	}

// 	r.ParseForm()
// 	text := r.FormValue("input")
// 	banner := r.FormValue("bannerfile")
// 	// ascii_Art, err := functions.Input(text, banner)
// 	fmt.Println(err)

// 	if err != nil {
// 		if err.Error() == "file not found" {

// 			// w.WriteHeader(http.StatusNotFound)
// 			ErrorPage(w, http.StatusNotFound, "404 - Not Found")
// 			return
// 		}
// 		if err.Error() == ("non ascii Character") {
// 			ErrorPage(w, http.StatusInternalServerError, "400 - Bad Request")
// 			return
// 		}

// 		ErrorPage(w, http.StatusInternalServerError, "500 - Internal Server Error")
// 		return
// 	}

// 	// fmt.Println(ascii_Art)
// 	data := ascii{AsciiArt: ascii_Art}

// 	home := template.Must(template.ParseFiles("template/index.html"))
// 	err2 := home.Execute(w, data)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}
// }
