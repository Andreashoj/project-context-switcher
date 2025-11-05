package main

import (
	"fmt"
	"os"
	"project-context-switcher/cmd"
	"project-context-switcher/internal/db"
)

func main() {
	DB, err := db.NewDB()
	if err != nil {
		fmt.Printf("starting the database failed: %s", err)
		return
	}

	defer DB.Close()

	rootCmd := cmd.NewRootCmd(DB)
	rootCmd.Init()
	if err = rootCmd.Execute(); err != nil {
		fmt.Printf("something went wrong initializing the cli commands: %s\n", err)
		os.Exit(1)
	}
}
