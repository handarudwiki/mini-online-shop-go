package main

import (
	"handarudwiki/mini-online-shop-go/apps/product"
	"handarudwiki/mini-online-shop-go/apps/user"
	"handarudwiki/mini-online-shop-go/external/database"
	"handarudwiki/mini-online-shop-go/internal/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "config.yaml"

	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	user.Init(router, db)
	product.Init(router, db)
	router.Listen(config.Cfg.App.Port)
}
