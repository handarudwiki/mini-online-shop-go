package transaction

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

func (h handler) CreateTransaction(ctx *fiber.Ctx) error {
	req := CreateTransactionRequestPayload{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithHttpCode(http.StatusBadRequest),
			infrafiber.WithError(err),
		).Send(ctx)
	}

	userPublicID := ctx.Locals("PUBLIC_ID").(string)
	req.UserPublicID = userPublicID

	err = h.svc.CreateTransaction(ctx.Context(), req)
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
		infrafiber.WithMessage("Transaction created"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h handler) GetTransactions(ctx *fiber.Ctx) error {
	userPublicId := ctx.Locals("USER_PUBLIC_ID").(string)

	transactions, err := h.svc.TransactionHistories(ctx.Context(), userPublicId)
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

	response := []TransactionHistoryResponse{}

	for _, transaction := range transactions {
		response = append(response, transaction.ToTransactionHistoryResponse())
	}
	return infrafiber.NewResponse(
		infrafiber.WithMessage("Get Transaction Histoes Success"),
		infrafiber.WithPayload(response),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}
