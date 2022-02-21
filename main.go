package main

import (
	shopApplication "backend-ft/application/shop"
	"backend-ft/common"
	shopController "backend-ft/controller/shop"
	shopInterface "backend-ft/interfaces/shop"
	shopRepository "backend-ft/repository/shop"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"log"
)

const RestPort = ":9090"

func main() {

	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/fita")
	if err != nil {
		log.Fatal("could not connect to database", err.Error())
	}
	defer db.Close()

	r := echo.New()
	r.Validator = &common.CustomValidator{
		Validator: validator.New(),
	}

	sr := shopRepository.NewShopRepository(db)
	sa := shopApplication.NewShopApplication(db, sr)
	si := shopInterface.NewShopInterface(db, sa)
	shopController := shopController.NewShopController(si)
	r.POST("/checkout", shopController.Checkout)
	r.GET("/cart", shopController.GetCart)
	r.POST("/cart", shopController.UpsertCart)
	r.Start(RestPort)

}
