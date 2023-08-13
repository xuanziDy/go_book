package handler

import "github.com/gin-gonic/gin"

func (h *UserHandler) RegisterRouter(r *gin.Engine) {
	r.Group("/api/v1")
	{
		r.POST("/users/signUp", h.SignUp)
		r.POST("/users/login", h.Login)
		r.POST("/users/logout", h.Logout)
		r.POST("/users/edit", h.Edit)
		r.POST("/users/profile", h.Profile)
	}
}
