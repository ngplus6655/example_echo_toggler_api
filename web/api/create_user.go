package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"github.com/sirupsen/logrus"
	"hello/models"
	"hello/middlewares"
)

func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		// POSTされたnameから新しいレコードを作成
		u := new(models.User)
		if err := c.Bind(u); err != nil {
			logrus.Error(err)
			return err
		}
		dbs.DB.Create(u)

		return c.JSON(fasthttp.StatusOK, u)
	}
}