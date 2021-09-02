package api

import (
	"github.com/allyite/goapi3/handlers"
	"github.com/allyite/goapi3/helper"
	"net/http"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	var collection = helper.ConnectDB()
	// Route / to handler function
  /*
	e.GET("/health-check", handlers.HealthCheck)

	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)
  */
	
	handler1 := func(c echo.Context) error {
        	return handlers.GetStats(c, collection)
        }

	e.GET("/cats/:data", handler1)
  	e.GET("/dogs", func(c echo.Context) error {       
	return c.String(http.StatusOK, "Hello, Dogs!\n")  
	})

}



/*
func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
*/
