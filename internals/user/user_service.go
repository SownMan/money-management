package user

import (
	"money-management/util"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const ()

type userService struct {
	Repository
}

func NewUserService(repository Repository) Service {
	return &userService{repository}
}

func (s *userService) GetUserById(id int) (User, error) {
	u, err := s.Repository.GetUserById(id)
	return u, err
}

func (s *userService) CreateUser(user UserRequest) (User, error) {
	u, _ := s.Repository.GetUserByEmail(user.Email)
	if u.ID != 0 {
		return User{}, nil
	}
	//hash password
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	newUser := User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	}

	r, err := s.Repository.CreateUser(newUser)
	if err != nil {
		return User{}, err
	}
	res := User{
		ID:        r.ID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  r.Password,
	}

	return res, nil
}

type MyJWTClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.MapClaims
}

func (s *userService) Login(request LoginRequest) (LoginResponse, error) {
	//get user by email
	user, err := s.Repository.GetUserByEmail(request.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	//check hashed pasword
	err = util.CheckPassword(request.Password, user.Password)
	if err != nil {
		return LoginResponse{}, err
	}

	//jwt token & claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:    strconv.Itoa(int(user.ID)),
		Email: user.Email,
		MapClaims: jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	})
	ss, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{accessToken: ss, Email: user.Email, ID: int(user.ID)}, nil
}

func (s *userService) UpdateUser(id int, user UserUpdateRequest) (User, error) {
	u, err := s.Repository.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	u.FirstName = user.FirstName
	u.LastName = user.LastName

	updatedUser, err := s.Repository.UpdateUser(u)
	if err != nil {
		return User{}, err
	}
	return updatedUser, nil
}

func (s *userService) DeleteUser(id int) (User, error) {
	panic("uniple")
}
