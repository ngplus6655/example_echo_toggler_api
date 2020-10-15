package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"strconv"
	"hello/models"
	"hello/middlewares"
)

func ToggleFavoriteToggler() echo.HandlerFunc {
	return func(c echo.Context) error {

		// response用のJSONを作る構造体
		type Response struct {
			UserId         uint
			BoolTogglerId  uint
			Favorite       bool
		}

		// DB接続
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		
		// パスパラメータをuintに変換する
		intUserId, _ := strconv.Atoi(c.Param("user_id"))
		uintUserId := uint(intUserId)
		intTogglerId, _ := strconv.Atoi(c.Param("toggler_id"))
		uintTogglerId := uint(intTogglerId)

		// Response構造体をインスタンス化
		var resJSON Response
		resJSON.UserId = uintUserId
		resJSON.BoolTogglerId = uintTogglerId

		// AppendするときにIDを指定したboolTogglerを渡す
		var boolToggler models.BoolToggler
		boolToggler.ID = uintTogglerId
		
		// Preloadでuserのリレーションを有効化してselect
		user := &models.User{}
		dbs.DB.Preload("Favorites", "bool_toggler_id = ?", uintTogglerId).
		Find(&user, uintUserId)

		// まだお気に入りされていなかった場合は、新しいレコードを追加
		if len(user.Favorites) < 1 {
			dbs.DB.Model(&user).Association("Favorites").Append(&boolToggler)
			resJSON.Favorite = true

		// お気に入り済みだった場合は、既存のレコードを削除
		} else {
			dbs.DB.Model(&user).Association("Favorites").Delete(&boolToggler)
			resJSON.Favorite = false
		}
		return c.JSON(fasthttp.StatusOK, resJSON)
	}
}
