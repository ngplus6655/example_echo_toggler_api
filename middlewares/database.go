package middlewares

import (
	"hello/database"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type DatabaseClient struct {
	DB *gorm.DB
}

// アクションから接続、操作ができるようにecho.Contextに登録
func DatabaseService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := database.Connect()
			d := DatabaseClient{DB: session}
			defer d.DB.Close()
			// sql文をロギング
			d.DB.LogMode(true)

			c.Set("dbs", &d)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}