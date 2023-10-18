package user

//USER TABLE
type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
}

//USER AND FRIEND TABLE
type UserUserLink struct {
	UserID   int `gorm:"not null" json:"user_id"`
	FriendID int `gorm:"not null" json:"friend_id"`
}

type UserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type LoginRequest struct {
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type LoginResponse struct {
	accessToken string
	ID          int    `json:"id"`
	Email       string `json:"email"`
}

type Repository interface {
	GetUserById(id int) (User, error)
	GetUserByEmail(email string) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(user User) (User, error)
}

type Service interface {
	GetUserById(id int) (User, error)
	CreateUser(user UserRequest) (User, error)
	Login(request LoginRequest) (LoginResponse, error)
	UpdateUser(id int, user UserUpdateRequest) (User, error)
	DeleteUser(id int) (User, error)
}
