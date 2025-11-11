package handlers

import (
	"fmt"
	"log"
	"net/http"
	"project-context-switcher/internal/models"
	"project-context-switcher/internal/services"
	"strconv"

	"github.com/go-chi/chi"
)

type ProjectHandler struct {
	projectService services.ProjectService
}

func NewProjectHandler(projectService services.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func RegisterProjectRoutes(r *chi.Mux, handler *ProjectHandler) {
	r.Post("/api/project", handler.create)
	r.Get("/api/project", handler.getAll)
	r.Get("/api/project/{id}", handler.get)
}

func (h *ProjectHandler) create(w http.ResponseWriter, r *http.Request) {
	payload, err := tryDecodeJSON[models.CreateProjectRequest](r.Body)
	if err != nil {
		log.Printf("failed converting the payload: %s", err)
		respondError(w, fmt.Sprintf("Something went wrong creating the project: %s", err), 500)
		return
	}

	project, err := h.projectService.Create(payload.Name, payload.Path)
	if err != nil {
		log.Printf("failed creating project: %s", err)
		respondError(w, fmt.Sprintf("Something went wrong creating the project: %s", err), 500)
		return
	}

	respondJSON(w, project, 201)
}

func (h *ProjectHandler) getAll(w http.ResponseWriter, r *http.Request) {
	projects, err := h.projectService.GetAll()
	if err != nil {
		log.Printf("failed retrieving the projects: %s", err)
		respondError(w, fmt.Sprintf("Something went wrong retrieving the projects: %s", err), 500)
		return
	}

	respondJSON(w, projects, 200)
}

func (h *ProjectHandler) get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		log.Printf("received invalid id from url: %s", idParam)
		respondError(w, "Invalid request", 500)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("couldn't convert given param into int got error: %s, from param: %s", err, idParam)
		respondError(w, "Invalid request", 500)
		return
	}

	container, err := h.projectService.Get(uint(id))
	if err != nil {
		log.Printf("error'd while getting the project: %s", err)
		respondError(w, "Invalid request", 500)
		return
	}

	respondJSON(w, container, 200)
}
