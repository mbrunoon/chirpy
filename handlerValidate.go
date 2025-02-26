package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mbrunoon/chirpy/helpers"
)

func handlerValidateChirp(res http.ResponseWriter, req *http.Request) {

	type reqParams struct {
		Body string `json:"body"`
	}

	type resParams struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(req.Body)
	params := reqParams{}
	err := decoder.Decode(&params)

	if err != nil {
		helpers.ResponseError(res, http.StatusInternalServerError, "Error while decoding", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		helpers.ResponseError(res, http.StatusBadRequest, "Chirp max length is 140", nil)
		return
	}

	cleanedBody := cleanProfane(params.Body)

	helpers.ResponseJSON(res, 200, resParams{
		CleanedBody: cleanedBody,
	})

}

func cleanProfane(text string) string {

	profaneWords := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	replacer := "****"

	separator := " "
	splitedText := strings.Split(text, separator)

	for i, word := range splitedText {
		if _, exists := profaneWords[strings.ToLower(word)]; exists {
			splitedText[i] = replacer
		}
	}

	return strings.Join(splitedText, separator)
}
