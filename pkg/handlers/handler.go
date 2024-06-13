package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "go-webapp/pkg/models"
    "go-webapp/pkg/utils"

    "github.com/gorilla/mux"
)

// GetUserHandler retrieves a user by ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
        return
    }

    user, err := models.GetUserByID(userID)
    if err != nil {
        utils.RespondWithError(w, http.StatusNotFound, "User not found")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, user)
}

// GetAllUsersHandler retrieves all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
    users := models.GetAllUsers()

    utils.RespondWithJSON(w, http.StatusOK, users)
}

// AddUserHandler adds a new user
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
    var newUser models.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    models.AddUser(newUser.Name, newUser.Email)

    utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User added successfully"})
}
