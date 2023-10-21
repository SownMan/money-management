package main

import (
	"money-management/initializer"
	"money-management/internals/group"
	"money-management/internals/user"
	"money-management/router"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDatabase()
	initializer.SyncDB()
}

func main() {

	userRepo := user.NewUserRepository(initializer.DB)
	userSvc := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userSvc)

	groupRepo := group.NewGroupRepository(initializer.DB)
	groupSvc := group.NewGroupService(groupRepo)
	groupHandler := group.NewGroupHandler(groupSvc)

	router.InitRouter(userHandler, groupHandler)
	router.Start(":8081")
}
