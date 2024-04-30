package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/svidzger/gtnt-backend/db"
	"github.com/svidzger/gtnt-backend/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User registered successfully")
}
