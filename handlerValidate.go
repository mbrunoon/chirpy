package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidateChirp(res http.ResponseWriter, req *http.Request) {

	type reqParams struct {
		Body string `json:"body"`
	}

	type resParams struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(req.Body)
	params := reqParams{}
	err := decoder.Decode(&params)

	if err != nil {
		responseError(res, http.StatusInternalServerError, "Error while decoding", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		responseError(res, http.StatusBadRequest, "Chirp max length is 140", nil)
		return
	}

	responseJSON(res, 200, resParams{
		Valid: true,
	})

}
