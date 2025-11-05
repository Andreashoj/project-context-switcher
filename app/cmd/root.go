package cmd

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

type RootCmd struct {
	db  *sql.DB
	cmd *cobra.Command
}

func NewRootCmd(db *sql.DB) RootCmd {
	return RootCmd{
		db:  db,
		cmd: &cobra.Command{},
	}
}

func (r *RootCmd) Init() {
	r.cmd.AddCommand(testCommand)
}

func (r *RootCmd) Execute() error {
	if err := r.cmd.Execute(); err != nil {
		return fmt.Errorf("something went wrong executing root command: %s", err)
	}

	return nil
}
