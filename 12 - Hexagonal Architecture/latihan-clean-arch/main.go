package main

import (
	"latian_clean_arch/app/database"
	"latian_clean_arch/transporter/router"

	"github.com/labstack/echo/v4"
)

func main() {
	db := database.InitSQL()

	e := echo.New()
	router.UserRouter(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}
