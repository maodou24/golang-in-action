package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maodou24/release-hub/internal/middleware"
	"github.com/maodou24/release-hub/internal/pkg/request"
	resp "github.com/maodou24/release-hub/internal/pkg/response"
	"github.com/maodou24/release-hub/internal/service"
)

func CreateRelease(c *gin.Context) {
	projectID, ok := routeID(c, "projectId")
	if !ok {
		return
	}
	var releaseReq request.ReleaseReq
	if err := c.ShouldBindJSON(&releaseReq); err != nil {
		c.JSON(http.StatusBadRequest, resp.Failure(http.StatusBadRequest, err.Error()))
		return
	}
	publicUserID, ok := middleware.UserID(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp.Failure(http.StatusUnauthorized, "user session required"))
		return
	}

	releaseReq.ProjectID = projectID
	release, err := service.CreateRelease(c.Request.Context(), publicUserID, releaseReq, tokenCipher)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	detail, err := service.GetProjectVersion(c.Request.Context(), release.ID)
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp.Success(detail))
}
