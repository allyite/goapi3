package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/allyite/goapi3/models"

	"github.com/labstack/echo"
)

//http://localhost:8000/cats/json?name=arnold&type=fluffy
func GetStats(c echo.Context) error {
	sName := c.QueryParam("name")
	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your state name is : %s\n", sName))
	} else if dataType == "json" {
		state := models.State{
			Name: sName,
		}
		return c.JSON(http.StatusOK, state)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data type as Sting or JSON",
		})
	}

}
