package main

// import postgres driver for its side effects
import _ "github.com/lib/pq"

import (
	"database/sql"
	"github.com/LukasMahieu/Chirpy/internal/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

func main() {
	// open db connection
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	Platform := os.Getenv("PLATFORM")
	log.Printf("PLATFORM = %q", Platform) // add this temporarily

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
		platform:       Platform,
	}

	// initialize server
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	dir := http.Dir(".")
	fh := http.FileServer(dir)
	mh := http.HandlerFunc(apiCfg.metricsHandler)
	rh := http.HandlerFunc(readinessHandler)
	reseth := http.HandlerFunc(apiCfg.resetHandler)
	chirpsh := http.HandlerFunc(apiCfg.postChirpsHandler)
	chirpsgeth := http.HandlerFunc(apiCfg.getChirpsHandler)
	chirpsidgeth := http.HandlerFunc(apiCfg.getChirpsIDHandler)
	usersh := http.HandlerFunc(apiCfg.usersHandler)
	loginh := http.HandlerFunc(apiCfg.loginHandler)

	// frontend
	mux.Handle("GET /app/", apiCfg.middlewareMetricsIncs(http.StripPrefix("/app/", fh)))

	// public api endpoints
	mux.Handle("GET /api/healthz", rh)
	mux.Handle("GET /api/chirps", chirpsgeth)
	mux.Handle("GET /api/chirps/{chirpID}", chirpsidgeth)
	mux.Handle("POST /api/chirps", chirpsh)
	mux.Handle("POST /api/users", usersh)
	mux.Handle("POST /api/login", loginh)

	// admin endpoints
	mux.Handle("GET /admin/metrics", mh)
	mux.Handle("POST /admin/reset", reseth)

	server.ListenAndServe()
}
