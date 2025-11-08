package cmd

import (
	"database/sql"
	"fmt"
	"project-context-switcher/internal/repos"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	db          *sql.DB
	cmd         *cobra.Command
	Router      *chi.Mux
	projectRepo repos.ProjectRepo
}

func NewRootCmd(db *sql.DB, router *chi.Mux, projectRepo repos.ProjectRepo) RootCmd {
	return RootCmd{
		db:          db,
		Router:      router,
		cmd:         &cobra.Command{},
		projectRepo: projectRepo,
	}
}

func (r *RootCmd) Init() {
	r.cmd.AddCommand(CreateWebCommand(r.Router))
	r.cmd.AddCommand(CreateProjectCommand(r.projectRepo))
}

func (r *RootCmd) Execute() error {
	if err := r.cmd.Execute(); err != nil {
		return fmt.Errorf("something went wrong executing root command: %s", err)
	}

	return nil
}
