package router

import (
	"github.com/gin-gonic/gin"
	c "github.com/maodou24/release-hub/internal/controller"
	"github.com/maodou24/release-hub/internal/middleware"
)

func New(sessionSecret string) *gin.Engine {
	r := gin.Default()

	// r.GET("/healthz", handler.Health)
	// r.GET("/:namespace/:projectName/:version/manifest.xml", projectHandler.GetManifest)

	api := r.Group("/api")
	api.Use(middleware.SessionAuth(sessionSecret))
	{
		api.POST("/users", c.UpsertUserHandler)
		api.GET("/users/me", c.CurrentUser)
		api.GET("/namespaces", c.QueryNamespacesHandler)
		api.DELETE("/namespaces/:namespaceId", c.DeleteNamespace)
		api.POST("/namespaces/:namespaceId/projects", c.CreateProject)
		api.GET("/namespaces/:namespaceId/projects", c.ListProjects)
		api.PUT("/projects/:projectId", c.UpdateProject)
		api.DELETE("/projects/:projectId", c.DeleteProject)
		api.GET("/projects/:projectId/releases", c.ListProjectVersions)
		api.POST("/projects/:projectId/releases", c.CreateRelease)
		api.GET("/releases/:releaseId", c.GetProjectVersion)
		api.PUT("/releases/:releaseId", c.UpdateProjectVersion)
		api.DELETE("/releases/:releaseId", c.DeleteProjectVersion)
		api.GET("/releases/:releaseId/manifest.xml", c.GetManifest)
	}

	return r
}
