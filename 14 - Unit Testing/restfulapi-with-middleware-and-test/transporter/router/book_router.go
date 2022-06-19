package router

import (
	"latian_clean_arch/app/middleware"
	"latian_clean_arch/repository/sql"
	"latian_clean_arch/service"
	"latian_clean_arch/transporter/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BookRouter(e *echo.Echo, db *gorm.DB) {

	repository := sql.NewBookRepositoryImplSql(db)
	service := service.NewBookServiceImpl(repository)
	controller := controller.NewBookController(service)

	e.GET("/books", controller.FindAll)
	e.POST("/books", controller.Create, middleware.IsAuthenticated)
	e.GET("/books/:id", controller.FindById)
	e.PUT("/books/:id", controller.Update, middleware.IsAuthenticated)
	e.DELETE("/books/:id", controller.Delete, middleware.IsAuthenticated)
}
