package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"hello/models"
	"hello/middlewares"
)

func CreateBoolToggler() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		boolToggler := models.BoolToggler{
			Toggler: false,
		}
		dbs.DB.Create(&boolToggler)
		return c.JSON(fasthttp.StatusOK, "BoolTogglerを追加しました")
	}
}