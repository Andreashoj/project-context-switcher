package cmd

import (
	"fmt"
	"log"
	"project-context-switcher/internal/repos"

	"github.com/spf13/cobra"
)

func CreateProjectCommand(projectRepo repos.ProjectRepo) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "create a project, use flags --path and --name to specify project details",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return fmt.Errorf("failed to get name flag: %w", err)
			}
			path, err := cmd.Flags().GetString("path")
			if err != nil {
				return fmt.Errorf("failed to get path flag: %w", err)
			}

			project, err := projectRepo.Create(name, path)
			if err != nil {
				return fmt.Errorf("failed to create porject: %w", err)
			}

			fmt.Printf("Created project %s with path: %s\n", project.Name, project.Path)

			return nil
		},
	}

	cmd.Flags().StringP("path", "p", "", "path to project")
	cmd.Flags().StringP("name", "n", "", "name of project")

	err := cmd.MarkFlagRequired("path")
	err = cmd.MarkFlagRequired("name")

	if err != nil {
		log.Panicf("somehting went wrong! %w", err)
	}

	return cmd
}
