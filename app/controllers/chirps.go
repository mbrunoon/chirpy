package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mbrunoon/chirpy/app/models"
	"github.com/mbrunoon/chirpy/helpers"
)

func (app *App) ChirpyList(res http.ResponseWriter, req *http.Request) {
	chirps, err := app.Models.GetChirps(context.Background())
	if err != nil {
		helpers.ResponseError(res, http.StatusInternalServerError, "error fetching chirps", err)
		return
	}

	chirpsSerialized := helpers.ChirpySerializerList(&chirps)
	helpers.ResponseJSON(res, http.StatusOK, chirpsSerialized)
}

func (app *App) ChirpyShow(res http.ResponseWriter, req *http.Request) {

	chirpyId := req.PathValue("chirpID")
	if chirpyId == "" {
		helpers.ResponseError(res, http.StatusBadRequest, "chirp id required", nil)
		return
	}

	chirpy, err := app.Models.GetChirp(context.Background(), uuid.MustParse(chirpyId))

	if err != nil {
		if err == sql.ErrNoRows {
			helpers.ResponseError(res, http.StatusNotFound, "chirp not found", err)
			return
		}

		helpers.ResponseError(res, http.StatusInternalServerError, "error fetching chirp", err)
		return
	}

	chirpySerialized := helpers.ChirpySerializer(&chirpy)
	helpers.ResponseJSON(res, http.StatusOK, chirpySerialized)
}

func (app *App) ChirpyCreate(res http.ResponseWriter, req *http.Request) {

	params := models.CreateChirpyParams{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)

	if err != nil {
		helpers.ResponseError(res, http.StatusInternalServerError, "error decoding params", err)
		return
	}

	err = helpers.ValidateChirpy(&params)
	if err != nil {
		helpers.ResponseError(res, http.StatusBadRequest, "chirp invalid", err)
		return
	}

	chirpy, err := app.Models.CreateChirpy(context.Background(), params)
	if err != nil {
		helpers.ResponseError(res, http.StatusInternalServerError, "error on create chirp", err)
		return
	}

	chirpySerialized := helpers.ChirpySerializer(&chirpy)
	helpers.ResponseJSON(res, http.StatusCreated, chirpySerialized)
}
