package service

import (
	"strateegy/project-service/controller/dto"
	interfaces "strateegy/project-service/interface"
	models "strateegy/project-service/models/project"
)

type ProjectService struct {
	repository interfaces.IProjectRepository
}

func NewProjectService(r interfaces.IProjectRepository) *ProjectService {
	return &ProjectService{
		repository: r,
	}
}

func (s *ProjectService) Store(data dto.ProjectRequestDTO, ID string) (dto.ProjectResponseDTO, error) {

	project := models.Project{
		Title:       data.Title,
		Description: data.Description,
		Nda_accept:  data.Nda_accept,
		Owner:       ID,
	}

	result, err := s.repository.Store(project)
	if err != nil {
		return dto.ProjectResponseDTO{}, err
	}

	return result, nil
}

func (s *ProjectService) AddUserInProject(data dto.AddUserDTO, ID string) error {

	result := s.repository.FindByProjectId(data.ProjectId)

	newUserList := append(result.Users, data.UserId)

	result.Users = newUserList

	s.repository.Save(result)

	return nil
}
