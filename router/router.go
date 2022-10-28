package router

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	rout := gin.Default()

	user := rout.Group("/user")
	{
		user.POST("/register", controllers.UserRegister)
		user.POST("/login", controllers.UserLogin)
		user.PUT("/", middlewares.Authentication(), controllers.UpdateUser)
		user.DELETE("/", middlewares.Authentication(), controllers.DeleteUser)
	}

	photo := rout.Group("/photo")
	{
		photo.Use(middlewares.Authentication())
		photo.POST("/", controllers.CreatePhoto)
		photo.GET("/", controllers.GetPhoto)
		photo.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photo.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	comment := rout.Group("/comment")
	{
		comment.Use(middlewares.Authentication())
		comment.POST("/", controllers.CreateComment)
		comment.GET("/", controllers.GetComment)
		comment.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		comment.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialmedia := rout.Group("/socialmedia")
	{
		socialmedia.Use(middlewares.Authentication())
		socialmedia.POST("/", controllers.CreateSocialMedia)
		socialmedia.GET("/", controllers.GetSocialMedia)
		socialmedia.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialmedia.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return rout
}
