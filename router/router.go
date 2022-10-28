package router

import (
	"final-project/handler"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	rout := gin.Default()

	user := rout.Group("/user")
	{
		user.POST("/register", handler.UserRegister)
		user.POST("/login", handler.UserLogin)
		user.PUT("/", middlewares.Authentication(), handler.UpdateUser)
		user.DELETE("/", middlewares.Authentication(), handler.DeleteUser)
	}

	photo := rout.Group("/photo")
	{
		photo.Use(middlewares.Authentication())
		photo.POST("/", handler.CreatePhoto)
		photo.GET("/", handler.GetPhoto)
		photo.PUT("/:photoId", middlewares.PhotoAuthorization(), handler.UpdatePhoto)
		photo.DELETE("/:photoId", middlewares.PhotoAuthorization(), handler.DeletePhoto)
	}

	comment := rout.Group("/comment")
	{
		comment.Use(middlewares.Authentication())
		comment.POST("/", handler.CreateComment)
		comment.GET("/", handler.GetComment)
		comment.PUT("/:commentId", middlewares.CommentAuthorization(), handler.UpdateComment)
		comment.DELETE("/:commentId", middlewares.CommentAuthorization(), handler.DeleteComment)
	}

	socialmedia := rout.Group("/socialmedia")
	{
		socialmedia.Use(middlewares.Authentication())
		socialmedia.POST("/", handler.CreateSocialMedia)
		socialmedia.GET("/", handler.GetSocialMedia)
		socialmedia.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), handler.UpdateSocialMedia)
		socialmedia.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), handler.DeleteSocialMedia)
	}

	return rout
}
