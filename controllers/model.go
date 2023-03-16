package controllers

// User adalah model untuk tabel users
type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
