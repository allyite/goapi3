package api

import (
	"github.com/allyite/goapi3/handlers"
	"net/http"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
  /*
	e.GET("/health-check", handlers.HealthCheck)

	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)
  */
	e.GET("/cats/:data", handlers.GetStats)
  	e.GET("/dogs", func(c echo.Context) error {       
	return c.String(http.StatusOK, "Hello, Dogs!\n")  
	})

}
/*
func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
*/
