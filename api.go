package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIServer struct {
	listenAddr string
	store      DB
}

type APIError struct {
	Error string
}

func NewAPIServer(listenAddr string, store DB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	mux := http.NewServeMux()
	godotenv.Load()

	// Posts
	mux.HandleFunc("GET /posts", makeHttpHandlerFunc(s.HandleGetPosts, false))
	mux.HandleFunc("POST /posts", makeHttpHandlerFunc(s.HandleCreatePost, true))

	// Categories
	mux.HandleFunc("GET /categories", makeHttpHandlerFunc(s.HandleGetCategories, false))
	mux.HandleFunc("POST /categories", makeHttpHandlerFunc(s.HandleCreateCategory, true))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	corsHandler := handlers.CORS(headersOk, originsOk, methodsOk)(mux)

	fmt.Println("Listening on port", s.listenAddr)
	http.ListenAndServe(":"+s.listenAddr, corsHandler)
}

func makeHttpHandlerFunc(f apiFunc, isProtected bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isProtected {
			secret := os.Getenv("SUPABASE_JWT_SECRET")
			token := r.Header.Get("Authorization")
			if token == "" {
				WriteJSON(w, http.StatusUnauthorized, APIError{Error: "Unauthorized"})
				return
			}

			_, err := parseJWT(token, []byte(secret))
			if err != nil {
				WriteJSON(w, http.StatusUnauthorized, APIError{Error: "Unauthorized"})

				return
			}
		}

		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func parseJWT(token string, hmacSecret []byte) (string, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	// Check if the token is valid
	if err != nil {
		return "", fmt.Errorf("error validating token: %v", err)
	} else if claims, ok := t.Claims.(*Claims); ok {
		return claims.Email, nil
	}

	return "", fmt.Errorf("error parsing token: %v", err)
}
