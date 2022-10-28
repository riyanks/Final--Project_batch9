package controllers

import (
	"encoding/json"
	"final-project/config"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	photoRequest := models.CreatePhoto{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&photoRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&photoRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	photo := models.Photo{
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		PhotoUrl: photoRequest.PhotoUrl,
		UserId:   userID,
	}

	err := db.Debug().Create(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	_ = db.First(&photo, photo.ID).Error

	photoString, _ := json.Marshal(photo)
	photoResponse := models.CreatePhotoResponse{}
	json.Unmarshal(photoString, &photoResponse)

	c.JSON(http.StatusCreated, photoResponse)
}

func GetPhoto(c *gin.Context) {
	db := config.GetDB()

	photos := []models.Photo{}

	err := db.Debug().Preload("User").Order("id asc").Find(&photos).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	photosString, _ := json.Marshal(photos)
	photosResponse := []models.PhotoResponse{}
	json.Unmarshal(photosString, &photosResponse)

	c.JSON(http.StatusOK, photosResponse)
}

func UpdatePhoto(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	photoRequest := models.UpdatePhoto{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&photoRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&photoRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	photo := models.Photo{}
	photo.ID = uint(photoId)
	photo.UserId = userID

	updateString, _ := json.Marshal(photoRequest)
	updateData := models.Photo{}
	json.Unmarshal(updateString, &updateData)

	err := db.Model(&photo).Updates(updateData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	photoString, _ := json.Marshal(photo)
	photoResponse := models.UpdatePhotoResponse{}
	json.Unmarshal(photoString, &photoResponse)

	c.JSON(http.StatusOK, photoResponse)
}

func DeletePhoto(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	photo := models.Photo{}
	photo.ID = uint(photoId)
	photo.UserId = userID

	err := db.Delete(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo Anda berhasil dihapus",
	})
}
