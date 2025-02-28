package security

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		reqToken := ctx.Request().Header.Get("Authorization")
		if reqToken == "" {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "missing authorization header"})
		}

		parts := strings.Split(reqToken, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization format"})
		}

		tokenStr := parts[1]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		}

		if role, ok := claims["role"].(string); !ok || role != "admin" {
			return ctx.JSON(http.StatusForbidden, map[string]string{"error": "access denied"})
		}
		return next(ctx)
	}
}
