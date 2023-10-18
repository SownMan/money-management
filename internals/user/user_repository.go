package user

import (
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) GetUserByEmail(email string) (User, error) {
	var u User
	err := r.db.First(&u, "email = ?", email).Error

	return u, err
}

func (r *userRepository) GetUserById(id int) (User, error) {
	var u User
	err := r.db.First(&u, id).Error
	return u, err
}

func (r *userRepository) DeleteUser(user User) (User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(user User) (User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
