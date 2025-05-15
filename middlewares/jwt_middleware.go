package middlewares

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)

			claims := token.Claims.(jwt.MapClaims)
			userID := claims["sub"].(string)
			role := claims["role"].(string)

			c.Locals("userID", userID)
			c.Locals("role", role)

			return c.Next()
		},
	})
}
