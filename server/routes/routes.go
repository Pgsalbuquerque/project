package routes

import (
	"strateegy/project-service/controller"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		project := main.Group("project")
		{
			project.POST("/", controller.CreateProject)
			project.POST("/user", controller.AddUserInProject)
		}
	}

	return router
}
