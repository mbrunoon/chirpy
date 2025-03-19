package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mbrunoon/chirpy/app/models"
	"github.com/mbrunoon/chirpy/helpers"
	"github.com/mbrunoon/chirpy/internal/auth"
)

func (app *App) Login(res http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		models.User
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		helpers.ResponseError(res, http.StatusUnauthorized, "Incorrect Login or Password", err)
		return
	}

	user, err := app.Models.GetUserByEmail(req.Context(), params.Email)
	if err != nil {
		helpers.ResponseError(res, http.StatusUnauthorized, "Incorrect Login or Password", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		helpers.ResponseError(res, http.StatusUnauthorized, "Incorrect Login or Password", err)
		return
	}

	helpers.ResponseJSON(res, http.StatusOK, response{
		User: models.User{
			ID:    user.ID,
			Email: user.Email,
		},
	})

}
