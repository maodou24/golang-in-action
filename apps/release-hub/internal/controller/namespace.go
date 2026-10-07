package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	resp "github.com/maodou24/release-hub/internal/pkg/response"
	"github.com/maodou24/release-hub/internal/service"
)

func QueryNamespacesHandler(c *gin.Context) {
	namespaces, err := service.QueryNamespaces(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Failure(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp.Success(namespaces))
}
