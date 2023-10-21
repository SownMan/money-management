package group

import (
	"time"
)

// GROUP TABLE
type Group struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	Description    string    `gorm:"not null" json:"description"`
	BalanceTarget  int       `gorm:"not null" json:"balance_target"`
	DueDate        time.Time `gorm:"not null" json:"due_date"`
	Cover          string    `json:"cover"`
	MemberCapacity int       `json:"member_capacity"`
	AdminCapacity  int       `json:"admin_capacity"`
	TotalBalance   int       `gorm:"not null" json:"total_balance"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// USER_GROUP_LINKS TABLE
type UserGroupLinks struct {
	UserID  int    `gorm:"not null" json:"user_id"`
	GroupID int    `gorm:"not null" json:"group_id"`
	Role    string `gorm:"not null" json:"role"`
}

type GroupRequest struct {
	Name           string    `json:"name" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	BalanceTarget  int       `json:"balance_target" binding:"required"`
	DueDate        time.Time `json:"due_date"`
	Cover          string    `json:"cover"`
	MemberCapacity int       `json:"member_capacity"`
	AdminCapacity  int       `json:"admin_capacity"`
}

type GroupUpdateRequest struct {
	Name           string    `json:"name" `
	Description    string    `json:"description"`
	BalanceTarget  int       `json:"balance_target"`
	DueDate        time.Time `json:"due_date"`
	Cover          string    `json:"cover"`
	MemberCapacity int       `json:"member_capacity"`
	AdminCapacity  int       `json:"admin_capacity"`
}

type Repository interface {
	FindById(id int) (Group, error)
	CreateGroup(group Group) (Group, error)
	UpdateGroup(group Group) (Group, error)
	DeleteGroup(group Group) (Group, error)

	FindLinkByGroupId(id int) (UserGroupLinks, error)
	CreateGroupLink(userGroup UserGroupLinks) (UserGroupLinks, error)
	DeleteGroupLink(userGroup UserGroupLinks) (UserGroupLinks, error)
}

type Service interface {
	FindById(id int) (Group, error)
	UpdateGroup(id int, groupReq GroupUpdateRequest) (Group, error)
	CreateGroup(userId int, group GroupRequest) (Group, error)
	DeleteGroup(id int) (Group, error)
}
