package infrafiber

import (
	"context"
	"handarudwiki/mini-online-shop-go/infra/response"
	"handarudwiki/mini-online-shop-go/internal/config"
	internalLog "handarudwiki/mini-online-shop-go/internal/config/log"
	"handarudwiki/mini-online-shop-go/utility"
	"log"
	"strings"
	"time"

	"github.com/NooBeeID/go-logging/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Trace() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		now := time.Now()
		traceId := uuid.New()
		c.Set("X-Trace-Id", traceId.String())

		data := map[logger.LogKey]interface{}{
			logger.TRACER_ID: traceId,
			logger.METHOD:    c.Route().Method,
			logger.PATH:      string(c.Context().URI().Path()),
		}

		ctx = context.WithValue(ctx, logger.DATA, data)

		internalLog.Log.Infof(ctx, "incoming request")

		c.SetUserContext(ctx)
		err := c.Next()
		data[logger.RESPONSE_TIME] = time.Since(now).Milliseconds()
		data[logger.RESPONSE_TYPE] = "ms"
		httpStatusCode := c.Response().Header.StatusCode()

		if httpStatusCode >= 200 && httpStatusCode <= 299 {
			ctx = context.WithValue(ctx, logger.DATA, data)
			internalLog.Log.Infof(ctx, "success")
		} else {
			respBody := c.Response().Body()
			data["response_body"] = string(respBody)
			ctx := context.WithValue(ctx, logger.DATA, data)
			internalLog.Log.Errorf(ctx, "failed to")
		}

		return err
	}
}

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
