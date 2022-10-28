package handler

import (
	"encoding/json"
	"final-project/config"
	"final-project/entity"
	"final-project/helpers"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socialMediaRequest := entity.CreateSocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&socialMediaRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&socialMediaRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	socialMedia := entity.SocialMedia{
		Name:           socialMediaRequest.Name,
		SocialMediaUrl: socialMediaRequest.SocialMediaUrl,
		UserId:         userID,
	}

	err := db.Debug().Create(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMediaString, _ := json.Marshal(socialMedia)
	socialMediaResponse := entity.CreateSocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	c.JSON(http.StatusCreated, socialMediaResponse)
}

func GetSocialMedia(c *gin.Context) {
	db := config.GetDB()

	socialMedias := []entity.SocialMedia{}

	err := db.Debug().Preload("User").Order("id asc").Find(&socialMedias).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMediaString, _ := json.Marshal(socialMedias)
	socialMediaResponse := []entity.SocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	c.JSON(http.StatusOK, socialMediaResponse)
}

func UpdateSocialMedia(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socialMediaRequest := entity.UpdateSocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&socialMediaRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&socialMediaRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	socialMedia := entity.SocialMedia{}
	socialMedia.ID = uint(socialMediaId)
	socialMedia.UserId = userID

	updateString, _ := json.Marshal(socialMediaRequest)
	updateData := entity.SocialMedia{}
	json.Unmarshal(updateString, &updateData)

	err := db.Model(&socialMedia).Updates(updateData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	_ = db.First(&socialMedia, socialMedia.ID).Error

	socialMediaString, _ := json.Marshal(socialMedia)
	socialMediaResponse := entity.UpdateSocialMediaResponse{}
	json.Unmarshal(socialMediaString, &socialMediaResponse)

	c.JSON(http.StatusOK, socialMediaResponse)
}

func DeleteSocialMedia(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	socialMedia := entity.SocialMedia{}
	socialMedia.ID = uint(socialMediaId)
	socialMedia.UserId = userID

	err := db.Delete(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
