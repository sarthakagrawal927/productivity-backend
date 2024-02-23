package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"golang.org/x/oauth2"
)

// Replace with your actual secret key and Google OAuth configuration
const (
	secretKey    = "your-secret-key"
	clientID     = "your-google-client-id"
	clientSecret = "your-google-client-secret"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main2() {
	r := mux.NewRouter()

	// Token handler
	r.HandleFunc("/auth/token", func(w http.ResponseWriter, r *http.Request) {
		// Validate authorization code and exchange for access token
		// ...

		// Validate access token and extract user information
		token, err := jwt.Parse(r.FormValue("access_token"), func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Use google-api-go-client to fetch user information
			config := &oauth2.Config{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
				RedirectURL:  "http://localhost:3000/api/auth/callback",
			}

			tok := &oauth2.Token{
				AccessToken: claims["access_token"].(string),
			}

			client := config.Client(r.Context(), tok)
			userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

			if err != nil {
				// Print the error for debugging
				fmt.Println(err)
				fmt.Println(userInfo)
				// Optionally: handle the error based on its type
				// (e.g., return an error response, redirect user)
			} else {
				// ... successful token validation and user information access ...
			}

		}

		// ... other routes and server configuration ...
	})
}
