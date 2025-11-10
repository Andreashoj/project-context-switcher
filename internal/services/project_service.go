package services

import (
	"fmt"
	"os"
	"project-context-switcher/internal/models"
	"project-context-switcher/internal/repos"

	"gopkg.in/yaml.v3"
)

type ProjectService interface {
	Create(name string, path string) (*models.Project, error)
	GetAll() ([]models.Project, error)
	GetContainers(path string) (*models.DockerCompose, error)
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
	// Check if docker-compose is present in the giving path, if it is - it should fail
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

	container, err := p.GetContainers(project.Path)
	if err != nil {
		return nil, fmt.Errorf("failed getting containers: %w", err)
	}

	return container, nil
}

func (p *projectService) GetContainers(path string) (*models.DockerCompose, error) {
	file, err := os.ReadFile("internal/services/testdata/docker-compose.yml")
	if err != nil {
		return nil, fmt.Errorf("failed reading file: %w", err)
	}

	var compose models.DockerCompose
	if err = yaml.Unmarshal(file, &compose); err != nil {
		fmt.Println("failed")
		return nil, fmt.Errorf("failed decoding the document: %w", err)
	}

	fmt.Println(compose)

	// Environment variables
	// Begin seeing if I can update the environment variables!
	// Post this to the frontend!

	return &compose, nil
}
