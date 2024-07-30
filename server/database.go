package Bitcoin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// type User struct {
// 	Firstname string `json:"first"`
// 	Lastname  string `json:"last"`
// 	Email     string `json:"email"`
// 	Username  string `json:"username"`
// 	Password  string `json:"password"`
// 	Location  string `json:"location"`
// 	Contract  string `json:"mobile"`
// }

func ReadUsers(filePath string) ([]UserID, error) {
	file, err := os.Open(filePath)
	fmt.Println(err)
	// if err != nil {
	// 	return nil, err
	// }
	defer file.Close()

	var User []UserID
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&User)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return User, nil
}

func WriteUsersToFile(filePath string, users []UserID) error {
	fmt.Println(users[len(users)-1])
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(users)
}

// func RegisterUsers(filename string, newUser User) error {
// 	users, err := readUsers(filename)
// 	if err != nil {
// 		return err
// 	}

// 	// Check if user already exists
// 	for _, user := range users {
// 		if user.Username == newUser.Username {
// 			return fmt.Errorf("user already exists")
// 		}
// 	}

// 	// Hash the password before storing
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	newUser.Password = string(hashedPassword)

// 	// Append new user
// 	users = append(users, newUser)
// 	return WriteUsers(filename, users)
// }

func UserExists(users []UserID, username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}

// func AuthenticateUser(filename, username, password string) (bool, error) {
// 	users, err := readUsers(filename)
// 	if err != nil {
// 		return false, err
// 	}

// 	for _, user := range users {
// 		if user.Username == username {
// 			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 			if err != nil {
// 				return false, nil
// 			}
// 			return true, nil
// 		}
// 	}

// 	return false, nil
// }
