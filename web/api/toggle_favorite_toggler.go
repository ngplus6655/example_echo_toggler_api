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
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		// パスパラメータをuintに変換する
		intUserId, _ := strconv.Atoi(c.Param("user_id"))
		uintUserId := uint(intUserId)
		intTogglerId, _ := strconv.Atoi(c.Param("toggler_id"))
		uintTogglerId := uint(intTogglerId)
		
		// 条件付きのsqlを生成する
		query := &models.Favorite{UserId: uintUserId, TogglerId: uintTogglerId,}
		// queryでデータが検索に引っかかった場合に値を保持する
		favorite := models.Favorite{}
		
		// まだお気に入りされていなかった場合は、新しいレコードを追加
		// お気に入りだった場合は、既存のレコードを削除
		if dbs.DB.Where(query).First(&favorite).RecordNotFound() {
			result := dbs.DB.Create(&query)
			return c.JSON(fasthttp.StatusOK, result)
		} else {
			// gormではデフォルトで論理削除(レコードを削除するのではなく
			// deleted_atというカラムにデータで削除を表現)を行うが、値をトグルする
			// データにおいては不要だと考えUnscopedをDeleteにつけている。
			dbs.DB.Unscoped().Delete(&favorite)
			return c.JSON(fasthttp.StatusOK, "お気に入りを解除しました。")
		}
	}
}
