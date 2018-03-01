package main

import (
	"net/http"

	"github.com/intelligentbee/echo/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/hello/:name", func(e echo.Context) error {
		name := e.Param("name")
		return e.String(http.StatusOK, "Hello "+name)
	})

	e.POST("/users", controllers.CreateUser)

	e.GET("/users/:id", controllers.GetUser)

	e.DELETE("/users/:id", controllers.DeleteUser)

	// Server
	e.Logger.Fatal(e.Start(":1234"))
}
