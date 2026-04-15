package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"user-crud/config"
	"user-crud/models"
)

// GET USERS
func GetUsers(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr != "" {
		id, _ := strconv.Atoi(idStr)

		var user models.User
		err := config.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", id).
			Scan(&user.ID, &user.Name)

		if err != nil {
			http.Error(w, "User not found", 404)
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}

	rows, _ := config.DB.Query("SELECT id, name FROM users")
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

// CREATE USER
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	result, _ := config.DB.Exec("INSERT INTO users(name) VALUES(?)", user.Name)
	id, _ := result.LastInsertId()

	user.ID = int(id)
	json.NewEncoder(w).Encode(user)
}

// UPDATE USER
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := config.DB.Exec("UPDATE users SET name=? WHERE id=?", user.Name, id)
	if err != nil {
		http.Error(w, "Update failed", 500)
		return
	}

	user.ID = id
	json.NewEncoder(w).Encode(user)
}

// DELETE USER
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	_, err := config.DB.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		http.Error(w, "Delete failed", 500)
		return
	}

	w.Write([]byte("Deleted"))
}
