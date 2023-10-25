package router

import (
	"money-management/internals/group"
	"money-management/internals/user"
	middlewre "money-management/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, groupHandler *group.Handler) {
	r = gin.Default()
	//user
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	r.GET("users", middlewre.RequireAuth, userHandler.GetUserByEmail)
	r.PUT("/users/:id", middlewre.RequireAuth, userHandler.UpdateUser)
	r.DELETE("users/:id", middlewre.RequireAuth, userHandler.DeleteUser)

	//friend
	r.GET("/users/friends", middlewre.RequireAuth, userHandler.GetAllFriend)
	r.POST("/users/friends", middlewre.RequireAuth, userHandler.AddFriend)
	r.DELETE("users/friends", middlewre.RequireAuth, userHandler.DeleteFriend)

	//group
	r.GET("/groups/:id", middlewre.RequireAuth, groupHandler.FindById)
	r.POST("/groups", middlewre.RequireAuth, groupHandler.CreateGroup)
	r.PUT("/groups/:id", middlewre.RequireAuth, groupHandler.UpdateGroup)
	r.DELETE("/groups/:id", middlewre.RequireAuth, groupHandler.DeleteGroup)

	// test
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test": time.Now(),
		})
	})

}

func Start(addr string) error {
	return r.Run(addr)
}
