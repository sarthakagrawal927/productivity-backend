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
	User  struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Image string `json:"image"`
	} `json:"user"`
}

func HandleGoogleAuth(c echo.Context) error {
	var req AuthRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Verify the Google token
	_, err := idtoken.Validate(c.Request().Context(), req.Token, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	// Create a session for the user
	userId, err := CreateUserWithEmailIfNotExists(req.User.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	// Create a session token
	sessionToken, err := createSessionToken(userId, req.User.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create session")
	}

	// Return the session token
	return c.JSON(http.StatusOK, map[string]string{
		"sessionToken": sessionToken,
	})
}

func createSessionToken(userID uint, email string) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
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

		if claims["email"] == nil || claims["email"].(string) == "" {
			return c.JSON(http.StatusUnauthorized, "Missing email claim")
		}

		c.Set("user_id", claims["id"].(float64))
		c.Set("email", claims["email"].(string))

		return next(c)
	}
}
