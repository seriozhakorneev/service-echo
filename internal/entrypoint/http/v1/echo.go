package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"service-echo/internal/usecase"
	"service-echo/pkg/logger"
)

type EchoRoutes struct {
	t usecase.Echo
	l logger.Interface
}

func newEchoRoutes(handler *gin.RouterGroup, t usecase.Echo, l logger.Interface) {
	r := &EchoRoutes{t, l}

	h := handler.Group("/echo")
	{
		h.POST("/reflect", r.reflect)
	}
}

// TODO: DOCS
func (r *EchoRoutes) reflect(c *gin.Context) {
	var data json.RawMessage

	err := c.BindJSON(&data)
	if err != nil {
		r.l.Error(err, "http - v1 - reflect")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	fmt.Println(string(data))
	c.JSON(http.StatusOK, data)
}
