package handlers

import (
	"net/http"
	"github.com/todo20200404/models"
	"github.com/labstack/echo"
)

func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetTasks())
}
