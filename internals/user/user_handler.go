package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewUserHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	u, err := h.Service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var userRequest UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.Service.CreateUser(userRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if res.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "email already exist",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	authId, _ := c.Get("user_id")
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	if id != authId {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you dont have permission to update this user",
		})
		return
	}

	u, err := h.Service.GetUserById(id)
	if u.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "record not found",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var userRequest UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.Service.UpdateUser(id, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var user LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := h.Service.Login(user)
	if u.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("jwt", u.accessToken, 3600, "/", "localhost", false, true)

	res := LoginResponse{
		Email: u.Email,
		ID:    u.ID,
	}

	c.JSON(http.StatusOK, res)

}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	authId, _ := c.Get("user_id")
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	if id != authId {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you dont have permission to update this user",
		})
		return
	}

	u, err := h.Service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, u)
}

func (h *Handler) GetAllFriend(c *gin.Context) {
	authId, _ := c.Get("user_id")
	users, err := h.Service.GetAllFriend(authId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) AddFriend(c *gin.Context) {
	authId, _ := c.Get("user_id")
	emailParam := c.Query("email")

	u, err := h.Service.AddFriend(emailParam, authId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("success add %s as a friend", u.FirstName),
	})
}

func (h *Handler) DeleteFriend(c *gin.Context) {
	authId, _ := c.Get("user_id")
	emailParam := c.Query("email")

	_, err := h.Service.DeleteFriend(emailParam, authId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "success delete friend")
}
