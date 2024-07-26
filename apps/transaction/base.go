package transaction

import (
	infrafiber "handarudwiki/mini-online-shop-go/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	transactionRouter := router.Group("transactions")

	transactionRouter.Use(infrafiber.CheckAuth())

	transactionRouter.Post("/", handler.CreateTransaction)
	transactionRouter.Get("/user/histories", handler.GetTransactions)
}
