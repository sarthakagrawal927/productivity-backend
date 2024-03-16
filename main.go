package main

import (
	/*
		"context"
		"fmt"
		"log"
		"os"
		"time"
	*/
	"fmt"
	"net/http"
	"strings"
	"todo/cmd/service"
	db "todo/pkg/database"

	//googlecalendar "todo/pkg/googlecalendar"

	"github.com/joho/godotenv"
	//"golang.org/x/oauth2/google"
	//"google.golang.org/api/option"
	//"google.golang.org/api/calendar/v3"
	jwt "github.com/golang-jwt/jwt/v5"
)

func validateJWT(w http.ResponseWriter, r *http.Request) {

	// Extract the JWT token from the authorization header
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := strings.SplitN(authHeader, " ", 2)[1] // Extract token from "Bearer <token>" format

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return your secret key (replace with your actual secret)
		return []byte("simple"), nil
	})

	fmt.Println(tokenString)
	fmt.Println(token)

	// Handle any errors during parsing
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, "Invalid signature", http.StatusUnauthorized)
			return
		}
		http.Error(w, fmt.Sprintf("Error parsing JWT duttaani , yeh sahi musibat hai: %v", err), http.StatusBadRequest)
		return
	}

	// Extract and verify email from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Invalid claims or token", http.StatusUnauthorized)
		return
	}

	if claims["email"] == nil || claims["email"].(string) == "" {
		http.Error(w, "Missing email claim", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome, %s!\n", claims["email"].(string))
	/*
		// Extract the JWT token from the authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.SplitN(authHeader, " ", 2)[1] // Extract token from "Bearer <token>" format

		// Create a new JWT parser with your secret key
		parser := jwt.NewParser(jwt.SigningMethodHS256)

		// Validate the token using parser
		claims := &jwt.StandardClaims{}
		token, err := parser.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Replace with your secure secret
			return []byte("zB0hWWCEcG5lqu2LDQx4FbWyRwEE8PNQgpvGqCD+no0="), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Invalid signature", http.StatusUnauthorized)
			} else {
				http.Error(w, fmt.Sprintf("Error parsing JWT: %v", err), http.StatusBadRequest)
			}
			return
		}

		// Extract and verify email from claims (optional)
		if claims.Subject != "" {
			fmt.Fprintf(w, "Welcome, %s!\n", claims.Subject) // Replace with appropriate response
		} else {
			http.Error(w, "Missing email claim", http.StatusUnauthorized)
			return
		}
	*/
}

func main() {
	/*
		ctx := context.Background()
		b, err := os.ReadFile("credentials.json")
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		// If modifying these scopes, delete your previously saved token.json.
		config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		client := googlecalendar.GetClient(config)

		srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Unable to retrieve Calendar client: %v", err)
		}

		t := time.Now().Format(time.RFC3339)
		events, err := srv.Events.List("primary").ShowDeleted(false).
			SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		}
		fmt.Println("Upcoming events:")
		if len(events.Items) == 0 {
			fmt.Println("No upcoming events found.")
		} else {
			for _, item := range events.Items {
				date := item.Start.DateTime
				if date == "" {
					date = item.Start.Date
				}
				fmt.Printf("%v (%v)\n", item.Summary, date)
			}
		}
	*/
	http.HandleFunc("/validate-jwt", validateJWT)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
	service.CreateService()

}

func init() {
	godotenv.Load()
	db.SetupDBConnection()
}
