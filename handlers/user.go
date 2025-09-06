package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raghavendrah25/RHGHospitality/types"
)

func HandleGetUser(c *gin.Context) {
	data := []types.User{
		{ID: "1", FirstName: "John", LastName: "Doe"},
		{ID: "2", FirstName: "Jane", LastName: "Smith"},
	}
	
	c.JSON(http.StatusOK, data)
}

func HandleGetUserByID(c *gin.Context) {
	userID := c.Params.ByName("id")
	data := types.User{
		ID:        userID,
		FirstName: "John",
		LastName:  "Doe",
	}

	c.JSON(http.StatusOK, data)
}
