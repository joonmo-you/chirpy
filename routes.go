package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) registerRoutes(mux *http.ServeMux) {
	cfg.registerUserRoutes(mux)
}

func (cfg *apiConfig) registerUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/users", func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Email string `json:"email"`
		}
		type User struct {
			ID        uuid.UUID `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Email     string    `json:"email"`
		}
		params := parameters{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&params)
		fmt.Println(params)
		if err != nil {
			fmt.Println(err)
			respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
			return
		}
		user, err := cfg.db.CreateUser(r.Context(), params.Email)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
			return
		}
		respondWithJSON(w, http.StatusCreated, User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		})
	})
}
