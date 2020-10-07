package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"hello/models"
	"hello/middlewares"
	"strconv"
)

func GetBoolToggler() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		intId, _ := strconv.Atoi(c.Param("id"))
		uintId := uint(intId)

		boolToggler := models.BoolToggler{}
		if dbs.DB.First(&boolToggler, uintId).RecordNotFound() {
			return c.JSON(fasthttp.StatusNotFound, "指定したidのboolTogglerが見つかりませんでした。")
		} else {
			return c.JSON(fasthttp.StatusOK, boolToggler.Toggler)
		}
	}
}