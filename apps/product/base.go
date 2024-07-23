package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRoute := router.Group("products")

	productRoute.Get("", handler.GetDetailProduct)
	productRoute.Get("/sku/:sku", handler.GetDetailProduct)
	productRoute.Post("", handler.CreateProduct)
}
