package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

// GetUser berdasarkan ID
func GetUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	user := Users{}
	err := db.QueryRow("SELECT * FROM users WHERE id=?", params["id"]).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// CreateUser
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	username := r.FormValue("username")
	email := r.FormValue("email")

	result, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	user := Users{
		ID:       int(id),
		Username: username,
		Email:    email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser
func UpdateUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	id, _ := strconv.Atoi(params["id"])

	var user Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("UPDATE users SET username=?, email=? WHERE id=?", user.Username, user.Email, id)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser
func DeleteUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	id, _ := strconv.Atoi(params["id"])

	result, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
