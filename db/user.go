package db

import (
	"time"

	"github.com/svidzger/gtnt-backend/models"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Create a new user
func CreateUser(user *models.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	// Prepare SQL statement
	query := `INSERT INTO users (username, email, password, created_at, updated_at, role, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	// Execute SQL statement
	err = db.QueryRow(query, user.Username, user.Email, string(hashedPassword), time.Now(), time.Now(), user.Role, user.IsActive).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetUser(id int) (models.User, error) {
	query := `SELECT id, username, email, created_at, updated_at, role, is_active FROM users WHERE id = $1`
	row := db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Role, &user.IsActive)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
