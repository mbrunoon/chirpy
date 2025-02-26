package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mbrunoon/chirpy/helpers"
)

func (app *App) UserCreate(res http.ResponseWriter, req *http.Request) {

	type reqParams struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	params := reqParams{}
	err := decoder.Decode(&params)

	if err != nil {
		helpers.ResponseError(res, http.StatusInternalServerError, "Error decoding params", err)
		return
	}

	user, err := app.Models.CreateUser(context.Background(), params.Email)

	if err != nil {
		helpers.ResponseError(res, http.StatusInternalServerError, "Error on CreateUser", err)
		return
	}

	userSerialized := helpers.UserSerializer(&user)

	helpers.ResponseJSON(res, 201, userSerialized)
}
