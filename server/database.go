package Bitcoin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)


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
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(users)
}


func UserExists(users []UserID, username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}
