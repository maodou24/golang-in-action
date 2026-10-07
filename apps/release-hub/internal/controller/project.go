package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maodou24/release-hub/internal/middleware"
	"github.com/maodou24/release-hub/internal/pkg/request"
	resp "github.com/maodou24/release-hub/internal/pkg/response"
	"github.com/maodou24/release-hub/internal/service"
	"gorm.io/gorm"
)

func CreateProject(c *gin.Context) {
	if _, ok := middleware.UserID(c); !ok {
		writeError(c, http.StatusUnauthorized, "user session required")
		return
	}
	namespaceID, ok := routeID(c, "namespaceId")
	if !ok {
		return
	}
	var projectReq request.ProjectReq
	if err := c.ShouldBindJSON(&projectReq); err != nil {
		writeError(c, http.StatusBadRequest, err.Error())
		return
	}

	projectReq.NamespaceID = namespaceID
	project, err := service.CreateProject(c.Request.Context(), projectReq)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp.Success(project))
}

func UpdateProject(c *gin.Context) {
	if _, ok := middleware.UserID(c); !ok {
		writeError(c, http.StatusUnauthorized, "user session required")
		return
	}
	projectID, ok := routeID(c, "projectId")
	if !ok {
		return
	}
	var projectReq request.ProjectReq
	if err := c.ShouldBindJSON(&projectReq); err != nil {
		writeError(c, http.StatusBadRequest, err.Error())
		return
	}
	project, err := service.UpdateProject(c.Request.Context(), projectID, projectReq)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(project))
}

func ListProjects(c *gin.Context) {
	namespaceID, ok := routeID(c, "namespaceId")
	if !ok {
		return
	}
	projects, err := service.ListProjects(c.Request.Context(), namespaceID)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(gin.H{"items": projects}))
}

func GetProjectVersion(c *gin.Context) {
	releaseID, ok := routeID(c, "releaseId")
	if !ok {
		return
	}
	release, err := service.GetProjectVersion(c.Request.Context(), releaseID)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(release))
}

func ListProjectVersions(c *gin.Context) {
	projectID, ok := routeID(c, "projectId")
	if !ok {
		return
	}
	versions, err := service.ListProjectVersions(c.Request.Context(), projectID)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(gin.H{"items": versions}))
}

func UpdateProjectVersion(c *gin.Context) {
	if _, ok := middleware.UserID(c); !ok {
		writeError(c, http.StatusUnauthorized, "user session required")
		return
	}
	releaseID, ok := routeID(c, "releaseId")
	if !ok {
		return
	}
	var releaseReq request.ReleaseReq
	if err := c.ShouldBindJSON(&releaseReq); err != nil {
		writeError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := service.UpdateProjectVersion(c.Request.Context(), releaseID, releaseReq); err != nil {
		writeServiceError(c, err)
		return
	}
	release, err := service.GetProjectVersion(c.Request.Context(), releaseID)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(release))
}

func DeleteProjectVersion(c *gin.Context) {
	if _, ok := middleware.UserID(c); !ok {
		writeError(c, http.StatusUnauthorized, "user session required")
		return
	}
	releaseID, ok := routeID(c, "releaseId")
	if !ok {
		return
	}
	if err := service.DeleteProjectVersion(c.Request.Context(), releaseID); err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func DeleteProject(c *gin.Context) {
	if _, ok := middleware.UserID(c); !ok {
		writeError(c, http.StatusUnauthorized, "user session required")
		return
	}
	projectID, ok := routeID(c, "projectId")
	if !ok {
		return
	}
	if err := service.DeleteProject(c.Request.Context(), projectID); err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func DeleteNamespace(c *gin.Context) {
	if _, ok := middleware.UserID(c); !ok {
		writeError(c, http.StatusUnauthorized, "user session required")
		return
	}
	namespaceID, ok := routeID(c, "namespaceId")
	if !ok {
		return
	}
	if err := service.DeleteNamespace(c.Request.Context(), namespaceID); err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func GetManifest(c *gin.Context) {
	releaseID, ok := routeID(c, "releaseId")
	if !ok {
		return
	}
	content, err := service.BuildManifest(c.Request.Context(), releaseID, c.DefaultQuery("remoteType", "https"))
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.Data(http.StatusOK, "application/xml; charset=utf-8", content)
}

func routeID(c *gin.Context, name string) (uint, bool) {
	value, err := strconv.ParseUint(c.Param(name), 10, 64)
	if err != nil || value == 0 {
		writeError(c, http.StatusBadRequest, name+" must be a positive integer")
		return 0, false
	}
	return uint(value), true
}

func writeServiceError(c *gin.Context, err error) {
	status := http.StatusBadRequest
	if errors.Is(err, gorm.ErrRecordNotFound) {
		status = http.StatusNotFound
	}
	writeError(c, status, err.Error())
}

func writeError(c *gin.Context, status int, message string) {
	c.JSON(status, resp.Failure(status, message))
}
