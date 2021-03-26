package auth

import (
	"context"
	"fmt"
	"net/http"

	d "github.com/allanger/pomoday-backend/third_party/postgres"
	"github.com/allanger/pomoday-backend/tools/hasher"
	"github.com/allanger/pomoday-backend/tools/logger"
)

var (
	log = logger.NewLogger("auth-middleware").Entry
	UserID string
)

// BasicAuth implements a simple middleware handler for adding basic http auth to a route.
func BasicAuth(realm string, creds map[string]string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok {
				log.Error("Not ok")
				basicAuthFailed(w, realm)
				return
			}
			u := &structUser{
				Username: user,
			}
			log.Debug("Getting user")
			getUser(context.Background(), u)
			if err := hasher.ComparePasswords(u.Password, pass); err != nil {
				log.Error(err)
				basicAuthFailed(w, realm)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func basicAuthFailed(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}

func getUser(ctx context.Context, u *structUser) (*structUser, error) {
	var getUserRequest = "SELECT id, password FROM users WHERE username=$1"
	database := d.Pool()
	var pass string
	err := database.QueryRow(ctx, getUserRequest, u.Username).Scan(&UserID, &pass)
	u.Password = pass
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Refactor cycle imports
type structUser struct {
	UserID   string `json:"userID" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
