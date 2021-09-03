package router

import (
	"github.com/allyite/goapi3/api"
	"github.com/allyite/goapi3/handlers"
//	"github.com/allyite/goapi3/helper"
	
	"net/http"
//	"api/middlewares"

	"github.com/labstack/echo"
)

func New(mClient *mongo.Client) *echo.Echo {
	
	// create a new echo instance
	e := echo.New()

	handler1 := func(c echo.Context) error {
        	return handlers.GetStats(c, mClient)
        }

	e.GET("/cats/:data", handler1)
	
	return e
	
	/*	
  	//create groups
	//adminGroup := e.Group("/admin")

	//set all middlewares
	//middlewares.SetMainMiddleWares(e)
	//middlewares.SetAdminMiddlewares(adminGroup)

	//set main routes
	MainGroup(e)

	//set groupRoutes
	//api.AdminGroup(adminGroup)
	e.GET("/health-check", handlers.HealthCheck)

	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)
	e.GET("/dogs", func(c echo.Context) error {       
	return c.String(http.StatusOK, "Hello, Dogs!\n")  
	})
  	*/
}

