package Bitcoin

import (
	// "fmt"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"text/template"
	// functions "server/functions"
)

type UserID struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type ascii struct {
	AsciiArt string
	Error    string
}
var mu sync.Mutex


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

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user UserID
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = RegisterUser("users.json", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User registered successfully")
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user UserID
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

func RegisterUser(filename string, newUser UserID) error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	var users []UserID
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil && err != io.EOF {
		return err
	}

	for _, user := range users {
		if user.Username == newUser.Username {
			return errors.New("user already exists")
		}
	}

	users = append(users, newUser)

	file.Truncate(0)
	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(filename, username, password string) (bool, error) {
    mu.Lock()
    defer mu.Unlock()

    file, err := os.Open(filename)
    if err != nil {
        return false, err
    }
    defer file.Close()

    var users []UserID
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&users)
    if err != nil && err != io.EOF {
        return false, err
    }

    for _, user := range users {
        if user.Username == username && user.Password == password {
            return true, nil
        }
    }

    return false, nil
}
