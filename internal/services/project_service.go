package services

import (
	"fmt"
	"project-context-switcher/internal/docker"
	"project-context-switcher/internal/models"
	"project-context-switcher/internal/repos"
)

type ProjectService interface {
	Create(name string, path string) (*models.Project, error)
	GetAll() ([]models.Project, error)
	Get(id uint) (*models.DockerCompose, error)
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
	_, err := docker.GetFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving docker-compose file: %w", err)
	}

	project, err := p.repo.Create(name, path)
	if err != nil {
		return nil, fmt.Errorf("failed creating the project: %w", err)
	}

	return project, nil
}

func (p *projectService) GetAll() ([]models.Project, error) {
	projects, err := p.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving projects")
	}

	return projects, nil
}

func (p *projectService) Get(id uint) (*models.DockerCompose, error) {
	project, err := p.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving project: %w", err)
	}

	file, err := docker.GetFile(project.Path)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving file: %w", err)
	}

	container, err := docker.GetContainers(file)
	if err != nil {
		return nil, fmt.Errorf("failed getting containers: %w", err)
	}

	return container, nil
}
