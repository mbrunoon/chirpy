package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/mbrunoon/chirpy/app/controllers"
	"github.com/mbrunoon/chirpy/app/models"
)

type ApiConfig struct {
	fileServerHits atomic.Int32
	models         *models.Queries
}

func (cfg *ApiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func main() {

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Error Postgres connection: %w", err)
	}

	dbQueries := models.New(db)

	apiCfg := ApiConfig{
		models: dbQueries,
	}

	app := controllers.App{
		Models: dbQueries,
	}

	mux := http.NewServeMux()

	fileServe := http.FileServer(http.Dir("./"))

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", fileServe)))

	mux.HandleFunc("GET /api/healthz", handlerHealthz)
	mux.HandleFunc("POST /api/users", app.UserCreate)
	mux.HandleFunc("POST /api/login", app.Login)

	mux.HandleFunc("GET /api/chirps", app.ChirpyList)
	mux.HandleFunc("GET /api/chirps/{chirpID}", app.ChirpyShow)
	mux.HandleFunc("POST /api/chirps", app.ChirpyCreate)

	mux.HandleFunc("POST /admin/reset", apiCfg.HandlerMetricsReset)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Start server error: %w", err)
	}

	fmt.Sprintln("Server running...")

}

func handlerHealthz(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(http.StatusText(http.StatusOK)))
}

func (cfg *ApiConfig) handlerMetrics(res http.ResponseWriter, req *http.Request) {
	hits := cfg.fileServerHits.Load()

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)

	htmlResponse := fmt.Sprintf(`
	<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
	</html>
	`, hits)

	res.Write([]byte(htmlResponse))
}

func (cfg *ApiConfig) HandlerMetricsReset(res http.ResponseWriter, req *http.Request) {
	cfg.fileServerHits.Swap(0)

	fmt.Println("PLATFORM:", os.Getenv("PLATFORM"))

	if os.Getenv("PLATFORM") == "development" {
		err := cfg.models.ResetUsers(context.Background())
		if err != nil {
			http.Error(res, "Failed to reset users", http.StatusInternalServerError)
			return
		}
	}

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Hits and users reset"))
}
