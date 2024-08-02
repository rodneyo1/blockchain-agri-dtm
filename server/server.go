package Bitcoin

import (
	// "fmt"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"text/template"
	// functions "server/functions"
)

type UserID struct {
	Firstname string `json:"first"`
	Lastname  string `json:"last"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Location  string `json:"location"`
	Contract  string `json:"mobile"`
	Gender  string `json:"gender"`
}

var mu sync.Mutex

// func Login(w http.ResponseWriter, r *http.Request) {
// 	if r.Method!=http.MethodGet{
// 			ErrorPage(w,http.StatusNotFound,"Not found")
// 	}
// 	home := template.Must(template.ParseFiles("./web/templates/login.html"))
// 	w.WriteHeader(http.StatusOK)

// 	home.Execute(w, nil)
// }
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method!=http.MethodGet{
			ErrorPage(w,http.StatusNotFound,"Not found")
	}
	home := template.Must(template.ParseFiles("./web/templates/index.html"))
	w.WriteHeader(http.StatusOK)

	home.Execute(w, nil)
}

func ErrorPage(w http.ResponseWriter, statusCode int, message string) {
	Error, err := template.New("error.html").ParseFiles("./web/templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(statusCode)

	Error.Execute(w, nil)
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	if r.Method!=http.MethodGet{
		ErrorPage(w,http.StatusNotFound,"Not found")
}
	home, err := template.New("checkout.html").ParseFiles("./web/templates/checkout.html")
	if err != nil {
		ErrorPage(w, http.StatusNotFound, "Not found")
	}
	w.WriteHeader(http.StatusOK)

	home.Execute(w, nil)
}

//	func ErrorPage(w http.ResponseWriter, statusCode int, message string) {
//		tmpl := template.Must(template.ParseFiles("./web/template/error.html"))
//		w.WriteHeader(statusCode)
//		data := ascii{Error: message}
//		tmpl.Execute(w, data)
//	}
func HandlerRegisterPAge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPage(w, http.StatusInternalServerError, "400 Bad Request")
		return
	}
	Register, err := template.New("registration.html").ParseFiles("./web/templates/registration.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, "400 Bad Request")
		return
	}
	w.WriteHeader(http.StatusOK)
	Register.Execute(w, nil)
}

func HandleLogin(w http.ResponseWriter,r *http.Request){
login,err:=template.New("login.html").ParseFiles("./web/templates/login.html")
if err!=nil{
	ErrorPage(w,http.StatusNotFound,"Not Found")
	return
}
login.Execute(w,nil)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	var user UserID
	user.Firstname = r.FormValue("first")
	user.Lastname = r.FormValue("last")
	user.Email = r.FormValue("email")
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	user.Contract = r.FormValue("mobile")
	user.Gender = r.FormValue("gender")
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(err)
	// if err != nil {
	// 	ErrorPage(w,http.StatusBadRequest,"Bad Request")
	// 	return
	// }
	
	// fmt.Println("this")
	err = RegisterUser("users.json", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	home := template.Must(template.ParseFiles("./web/templates/index.html"))
	// w.WriteHeader(http.StatusOK)

	home.Execute(w, nil)
}

// func HandleLogin(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var user UserID
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	authenticated, err := AuthenticateUsers("users.json", user.Username, user.Password)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if !authenticated {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintln(w, "Login successful")
// }

func RegisterUser(filename string, newUser UserID) error {
	mu.Lock()
	defer mu.Unlock()

	// Read existing users from file
	users, err := ReadUsers(filename)
	fmt.Println(err)
	if err != nil {
		return err
	}

	// Check if user already exists
	if UserExists(users, newUser.Username) {
		return errors.New("user already exists")
	}

	// Append new user to list
	users = append(users, newUser)

	// Write updated users to file
	return WriteUsersToFile(filename, users)
}

// func AuthenticateUsers(filename, username, password string) (bool, error) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return false, err
// 	}
// 	defer file.Close()

// 	var users []UserID
// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(&users)
// 	if err != nil && err != io.EOF {
// 		return false, err
// 	}

// 	for _, user := range users {
// 		if user.Username == username && user.Password == password {
// 			return true, nil
// 		}
// 	}

// 	return false, nil
// }
