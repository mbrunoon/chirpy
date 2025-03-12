package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mbrunoon/chirpy/helpers"
)

type reqParams struct {
	Email string `json:"email"`
}

func (app *App) UserCreate(res http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	params := reqParams{}
	err := decoder.Decode(&params)

	if err != nil {
		helpers.ResponseError(res, http.StatusBadRequest, "Error decoding email parameter", err)
		return
	}

	user, err := app.Models.CreateUser(req.Context(), params.Email)

	if err != nil {
		helpers.ResponseError(res, http.StatusBadRequest, "Error creating user with email: "+params.Email, err)
		return
	}

	serializedUser := helpers.UserSerializer(&user)

	helpers.ResponseJSON(res, http.StatusCreated, serializedUser)
}
