package models

import (
    "errors"
)

// User represents a user entity
type User struct {
    ID    int
    Name  string
    Email string
}

var (
    users = []User{
        {ID: 1, Name: "Alice", Email: "alice@example.com"},
        {ID: 2, Name: "Bob", Email: "bob@example.com"},
    }
)

// GetUserByID retrieves a user by their ID
func GetUserByID(id int) (*User, error) {
    for _, u := range users {
        if u.ID == id {
            return &u, nil
        }
    }
    return nil, errors.New("user not found")
}

// GetAllUsers returns all users
func GetAllUsers() []User {
    return users
}

// AddUser adds a new user to the list
func AddUser(name, email string) {
    newID := len(users) + 1
    newUser := User{ID: newID, Name: name, Email: email}
    users = append(users, newUser)
}
