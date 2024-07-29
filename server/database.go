package Bitcoin

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}


func readUsers(filePath string) ([]User, error) {
    var users []User
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    data, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, &users)
    if err != nil {
        return nil, err
    }

    return users, nil
}

func writeUsers(filePath string, users []User) error {
    data, err := json.MarshalIndent(users, "", "  ")
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(filePath, data, 0644)
    if err != nil {
        return err
    }

    return nil
}



func addUser(filePath string, newUser User) error {
    users, err := readUsers(filePath)
    if err != nil {
        return err
    }

    users = append(users, newUser)
    return writeUsers(filePath, users)
}
