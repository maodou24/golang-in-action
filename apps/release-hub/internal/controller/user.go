package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maodou24/release-hub/internal/middleware"
	req "github.com/maodou24/release-hub/internal/pkg/request"
	resp "github.com/maodou24/release-hub/internal/pkg/response"
	"github.com/maodou24/release-hub/internal/pkg/security"
	"github.com/maodou24/release-hub/internal/service"
)

var tokenCipher *security.TokenCipher
var sessionSecret string

func ConfigureAuth(cipher *security.TokenCipher, secret string) {
	tokenCipher = cipher
	sessionSecret = secret
}

func UpsertUserHandler(c *gin.Context) {
	var userReq req.User
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, resp.Failure(http.StatusBadRequest, err.Error()))
		return
	}
	user, err := service.UpsertUser(c.Request.Context(), userReq, tokenCipher)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Failure(http.StatusBadRequest, err.Error()))
		return
	}
	cookie := security.SignUserSession(user.PublicID, sessionSecret)
	http.SetCookie(c.Writer, &http.Cookie{
		Name: security.UserSessionCookie, Value: cookie, Path: "/", HttpOnly: true,
		SameSite: http.SameSiteLaxMode, MaxAge: 60 * 60 * 24 * 30,
	})
	c.JSON(http.StatusCreated, resp.Success(resp.AuthenticatedUser{Authenticated: true, User: *user}))
}

func CurrentUser(c *gin.Context) {
	publicID, ok := middleware.UserID(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp.Failure(http.StatusUnauthorized, "user session required"))
		return
	}
	user, err := service.GetUser(c.Request.Context(), publicID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, resp.Failure(http.StatusUnauthorized, err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(resp.AuthenticatedUser{Authenticated: true, User: *user}))
}
