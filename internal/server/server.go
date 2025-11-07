package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

func RunWebServer(r *chi.Mux) error {
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	staticFiles := http.FileServer(http.Dir("internal/ui/dist"))
	r.Handle("/*", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		directoryPath := strings.Contains(filepath.Base(request.URL.Path), ".")
		if directoryPath {
			staticFiles.ServeHTTP(writer, request)
		} else {
			http.ServeFile(writer, request, "internal/ui/dist/index.html")
		}
	}))

	go func() {
		fmt.Printf("Started dashboard on http://localhost:8080")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("something went wrong starting the http listener: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // CTRL+C and kill <pid>
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed stopping the http server: %w", err)
	}

	return nil
}
