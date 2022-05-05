package postgres

import (
	"strateegy/project-service/controller/dto"
	"strateegy/project-service/database"
	models "strateegy/project-service/models/project"

	"github.com/lib/pq"
)

type ProjectRepository struct {
}

func (s *ProjectRepository) Store(p models.Project) (dto.ProjectResponseDTO, error) {
	db := database.GetDB()

	admins := pq.StringArray{p.Owner}

	project := models.Project{
		Title:       p.Title,
		Description: p.Description,
		Users:       pq.StringArray{},
		Admins:      admins,
		Nda_accept:  p.Nda_accept,
		Owner:       p.Owner,
	}

	db.Create(&project)

	return dto.ProjectResponseDTO{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		Users:       project.Users,
		Admins:      project.Admins,
		Owner:       project.Owner,
		Nda_accept:  project.Nda_accept,
		Link:        project.Link,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}, nil

}

func (s *ProjectRepository) FindByProjectId(ID int32) models.Project {
	db := database.GetDB()

	var project models.Project

	db.Find(&project, "id = ?", ID)

	return project

}

func (s *ProjectRepository) Save(p models.Project) models.Project {
	db := database.GetDB()

	db.Save(&p)

	return p
}
