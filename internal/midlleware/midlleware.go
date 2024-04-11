package midlleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io"
	"log"
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

func LoggerRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		log.Print("Request Headers: ")
		for name, values := range r.Header {
			for _, value := range values {
				log.Printf("%s: %s ", name, value)
			}
		}
		log.Println()
		if r.Body != nil {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Failed to read request body:", err)
			}
			defer r.Body.Close()
			log.Println("Request Body:", string(body))
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		next.ServeHTTP(w, r)
	})
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
			if errors.Is(err, http.ErrNoCookie) {
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
			if errors.Is(err, jwt.ErrSignatureInvalid) {
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
