package controller

import (
	"strateegy/project-service/controller/dto"
	repository "strateegy/project-service/repositories/postgres"
	service "strateegy/project-service/service"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var dto dto.ProjectRequestDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})
		return
	}

	repo := &repository.ProjectRepository{}
	service := service.NewProjectService(repo)

	result, err := service.Store(dto, "1823940")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(201, result)
}

func AddUserInProject(c *gin.Context) {
	var dto dto.AddUserDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Bad Request",
		})
		return
	}

	repo := &repository.ProjectRepository{}
	service := service.NewProjectService(repo)

	ID := "10445"

	service.AddUserInProject(dto, ID)
}
