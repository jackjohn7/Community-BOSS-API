package v1

import "github.com/gin-gonic/gin"

func SetupGroup(group *gin.RouterGroup) {
	// add all of the appropriate routes
	group.GET("/hello", helloWorldHandler)
	// Functions for getting information about quarters
	GetQuarters(group)
	GetAllQuarters(group)
	GetLatestQuarter(group)

	// Functions for getting information about subjects
	GetSubjectsBySeasonAndYear(group)
	GetSubjectsByQID(group)

	// Functions for getting information about courses
	GetCoursesBySYS(group)
	GetCoursesBySID(group)

	// Functions for getting information about sections
	GetSectionsByCID(group)
}

// @BasePath /beta

// PingExample godoc
// @Summary Hello, world!
// @Schemes
// @Description Verifies that the server is properly able to receive requests and serve responses
// @Tags HealthCheck
// @Produce json
// @Success 200 {string} Helloworld
// @Router /hello [get]
func helloWorldHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}
