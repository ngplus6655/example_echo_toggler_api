package routes

import (
	"hello/web/api"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
    {
			g.GET("/hello", api.ReturnHello())

			g.GET("/bool_toggler/:id", api.GetBoolToggler())
			g.POST("/bool_togglers", api.CreateBoolToggler())
			g.PUT("/bool_toggler/:id/toggle", api.ToggleBoolToggler())

			g.POST("/users", api.CreateUser())
			g.GET("/user/:id", api.ShowUserInfo())
			g.POST("/favorite/users/:user_id/bool_togglers/:toggler_id", 
			api.ToggleFavoriteToggler())
		}
}