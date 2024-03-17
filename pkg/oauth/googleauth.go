package oauth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

func ValidateJWT(c echo.Context) error {
	// Extract the JWT token from the authorization header
	authHeader := c.Request().Header.Get("Authorization")

	if authHeader == "" {
		// Handle missing header error immediately
		return c.JSON(http.StatusUnauthorized, "Missing authorization header")
	}

	tokenString := strings.SplitN(authHeader, " ", 2)[1] // Extract token from "Bearer <token>" format

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return secret key (replace with your actual secret)
		return []byte("simple"), nil
	})

	fmt.Println(tokenString)
	fmt.Println(token)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, "Invalid signature")
		}
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error parsing JWT: %v", err))
	}

	// Extract and verify email from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.JSON(http.StatusUnauthorized, "Invalid claims or token")
	}

	if claims["email"] == nil || claims["email"].(string) == "" {
		return c.JSON(http.StatusUnauthorized, "Missing email claim")
	}

	// Access authenticated user's email from claims
	email := claims["email"].(string)

	// Proceed with actions for a valid user with email

	// Optionally, set a key-value pair in the context for later access
	c.Set("user", email)

	return c.String(http.StatusOK, "Welcome remeber Push through the Pain. Giving Up Hurts More, "+email+"!")
}
