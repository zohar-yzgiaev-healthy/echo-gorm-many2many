package handlers

import (
	"github.com/labstack/echo"
	database "gorm-many2many/database"
	"gorm-many2many/models"
	"net/http"
)

func GetPlace(c echo.Context) error {
	uid := c.QueryParam("uid")
	db := database.Get()
	place := models.Place{}
	if findError := db.Debug().Preload("Features").Preload("Photos").Where("id = ?", uid).Find(&place).Error; findError != nil {
		return c.JSON(http.StatusInternalServerError, "Unable to find place")
	}
	return c.JSON(http.StatusOK, place)
}
