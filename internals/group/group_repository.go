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

func (r *groupRepository) FindLinkByGroupId(id int) (UserGroupLinks, error) {
	var userGroup UserGroupLinks
	err := r.db.Raw("SELECT * FROM user_group_links WHERE group_id = ?", id).First(&userGroup).Error
	return userGroup, err
}

func (r *groupRepository) CreateGroupLink(userGroup UserGroupLinks) (UserGroupLinks, error) {
	err := r.db.Create(&userGroup).Error
	return userGroup, err
}

func (r *groupRepository) DeleteGroupLink(userGroup UserGroupLinks) (UserGroupLinks, error) {
	err := r.db.Raw("DELETE FROM user_group_links WHERE group_id = ?", userGroup.GroupID).Scan(&userGroup).Error
	return userGroup, err
}
