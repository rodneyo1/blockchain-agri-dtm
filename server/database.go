package Bitcoin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
    bcrypt "golang.org/x/crypto/bcrypt"
)

type User struct {
	Firstname string `json:"first"`
	Lastname  string `json:"last"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
    Location string `json:"location"`

	// Username string `json:"username"`
	// Password string `json:"password"`
}

func readUsers(filePath string) ([]User, error) {
	var users []User
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func WriteUsers(filePath string, users []User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0o644)
	if err != nil {
		return err
	}

	return nil
}


func RegisterUsers(filename string, newUser User) error {
	users, err := readUsers(filename)
	if err != nil {
		return err
	}

	// Check if user already exists
	for _, user := range users {
		if user.Username == newUser.Username {
			return fmt.Errorf("user already exists")
		}
	}

	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)

	// Append new user
	users = append(users, newUser)
	return WriteUsers(filename, users)
}

func AuthenticateUser(filename, username, password string) (bool, error) {
	users, err := readUsers(filename)
	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				return false, nil
			}
			return true, nil
		}
	}

	return false, nil
}
