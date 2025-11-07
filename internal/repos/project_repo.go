package repos

import (
	"database/sql"
	"project-context-switcher/internal/models"
)

type ProjectRepo interface {
	Create(name string, path string) (*models.Project, error)
	Delete(id uint) error
	Update(name string, path string) (*models.Project, error)
}

type projectRepo struct {
	DB *sql.DB
}

func (p projectRepo) Create(name string, path string) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectRepo) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (p projectRepo) Update(name string, path string) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func NewProjectRepo(DB *sql.DB) ProjectRepo {
	return projectRepo{DB: DB}
}
