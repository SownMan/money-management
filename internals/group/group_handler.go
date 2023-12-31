package group

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewGroupHandler(s Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) CreateGroup(c *gin.Context) {
	userId, _ := c.Get("user_id")
	idInt := userId.(int)
	var groupReq GroupRequest
	if err := c.ShouldBindJSON(&groupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	g, err := h.Service.CreateGroup(idInt, groupReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, g)
}

func (h *Handler) UpdateGroup(c *gin.Context) {
	authId, _ := c.Get("user_id")
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	_, err := h.Service.FindLinkByUserAndGroup(id, authId.(int))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "no permission to edit this group"})
		return
	}

	var groupReq GroupUpdateRequest
	if err := c.ShouldBindJSON(&groupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	group, err := h.Service.UpdateGroup(id, groupReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": group,
	})
}

func (h *Handler) FindById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g, err := h.Service.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, g)
}

func (h *Handler) DeleteGroup(c *gin.Context) {
	authId, _ := c.Get("user_id")
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	_, err := h.Service.FindLinkByUserAndGroup(id, authId.(int))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "no permission to edit this group"})
		return
	}

	group, err := h.Service.DeleteGroup(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"group":   group,
		"message": "success deleting group",
	})
}
