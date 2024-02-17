package middleware

// FOR THIS APPROACH TO WORK GCP IDENTITY PLATFORM SERVICE IS REQUIRED

import (
	"context"
	"fmt"
	"net/http"

	"github.com/openshift/oauth2-server/v4/config"
	"github.com/openshift/oauth2-server/v4/handlers"
	"github.com/openshift/oauth2-server/v4/server"
)

func main() {
	ctx := context.Background()

	// Replace with your Google OAuth credentials
	config := &config.ServerConfig{
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
		Issuer:       "http://localhost:8080",
	}

	provider := config.NewProvider("google").(*config.GoogleProvider)
	provider.Scopes = []string{"profile", "email"}

	server, err := server.NewServer(ctx, config)
	if err != nil {
		fmt.Println("Error creating server:", err)
		return
	}

	http.HandleFunc("/authorize", handlers.AuthorizeHandler(server))
	http.HandleFunc("/token", handlers.TokenHandler(server))

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// verify tokens in protected routes
// func protectedHandler(w http.ResponseWriter, r *http.Request) {
// 	token, err := server.AuthenticateToken(r, "Bearer")
// 	if err != nil {
// 		http.Error(w, "Invalid token", http.StatusUnauthorized)
// 		return
// 	}

// 	// Access user information from token and handle request
// }
