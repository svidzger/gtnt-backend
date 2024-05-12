package db

import (
	"time"

	"github.com/svidzger/gtnt-backend/models"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Create a new user in the database
func CreateUser(user *models.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// Check for errors during hashing
	if err != nil {
		return err
	}
	// Prepare SQL statement
	query := `INSERT INTO users (username, email, password, created_at, updated_at, role, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	// Execute SQL statement
	err = db.QueryRow(query, user.Username, user.Email, string(hashedPassword), time.Now(), time.Now(), user.Role, user.IsActive).Scan(&user.ID)
	// Check for errors during execution of SQL statement
	if err != nil {
		return err
	}

	return nil
}

// Get a user by username from the database
func GetUser(username string) (*models.User, error) {
	// Create a new user object
	user := &models.User{}
	// Prepare SQL statement
	query := `SELECT id, username, email, password, created_at, updated_at, role, is_active FROM users WHERE username = $1`
	// Execute SQL statement
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Role, &user.IsActive)
	// Check for errors during execution of SQL statement
	if err != nil {
		return nil, err
	}
	// Return the user object
	return user, nil
}
