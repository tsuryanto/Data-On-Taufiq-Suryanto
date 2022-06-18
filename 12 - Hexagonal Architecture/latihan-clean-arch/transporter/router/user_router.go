package router

import (
	"latian_clean_arch/repository/sql"
	"latian_clean_arch/service"
	"latian_clean_arch/transporter/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRouter(e *echo.Echo, db *gorm.DB) {

	repository := sql.NewUserRepositoryImplSql(db)
	service := service.NewUserServiceImpl(repository)
	controller := controller.NewUserController(service)

	e.GET("/users", controller.FindAll)
	e.POST("/users", controller.Create)
	e.GET("/users/:id", controller.FindById)
	e.PUT("/users/:id", controller.Update)
	e.DELETE("/users/:id", controller.Delete)
}
