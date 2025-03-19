package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mbrunoon/chirpy/app/models"
	"github.com/mbrunoon/chirpy/helpers"
	"github.com/mbrunoon/chirpy/internal/auth"
)

type reqParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *App) UserCreate(res http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	params := reqParams{}
	err := decoder.Decode(&params)

	if err != nil {
		helpers.ResponseError(res, http.StatusBadRequest, "Error decoding email parameter", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		helpers.ResponseError(res, http.StatusBadRequest, "Error hashing password", err)
		return
	}

	user, err := app.Models.CreateUser(req.Context(), models.CreateUserParams{
		Email:          params.Email,
		HashedPassword: hashedPassword,
	})

	if err != nil {
		helpers.ResponseError(res, http.StatusBadRequest, "Error creating user with email: "+params.Email, err)
		return
	}

	serializedUser := helpers.UserSerializer(&user)

	helpers.ResponseJSON(res, http.StatusCreated, serializedUser)
}
