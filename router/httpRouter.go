package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thanawatpetchuen/ginstructure/handler"
)

func NewHTTPRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	v1.GET("/todos", h.TODO)

	return r
}
