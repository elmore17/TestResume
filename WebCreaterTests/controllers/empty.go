package controllers

import (
	"TestResume/WebCreaterTests/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindUser(c *gin.Context) {
	var user []models.User
	models.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: input.Name, Email: input.Email}
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUserId(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.UserUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func CreateBook(c *gin.Context) {
	var input models.InputBook
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.First(&user, input.UserId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, UserId: input.UserId}
	models.DB.Create(&book)

	var result models.Book
	if err := models.DB.Preload("User").First(&result, book.Id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
