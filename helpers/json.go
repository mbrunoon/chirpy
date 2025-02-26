package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseError(res http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}

	if code >= 500 {
		log.Printf("Erro 5xx: %s", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	ResponseJSON(res, code, errorResponse{
		Error: msg,
	})
}

func ResponseJSON(res http.ResponseWriter, code int, payload interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	resData, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		res.WriteHeader(500)
		return
	}

	res.WriteHeader(code)
	res.Write(resData)
}
