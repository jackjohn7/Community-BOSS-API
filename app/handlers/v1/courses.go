package v1

import (
	// "github.com/JingusJohn/Community-BOSS-API/storage"
	"github.com/gin-gonic/gin"
)

func GetCoursesBySeasonAndYear(group *gin.RouterGroup) {
	group.GET("/courses/:season/:year", func(ctx *gin.Context) {

	})
}
