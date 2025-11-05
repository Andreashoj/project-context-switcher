package main

import (
	"fmt"
	"net/http"
	"project-context-switcher/cmd"
	"project-context-switcher/internal/db"

	"github.com/go-chi/chi"
)

func main() {
	DB, err := db.NewDB()
	if err != nil {
		fmt.Printf("starting the database failed: %s", err)
		return
	}

	defer DB.Close()

	r := chi.NewRouter()
	rootCmd := cmd.NewRootCmd(DB)
	rootCmd.Init()
	if err = rootCmd.Execute(); err != nil {
		fmt.Printf("something went wrong initializing the cli commands: %s\n", err)
		return
	}

	r.Route("/api", func(api chi.Router) {
		api.Get("/test", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "tester")
		})
	})

	// Serve frontend on specific port on command: pcs web
	// Create svelte project

	staticFiles := http.FileServer(http.Dir("./internal/ui/dist"))
	r.Handle("/*", staticFiles)
	if err = http.ListenAndServe("localhost:8080", r); err != nil {
		fmt.Printf("something went wrong starting the http listener: %s", err)
		return
	}
}
