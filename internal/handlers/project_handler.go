package handlers

import (
	"net/http"
	"project-context-switcher/internal/services"

	"github.com/go-chi/chi"
)

type ProjectHandler struct {
	projectService services.ProjectService
}

func NewProjectHandler(projectService services.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func RegisterProjectRoutes(r *chi.Mux, handler *ProjectHandler) {
	r.Route("/api/project", func(p chi.Router) {
		p.Get("/", handler.create)
		p.Get("/", handler.getAll)
	})
}

func (h *ProjectHandler) create(w http.ResponseWriter, r *http.Request) {

}

func (h *ProjectHandler) getAll(w http.ResponseWriter, r *http.Request) {
	projects, err := h.projectService.GetAll()
	if err != nil {
		respondError(w, "Something went wrong retrieving the projects", 500)
	}

	respondJSON(w, projects, 200)
}
