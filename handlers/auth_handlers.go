package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/svidzger/gtnt-backend/db"
	"github.com/svidzger/gtnt-backend/models"
)

// Set the CORS origirn, methods, and headers
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// setupResponse sets the necessary headers to enable CORS and handle preflight requests
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	enableCors(w)
	if (*req).Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}

// RegisterHandler handles the registration of a new user
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	setupResponse(&w, r)
	// Create a new user struct
	user := models.User{}
	// Decode the JSON body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	// Check if there was an error while decoding the JSON body
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Call the CreateUser function with the user struct to add the user to the database
	err = db.CreateUser(&user)
	// Check if there was an error while creating the user
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Return a success message
	fmt.Fprintln(w, "User registered successfully")
}

// LoginHandler is a handler function to authenticate a user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	setupResponse(&w, r)
	// Create a new user struct
	user := models.User{}
	// Decode the JSON body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get the user from the database
	dbUser, err := db.GetUser(user.Username)
	// Check if the user was not found
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	// JWT token variables. Key is the secret key used to sign the token,
	// t is the token, and s is the encoded token string
	var (
		key []byte
		t   *jwt.Token
		s   string
	)
	// Assign the secret key to a variable
	key = []byte(os.Getenv("JWT_SECRET_KEY"))
	// Create a new token with the username, role, and expiration time
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  dbUser.Username,
		"role": dbUser.Role,
		"exp":  time.Now().Add(time.Hour * 12).Unix(),
	})
	// Get the complete, signed token as a string and assign it to the variable s
	s, err = t.SignedString(key)
	// Check if there was an error while generating the token
	if err != nil {
		http.Error(w, "Error while generating token", http.StatusInternalServerError)
		return
	}
	// Write the token to the response
	w.Write([]byte(s))
}
