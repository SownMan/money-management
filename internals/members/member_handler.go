package members

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewMemberHandler(service Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateMember(c *gin.Context) {
	authId, _ := c.Get("user_id")
	var memberReq MemberRequest
	if err := c.ShouldBindJSON(&memberReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	member, err := h.Service.CreateMember(memberReq, authId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, member)
}
