package main

import (
	"fmt"
	"log"
	"project-context-switcher/cmd"
	"project-context-switcher/internal/db"
	"project-context-switcher/internal/handlers"
	"project-context-switcher/internal/repos"
	"project-context-switcher/internal/server"
	"project-context-switcher/internal/services"

	"github.com/go-chi/chi"
)

func main() {
	DB, err := db.NewDB() // Consider if this should be removed from main, as it would run always
	if err != nil {
		fmt.Printf("starting the database failed: %s", err)
		return
	}

	defer DB.Close()

	r := chi.NewRouter()
	server.SetupCors(r)

	// repos
	projectRepo := repos.NewProjectRepo(DB)

	// services
	projectService := services.NewProjectService(projectRepo)

	// handler
	projectHandler := handlers.NewProjectHandler(projectService)
	handlers.RegisterProjectRoutes(r, projectHandler)

	server.RegisterWebServerRoute(r)

	// Put this into dev mode
	if err = server.RunWebServer(r); err != nil {
		log.Printf("failed starting the web server: %s", err)
		return
	}

	rootCmd := cmd.NewRootCmd(DB, r, projectRepo)
	rootCmd.Init()

	// Create some CRUD for creating projects?
	// When I have crud for projects, they can be served to frontend
	// Analyze projects and get data to showcase - this should enable the whole docker analyze implementation

	if err = rootCmd.Execute(); err != nil {
		fmt.Printf("something went wrong initializing the cli commands: %s\n", err)
		return
	}

}
