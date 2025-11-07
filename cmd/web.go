package cmd

import (
	"fmt"
	"project-context-switcher/internal/server"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
)

func CreateWebCommand(r *chi.Mux) *cobra.Command {
	return &cobra.Command{
		Use:  "web",
		Long: "Launches a web server, where the projects can be managed",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := server.RunWebServer(r)
			if err != nil {
				return fmt.Errorf("failed starting the webserver: %w", err)
			}

			return nil
		},
	}
}
