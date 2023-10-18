package router

import (
	"money-management/internals/group"
	"money-management/internals/user"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, groupHandler *group.Handler) {
	r = gin.Default()
	//user
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	r.PUT("/users/:id", userHandler.UpdateUser)

	//group
	r.GET("/groups/:id", groupHandler.FindById)
	r.POST("/groups", groupHandler.CreateGroup)
	r.PUT("/groups/:id", groupHandler.UpdateGroup)
	r.DELETE("/groups/:id", groupHandler.DeleteGroup)

}

func Start(addr string) error {
	return r.Run(addr)
}
