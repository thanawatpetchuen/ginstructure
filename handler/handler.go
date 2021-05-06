package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanawatpetchuen/ginstructure/service"
)

type Handler struct {
	todoService *service.TodoService
}

func NewHandler(todoService *service.TodoService) *Handler {
	return &Handler{todoService: todoService}
}

func (h *Handler) TODO(c *gin.Context) {
	response := h.todoService.GetTodos()

	c.JSON(http.StatusOK, response)
}
