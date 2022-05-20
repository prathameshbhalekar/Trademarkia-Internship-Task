package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/db"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/models"
	"gorm.io/gorm"
)

func FindInRange(c *gin.Context) {
	id_str := c.Query("id")
	k_str := c.Query("k")
	k, k_err := strconv.Atoi(k_str)
	id, id_err := strconv.Atoi(id_str)

	if id_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if k_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "distance should be number"})
		return
	}

	var user models.Users
	res := db.DB.Where("id = ?", id).First(&user)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var usersInRange []models.Users
	db.DB.Where("id != ? AND Location > ? AND Location < ?", id, user.Location-k, user.Location+k).Find(&usersInRange)

	c.JSON(http.StatusOK, gin.H{"data": usersInRange})

}

func FindUsersByName(c *gin.Context) {
	q := c.Query("q")
	pattern := []string{"%", q, "%"}
	var users []models.Users
	db.DB.Where("Name LIKE ?", strings.Join(pattern, "")).Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})

}
