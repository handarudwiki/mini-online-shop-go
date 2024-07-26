package infrafiber

import (
	"handarudwiki/mini-online-shop-go/infra/response"
	"handarudwiki/mini-online-shop-go/internal/config"
	"handarudwiki/mini-online-shop-go/utility"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization")
		if authorization == "" {
			return NewResponse(
				WithError(response.ErrUnauthorized),
			).Send(ctx)
		}
		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("token invalid")
			return NewResponse(
				WithError(response.ErrUnauthorized),
			).Send(ctx)
		}

		token := bearer[1]

		publicID, role, err := utility.ValidateToken(token, config.Cfg.App.Encryption.JWTSecret)
		if err != nil {
			return NewResponse(
				WithError(response.ErrUnauthorized),
			).Send(ctx)
		}
		ctx.Locals("ROLE", role)
		ctx.Locals("PUBLIC_ID", publicID)

		return ctx.Next()
	}
}

func CheckRoles(authorizedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("ROLE").(string)

		isExist := false

		for _, authorizedRole := range authorizedRoles {
			if authorizedRole == role {
				isExist = true
			}
		}

		if !isExist {
			return NewResponse(
				WithError(response.ErrForbiddenAccess),
			).Send(c)
		}

		return c.Next()
	}
}
