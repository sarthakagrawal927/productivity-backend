package middleware

import (
	"net/http"
	"todo/pkg/utils"

	"github.com/labstack/echo/v4"
)

func printAllCookies(c echo.Context) {
	cookies := c.Cookies()
	for _, cookie := range cookies {
		print(cookie.Name)
		print(cookie.Value)
	}
}

// add a middleware to parse the JWT token in cookie, and return the email
// Process is the middleware function.
func AttachUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		printAllCookies(c)
		if err := next(c); err != nil {
			c.Error(err)
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

		// Access authenticated user's email from claims
		email := claims["email"].(string)
		userId := claims["ID"].(uint)

		print("User ID: ", userId)
		print("User Email: ", email)
		return nil
	}
}
