package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"strconv"
	"hello/models"
	"hello/middlewares"
)

func ToggleBoolToggler() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		// パスパラメータはstringなのでuintに変換する
		intId, _ := strconv.Atoi(c.Param("id"))
		uintId := uint(intId)

		boolToggler := models.BoolToggler{}
		if dbs.DB.First(&boolToggler, uintId).RecordNotFound() {
			// idの指定が誤っている場合はステータスコード404を返す。
			return c.JSON(fasthttp.StatusNotFound, "指定したidのboolTogglerが見つかりませんでした。")
		} else {
			// bool値を反転させて保存する
			boolToggler.Toggler = !boolToggler.Toggler
			dbs.DB.Save(&boolToggler)
			return c.JSON(fasthttp.StatusOK, boolToggler)
		}
	}
}