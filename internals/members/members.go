package members

import "time"

type Members struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID int    `gorm:"not null" json:"user_id"`
	Name   string `gorm:"not null" json:"name"`
}

type GroupMembers struct {
	MemberID  uint      `gorm:"primaryKey" json:"member_id"`
	GroupID   int       `gorm:"not null" json:"group_id"`
	Name      string    `gorm:"not null" json:"name"`
	Balance   int       `gorm:"not null" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type MemberHistory struct {
	MemberID       int    `gorm:"not null" json:"member_id"`
	UserID         int    `gorm:"not null" json:"user_id"`
	BalanceUpdated int    `gorm:"not null" json:"balance_updated"`
	Evidence       string `gorm:"not null" json:"evidence"`
	CreatedAt      time.Time
}
