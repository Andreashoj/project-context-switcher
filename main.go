package main

import (
	"fmt"
	"project-context-switcher/cmd"
	"project-context-switcher/internal/db"

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

	rootCmd := cmd.NewRootCmd(DB, r)
	rootCmd.Init()

	if err = rootCmd.Execute(); err != nil {
		fmt.Printf("something went wrong initializing the cli commands: %s\n", err)
		return
	}
}
