package middleware

import (
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/labstack/echo/v4"
)

func UserHandler(client clerk.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		sessClaims, ok := ctx.Value(clerk.ActiveSessionClaims).(*clerk.SessionClaims)
		if !ok {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		user, err := client.Users().Read(sessClaims.Subject)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, "Welcome "+*user.FirstName)
	}
}
