package members

import "gorm.io/gorm"

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepository {
	return &memberRepository{db}
}

func (r *memberRepository) CreateMember(member Members) (Members, error) {
	err := r.db.Create(&member).Error
	return member, err
}
