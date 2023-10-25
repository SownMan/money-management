package initializer

import (
	"money-management/internals/group"
	"money-management/internals/members"
	"money-management/internals/user"
)

func SyncDB() {
	DB.AutoMigrate(&user.User{}, &user.UserUserLink{}, &group.Group{}, &group.UserGroupLinks{}, &members.Members{}, &members.GroupMembers{}, &members.MemberHistory{})
}
