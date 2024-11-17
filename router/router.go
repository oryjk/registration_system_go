package router

import (
	activityService "awesomeProject/activity/service"
	"awesomeProject/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	apiGroup := r.Group("/go_api")
	{
		userGroup := apiGroup.Group("/users")
		{
			userGroup.GET("/:userId", func(c *gin.Context) {

				userId := c.Param("userId")
				userInfo, err := service.QueryUserById(userId)
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
					return
				}

				c.JSON(http.StatusOK, userInfo)
			})

			userGroup.GET("/all", func(c *gin.Context) {
				userInfos, err := service.QueryAllUserInfo()
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
					return
				}

				c.JSON(http.StatusOK, userInfos)

			})

			userGroup.GET("/managers", func(c *gin.Context) {
				userInfos, err := service.FindManager()
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
					return
				}

				c.JSON(http.StatusOK, userInfos)
			})
		}

		activityGroup := apiGroup.Group("/activity")
		{
			activityGroup.POST("/create", activityService.CreateActivity)
			activityGroup.GET("/all", activityService.QueryAllActivity)
			activityGroup.GET("/:id", activityService.QueryActivityById)
		}
	}
}
