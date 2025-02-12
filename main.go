package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	fileServe := http.FileServer(http.Dir("./"))
	mux.Handle("/app/", http.StripPrefix("/app", fileServe))

	mux.HandleFunc("/healthz", handlerHealthz)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Start server error: %w", err)
	}

	fmt.Sprintln("Server running...")

}

func handlerHealthz(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(http.StatusText(http.StatusOK)))
}
