package user

import (
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
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

func (r *userRepository) GetUserLink(userId, friendId int) (UserUserLink, error) {
	var link UserUserLink
	err := r.db.Raw("SELECT * FROM user_user_links WHERE user_id = ? AND friend_id = ?", userId, friendId).Scan(&link).Error
	if err != nil {
		err = r.db.Raw("SELECT * FROM user_user_links WHERE user_id = ? AND friend_id = ?", friendId, userId).Scan(&link).Error
		if err != nil {
			return UserUserLink{}, err
		}
	}
	return link, err
}

func (r *userRepository) GetAllFriend(id int) ([]User, error) {
	var links []UserUserLink
	err := r.db.Raw("SELECT * FROM user_user_links WHERE user_id = ? OR friend_id = ?", id, id).Scan(&links).Error

	var users []User

	for _, v := range links {
		if v.UserID != id {
			var tempUser User
			r.db.First(&tempUser, v.UserID)
			users = append(users, tempUser)
		} else {
			var tempUser User
			r.db.First(&tempUser, v.FriendID)
			users = append(users, tempUser)
		}
	}

	return users, err
}

func (r *userRepository) AddFriend(link UserUserLink) (UserUserLink, error) {
	err := r.db.Create(&link).Error
	return link, err
}

func (r *userRepository) DeleteFriend(link UserUserLink) (UserUserLink, error) {
	err := r.db.Raw("DELETE FROM user_user_links WHERE user_id = ? AND friend_id = ?", link.UserID, link.FriendID).Scan(&link).Error
	return link, err
}
