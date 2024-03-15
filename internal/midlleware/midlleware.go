package midlleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io"
	"net/http"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}
type Uid struct {
	ID uint64 `json:"user_id"`
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		temp, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		cookie, err := r.Cookie("token")
		if err != nil {
			if err.Error() == http.ErrNoCookie.Error() {
				http.Error(w, "Token not found", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		tokenString := cookie.Value
		u := &Uid{}
		if err := json.Unmarshal(temp, u); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			if err.Error() == jwt.ErrSignatureInvalid.Error() {
				http.Error(w, "Invalid token signature", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			if claims.UserID != u.ID {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(temp))
		next.ServeHTTP(w, r)
	})
}
