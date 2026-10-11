package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maodou24/release-hub/internal/middleware"
	"github.com/maodou24/release-hub/internal/model"
	"github.com/maodou24/release-hub/internal/pkg/request"
	resp "github.com/maodou24/release-hub/internal/pkg/response"
	"github.com/maodou24/release-hub/internal/service"
	"gorm.io/gorm"
)

func CreateNamespace(c *gin.Context) {
	if _, ok := middleware.UserID(c); !ok {
		c.JSON(http.StatusUnauthorized, resp.Failure(http.StatusUnauthorized, "user session required"))
		return
	}
	var namespaceReq request.NamespaceReq
	if err := c.ShouldBindJSON(&namespaceReq); err != nil {
		c.JSON(http.StatusBadRequest, resp.Failure(http.StatusBadRequest, err.Error()))
		return
	}
	namespace, err := service.GetOrCreateNamespace(c.Request.Context(), namespaceReq.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Failure(http.StatusBadRequest, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, resp.Success(namespace))
}

func QueryNamespacesHandler(c *gin.Context) {
	namespaces, err := service.QueryNamespaces(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Failure(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp.Success(namespaces))
}

func UpdateNamespace(c *gin.Context) {
	if _, ok := middleware.UserID(c); !ok {
		writeError(c, http.StatusUnauthorized, "user session required")
		return
	}
	id, ok := parseNamespaceID(c)
	if !ok {
		return
	}
	var payload request.NamespaceReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		writeError(c, http.StatusBadRequest, err.Error())
		return
	}
	namespace, err := service.UpdateNamespace(c.Request.Context(), id, payload.Name)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(namespace))
}

func DeleteNamespace(c *gin.Context) {
	if _, ok := middleware.UserID(c); !ok {
		c.JSON(http.StatusUnauthorized, resp.Failure(http.StatusUnauthorized, "user session required"))
		return
	}
	namespaceID, ok := parseNamespaceID(c)
	if !ok {
		return
	}
	if err := service.DeleteNamespace(c.Request.Context(), namespaceID); err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, model.ErrNamespaceNotEmpty) {
			status = http.StatusConflict
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		c.JSON(status, resp.Failure(status, err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func parseNamespaceID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("namespaceId"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, resp.Failure(http.StatusBadRequest, "namespaceId must be a positive integer"))
		return 0, false
	}
	return uint(id), true
}
