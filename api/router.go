package api

import (
	"net/http"
	"time"

	v1 "github.com/allanger/pomoday-backend/api/v1"
	l "github.com/allanger/pomoday-backend/middleware/logger"
	database "github.com/allanger/pomoday-backend/third_party/postgres"
	"github.com/allanger/pomoday-backend/tools/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	log = logger.NewLogger("router").Entry
	err error
	r   *chi.Mux
)

// TODO: Setup CORS
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		log.Printf("Should set headers")
		if r.Method == "OPTIONS" {
			log.Printf("Should return for OPTIONS")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// NewRouter configures router and endpoints
func NewRouter() error {
	if err = database.OpenConnectionPool(); err != nil {
		log.Fatal(err)
		return err
	}
	r = chi.NewRouter()
	r.Use(middleware.Timeout(15 * time.Second))
	r.Use(l.Logger("router", logger.Entry))
	r.Use(cors)
	r.Mount("/", v1.Router())
	return nil
}
