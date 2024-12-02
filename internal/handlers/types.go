package handlers

import "medods/internal/repository"

type Response struct {
	Err     string `json:"error,omitempty"`
	Warning string `json:"warning,omitempty"`
	Success string `json:"success,omitempty"`
}

type Handler struct {
	rep *repository.Repository
}

func NewHandler(rep *repository.Repository) *Handler {
	return &Handler{
		rep: rep,
	}
}
