package repos

import (
	"database/sql"
	"fmt"
	"project-context-switcher/internal/models"
)

type ProjectRepo interface {
	Create(name string, path string) (*models.Project, error)
	Delete(id uint) error
	Update(name string, path string) (*models.Project, error)
	GetAll() ([]models.Project, error)
}

type projectRepo struct {
	DB *sql.DB
}

func NewProjectRepo(DB *sql.DB) ProjectRepo {
	return projectRepo{DB: DB}
}

func (p projectRepo) GetAll() ([]models.Project, error) {
	rows, err := p.DB.Query(`SELECT id, name, path FROM projects`)
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err = rows.Scan(&project.Id, &project.Name, &project.Path)
		if err != nil {
			return nil, fmt.Errorf("failed mapping project entity: %w", err)
		}

		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed iterating through projects: %w", err)
	}

	return projects, nil
}

func (p projectRepo) Create(name string, path string) (*models.Project, error) {
	var project models.Project
	err := p.DB.
		QueryRow(`INSERT INTO projects (path, name) VALUES ($1, $2) RETURNING id, path, name, updated_at, created_at`, path, name).
		Scan(&project.Id, &project.Path, &project.Name, &project.UpdatedAt, &project.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return &project, nil
}

func (p projectRepo) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (p projectRepo) Update(name string, path string) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}
