package middleware

import (
	"log" // To record app information (Untuk mencatat informasi applikasi)
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Mencatat waktu awal request
		start := time.Now()

		// Meneruskan request ke handler berikutnya
		next.ServeHTTP(w, r)

		// Menghitung durasi request
		duration := time.Since(start)

		// Mencatat informasi request
		log.Printf(
			"%s %s - %v",
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}
