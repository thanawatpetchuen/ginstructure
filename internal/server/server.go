package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(port int, h *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: h,
	}

	return srv
}
