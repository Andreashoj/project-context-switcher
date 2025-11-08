package services

import (
	"fmt"
	"project-context-switcher/internal/models"
	"project-context-switcher/internal/repos"
)

type ProjectService interface {
	Create(name string, path string) (*models.Project, error)
	GetAll() ([]models.Project, error)
}

type projectService struct {
	repo repos.ProjectRepo
}

func NewProjectService(repo repos.ProjectRepo) ProjectService {
	return &projectService{
		repo: repo,
	}
}

func (p *projectService) Create(name string, path string) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p *projectService) GetAll() ([]models.Project, error) {
	projects, err := p.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving projects")
	}

	return projects, nil
}
