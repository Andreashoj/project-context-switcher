package handlers

import (
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
	r.Route("/api/project", func(p chi.Router) {
		p.Post("/", handler.create)
		p.Get("/", handler.getAll)
		p.Get("/:id", handler.get)
	})
}

func (h *ProjectHandler) create(w http.ResponseWriter, r *http.Request) {
	payload, err := tryDecodeJSON[models.CreateProjectRequest](r.Body)
	if err != nil {
		log.Printf("failed converting the payload: %s", err)
		respondError(w, "Something went wrong creating the project", 500)
		return
	}

	project, err := h.projectService.Create(payload.Name, payload.Path)
	if err != nil {
		log.Printf("failed creating project: %s", err)
		respondError(w, "Something went wrong creating the project", 500)
		return
	}

	respondJSON(w, project, 201)
}

func (h *ProjectHandler) getAll(w http.ResponseWriter, r *http.Request) {
	projects, err := h.projectService.GetAll()
	if err != nil {
		log.Printf("failed retrieving the projects: %s", err)
		respondError(w, "Something went wrong retrieving the projects", 500)
		return
	}

	respondJSON(w, projects, 200)
}

func (h *ProjectHandler) get(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam != "" {
		log.Printf("received invalid id from url")
		respondError(w, "Invalid request", 500)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("couldn't convert given param into int")
		respondError(w, "Invalid request", 500)
		return
	}

	_, err = h.projectService.Get(uint(id))
	if err != nil {
		log.Printf("error'd while getting the project: %s", err)
		respondError(w, "Invalid request", 500)
		return
	}

	respondJSON(w, "Got project", 201)
}
