package cmd

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	db     *sql.DB
	cmd    *cobra.Command
	Router *chi.Mux
}

func NewRootCmd(db *sql.DB, router *chi.Mux) RootCmd {
	return RootCmd{
		db:     db,
		Router: router,
		cmd:    &cobra.Command{},
	}
}

func (r *RootCmd) Init() {
	r.cmd.AddCommand(CreateWebCommand(r.Router))
}

func (r *RootCmd) Execute() error {
	if err := r.cmd.Execute(); err != nil {
		return fmt.Errorf("something went wrong executing root command: %s", err)
	}

	return nil
}
