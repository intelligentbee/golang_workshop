package controllers

import (
	"net/http"
	"strconv"

	"github.com/intelligentbee/echo/config"
	"github.com/intelligentbee/echo/models"
	"github.com/labstack/echo"
)

// create user - response json
// func CreateUser(c echo.Context) error {
// 	return c.JSON(http.StatusCreated, `{“id”: 123, “username”: “John”}`)
// }

// create user - save in database
func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	db := config.OpenDB()
	defer db.Close()

	db.Create(&user)

	return c.JSON(http.StatusCreated, user)
}

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user := models.User{}

	db := config.OpenDB()
	defer db.Close()

	db.Where(&models.User{ID: uint(id)}).First(&user)

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	db := config.OpenDB()
	defer db.Close()

	db.Where(&models.User{ID: uint(id)}).Delete(&models.User{})

	return c.NoContent(http.StatusOK)
}
