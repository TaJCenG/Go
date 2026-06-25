package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"user-service/pkg/response"
)

// POST /users
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if u.Name == "" || u.Email == "" {
		response.Error(w, http.StatusBadRequest, "Name and Email required")
		return
	}
	if err := CreateUser(u); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, u)
}

// GET /users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := GetUsers()
	response.JSON(w, http.StatusOK, users)
}

// GET /users/{id}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)
	u, err := GetUser(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, u)
}

// PUT /users/{id}
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	u.ID = id
	if u.Name == "" || u.Email == "" {
		response.Error(w, http.StatusBadRequest, "Name and Email required")
		return
	}
	if err := UpdateUser(id, u); err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, u)
}

// DELETE /users/{id}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)
	if err := DeleteUser(id); err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}
