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
		// If it's a directory-like request without a file extension, serve index.html
		if request.URL.Path == "/" || !strings.Contains(filepath.Base(request.URL.Path), ".") {
			http.ServeFile(writer, request, "internal/ui/dist/index.html")
		} else {
			staticFiles.ServeHTTP(writer, request)
		}
	}))

	// Background worker started
	go func() {
		fmt.Printf("Started dashboard on http://localhost:8080")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("something went wrong starting the http listener: %s", err)
		}
	}()

	// Create channel with type of os.Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // CTRL+C and kill <pid>
	<-quit                                               // Block and wait for the channel to receive a value of sigterm or sigint

	// Init empty context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Cancel context gracefully when main function stops running

	if err := srv.Shutdown(ctx); err != nil { // shutdown http server, and provide it with a context that cancels the shutdown after 5 seconds if it's stuck
		fmt.Printf("something went wrong stopping the http server")
		return err
	}

	return nil
}
