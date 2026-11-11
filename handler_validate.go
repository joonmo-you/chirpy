package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

const MaxChirpLength = 140

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type payload struct {
		CleanedBody string `json:"cleaned_body"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	if len(params.Body) > MaxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	cleanedBody := []string{}
	for _, word := range strings.Split(params.Body, " ") {
		if slices.Contains([]string{"kerfuffle", "sharbert", "fornax"}, strings.ToLower(word)) {
			cleanedBody = append(cleanedBody, "****")
		} else {
			cleanedBody = append(cleanedBody, word)
		}
	}
	respondWithJSON(w, http.StatusOK, payload{CleanedBody: strings.Join(cleanedBody, " ")})
}
