package product

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

func (h handler) CreateProduct(ctx *fiber.Ctx) error {
	req := CreateProductRequestPayload{}

	err := ctx.BodyParser(&req)

	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	err = h.svc.CreateProduct(ctx.Context(), req)

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
		infrafiber.WithMessage("create product success"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h handler) GetListProduct(ctx *fiber.Ctx) error {
	req := ListProductRequestPayload{}

	err := ctx.QueryParser(&req)

	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("Invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	products, err := h.svc.ListProducts(ctx.Context(), req)

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

	productsResponse := NewProductListResponse(products)

	return infrafiber.NewResponse(
		infrafiber.WithQuery(productsResponse),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithQuery(req.GenerateDefaultValue()),
		infrafiber.WithMessage("get list product success"),
	).Send(ctx)

}

func (h handler) GetDetailProduct(ctx *fiber.Ctx) error {
	sku := ctx.Params("sku", "")

	if sku == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	product, err := h.svc.DetailProduct(ctx.Context(), sku)

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

	productResponse := ProductDetailResponse{
		ID:        product.ID,
		Name:      product.Name,
		SKU:       product.SKU,
		Stock:     int(product.Stock),
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.Updated_at,
	}

	return infrafiber.NewResponse(
		infrafiber.WithQuery(productResponse),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get detail product success"),
	).Send(ctx)
}
