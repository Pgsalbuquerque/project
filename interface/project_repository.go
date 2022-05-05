package interfaces

import (
	"strateegy/project-service/controller/dto"
	models "strateegy/project-service/models/project"
)

type IProjectRepository interface {
	Store(p models.Project) (dto.ProjectResponseDTO, error)
	FindByProjectId(ID int32) models.Project
	Save(p models.Project) models.Project
}
