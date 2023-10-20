package router

import (
	"money-management/internals/group"
	"money-management/internals/user"
	middlewre "money-management/middleware"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, groupHandler *group.Handler) {
	r = gin.Default()
	//user
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	r.PUT("/users/:id", middlewre.RequireAuth, userHandler.UpdateUser)

	//group
	r.GET("/groups/:id", middlewre.RequireAuth, groupHandler.FindById)
	r.POST("/groups", middlewre.RequireAuth, groupHandler.CreateGroup)
	r.PUT("/groups/:id", middlewre.RequireAuth, groupHandler.UpdateGroup)
	r.DELETE("/groups/:id", middlewre.RequireAuth, groupHandler.DeleteGroup)

}

func Start(addr string) error {
	return r.Run(addr)
}
