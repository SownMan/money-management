package members

import "time"

//Member TABLE
type Members struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID int    `gorm:"not null" json:"user_id"`
	Name   string `gorm:"not null" json:"name"`
}

type MemberRequest struct {
	Name string `json:"name" binding:"required"`
}

//Group Member Table
type GroupMembers struct {
	ID        uint      `gorm:"primaryKey" json:"member_id"`
	GroupID   int       `gorm:"not null" json:"group_id"`
	Name      string    `gorm:"not null" json:"name"`
	Balance   int       `gorm:"not null" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

//Member History
type MemberHistory struct {
	MemberID       int    `gorm:"not null" json:"member_id"`
	UserID         int    `gorm:"not null" json:"user_id"`
	BalanceUpdated int    `gorm:"not null" json:"balance_updated"`
	Evidence       string `gorm:"not null" json:"evidence"`
	CreatedAt      time.Time
}

type Repository interface {
	CreateMember(Members) (Members, error)
}

type Service interface {
	CreateMember(member MemberRequest, id int) (Members, error)
}
