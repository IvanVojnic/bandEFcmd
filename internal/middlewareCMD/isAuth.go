package middlewareCMD

import (
	"cmdMS/internal/errorwrapper"
	"cmdMS/internal/utils"
	"cmdMS/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JwtAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if c.Path() == "/auth/createUser" || c.Path() == "/auth/signIn" || c.Path() == "/auth/refreshToken" {
				return next(c)
			}
			var tokens models.Tokens
			req := c.Request()
			headers := req.Header
			atHeader := headers.Get("Authorization")
			if atHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "no access token in header")
			}
			atHeaderArr := strings.Split(atHeader, " ")
			tokens.AccessToken = atHeaderArr[1]
			userID, errIsAuth := utils.IsAuthorized(tokens.AccessToken)
			if errIsAuth != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, errorwrapper.ErrorResponse{Message: errIsAuth.Error()})
			}
			c.Set("user_id", userID)
			return next(c)
		}
	}
}
