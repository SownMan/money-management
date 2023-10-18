package group

import (
	"gorm.io/gorm"
)

func NewGroupRepository(db *gorm.DB) Repository {
	return &groupRepository{db}
}

type groupRepository struct {
	db *gorm.DB
}

// CreateGroup implements Repository.
func (r *groupRepository) CreateGroup(group Group) (Group, error) {
	err := r.db.Create(&group).Error
	return group, err

}

func (r *groupRepository) CreateGroupLink(userGroup UserGroupLinks) (UserGroupLinks, error) {
	err := r.db.Create(&userGroup).Error
	return userGroup, err
}

func (r *groupRepository) UpdateGroup(group Group) (Group, error) {
	err := r.db.Save(&group).Error
	return group, err
}

// FindById implements Repository.
func (r *groupRepository) FindById(id int) (Group, error) {
	var g Group
	err := r.db.First(&g, id).Error
	return g, err
}

func (r *groupRepository) DeleteGroup(group Group) (Group, error) {
	err := r.db.Delete(&group).Error
	return group, err
}
