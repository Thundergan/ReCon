package config

import (
	"Remote/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	//Instantiates the SmartController
	remoteController := controllers.RemoteController{}

	//Creates a default router via gin
	r := gin.Default()

	//Adds the route group "/api/v1" to the router
	v1 := r.Group("/api/v1")
	{
		//Defines the GET route, full route is "/api/v1/exec"
		v1.GET("/exec/:host", remoteController.Exec)
	}

	//Returns the gin router a.k.a. Engine
	return r

}
