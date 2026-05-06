package auth

import (
	"log"
	"pec2/internal/db"
)

// Set stores a token → email mapping in the database.
func Set(token, email string) {
	_, err := db.DB.Exec("INSERT INTO sesiones (token, email) VALUES (?, ?)", token, email)
	if err != nil {
		log.Println("Error guardando sesión:", err)
	}
}

// Get retrieves the email associated with a token from the database.
func Get(token string) (string, bool) {
	var email string
	err := db.DB.QueryRow("SELECT email FROM sesiones WHERE token = ?", token).Scan(&email)
	if err != nil {
		return "", false
	}
	return email, true
}

// Delete removes a token from the database.
func Delete(token string) {
	_, err := db.DB.Exec("DELETE FROM sesiones WHERE token = ?", token)
	if err != nil {
		log.Println("Error eliminando sesión:", err)
	}
}
