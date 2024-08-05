package middleware

import (
	"net/http"
	"os"
	db "todo/pkg/database"
	"todo/pkg/models"
	"todo/pkg/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

type AuthRequest struct {
	Token string `json:"token"`
}

func HandleGoogleAuth(c echo.Context) error {
	var req AuthRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Verify the Google token
	payload, err := idtoken.Validate(c.Request().Context(), req.Token, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	// map[at_hash:OXdDbsLLoNM2y_zxHaUVyw aud:616778546290-pk54dktkqqsno31b418bclt7lhfga8oq.apps.googleusercontent.com azp:616778546290-pk54dktkqqsno31b418bclt7lhfga8oq.apps.googleusercontent.com email:sarthakagrawal927@gmail.com email_verified:true exp:1.722788293e+09 family_name:Agrawal given_name:Sarthak iat:1.722784693e+09 iss:https://accounts.google.com name:Sarthak Agrawal picture:https://lh3.googleusercontent.com/a/ACg8ocK4NgzJQmR5BR0hh_HHxLfZdCxgCXkJVERkcfPrK1zePlndoXT7=s96-c sub:110994534743808702789]
	// fmt.Printf("%+v", payload.Claims)

	// Create a session for the user, can store name & photo as well
	userId, err := CreateUserWithEmailIfNotExists(payload.Claims["email"].(string))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	// Create a session token
	sessionToken, err := createSessionToken(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create session")
	}

	// Return the session token
	return c.JSON(http.StatusOK, map[string]string{
		"sessionToken": sessionToken,
	})
}

func createSessionToken(userID uint) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID

	// Sign the token with a secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateUserWithEmailIfNotExists(email string) (uint, error) {
	user := &models.User{Email: email}
	queryResult := db.DB_CONNECTION.GetDB().First(&user, "email = ?", email) // Modify the query to use the email field
	if queryResult.RowsAffected == 0 {
		user.Email = email
		queryResult = db.DB_CONNECTION.GetDB().Create(user)
	}
	if queryResult.Error != nil {
		return 0, queryResult.Error
	}
	return user.ID, nil
}

func AttachUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ignore /api/auth/google
		if c.Path() == "/api/auth/google" {
			return next(c)
		}

		cookie, err := c.Cookie("auth")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Missing auth cookie")
		}

		claims, err := utils.ParseJWT(cookie.Value)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		if claims["id"] == nil || claims["id"].(float64) == 0 {
			return c.JSON(http.StatusUnauthorized, "Missing email claim")
		}

		c.Set("user_id", uint(claims["id"].(float64)))

		return next(c)
	}
}
