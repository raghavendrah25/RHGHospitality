package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raghavendrah25/RHGHospitality/db"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Params.ByName("id")
	user, err := h.userStore.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
