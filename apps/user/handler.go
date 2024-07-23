package user

import (
	infrafiber "handarudwiki/mini-online-shop-go/infra/fiber"
	"handarudwiki/mini-online-shop-go/infra/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}
func (h handler) register(ctx *fiber.Ctx) error {
	req := RegisterRequestPayload{}

	err := ctx.BodyParser(&req)
	if err != nil {
		myError := response.ErrorBadRequest

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myError),
			infrafiber.WithHttpCode(http.StatusBadRequest),
			infrafiber.WithMessage("register fail"),
		).Send(ctx)
	}

	err = h.svc.register(ctx.Context(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}
	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("register success"),
	).Send(ctx)
}

func (h handler) login(ctx *fiber.Ctx) error {
	req := LoginRequestPayload{}
	err := ctx.BodyParser(&req)

	if err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithMessage("login fail"),
		).Send(ctx)
	}

	token, err := h.svc.login(ctx.Context(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(err.Error()),
		).Send(ctx)

	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("login success"),
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithPayload(map[string]interface{}{
			"acces_token": token,
		}),
	).Send(ctx)
}
