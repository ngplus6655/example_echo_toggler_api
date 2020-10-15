package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"hello/models"
	"hello/middlewares"
	"strconv"
)

func ShowUserInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		// パスパラメータをuintに変換する
		intUserId, _ := strconv.Atoi(c.Param("id"))
		uintUserId := uint(intUserId)
		
		user := models.User{}
		dbs.DB.First(&user, uintUserId)
		return c.JSON(fasthttp.StatusOK, user)
	}
}
