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

		type ResToggler struct {
			BoolTogglerId uint
			Toggler       bool
		}

		type Response struct {
			UserId uint
			Name   string
			Favorites []ResToggler
		}

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		// パスパラメータをuintに変換する
		intUserId, _ := strconv.Atoi(c.Param("id"))
		uintUserId := uint(intUserId)
		
		// Preloadでuserのリレーションを有効化してselect
		user := models.User{}
		dbs.DB.Preload("Favorites").Find(&user, uintUserId)

		var resJSON Response
		resJSON.UserId = user.ID
		resJSON.Name   = user.Name
		for _, v := range user.Favorites {
			toggler := ResToggler{
				BoolTogglerId: v.ID,
				Toggler: v.Toggler,
			}
			resJSON.Favorites = append(resJSON.Favorites, toggler)
		}
		return c.JSON(fasthttp.StatusOK, resJSON)
	}
}
