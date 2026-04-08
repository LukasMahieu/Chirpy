package main

import (
	"github.com/LukasMahieu/Chirpy/internal/database"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
}

func (cfg *apiConfig) middlewareMetricsIncs(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// middleware logic
			cfg.fileserverHits.Add(1)

			// transfer control back to handler
			next.ServeHTTP(w, r)
		})
}
