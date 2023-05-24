package v1

import "github.com/gin-gonic/gin"


func SetupGroup(group *gin.RouterGroup) {
  // add all of the appropriate routes
  group.GET("/hello", helloWorldHandler)
}

func helloWorldHandler(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "message": "Hello, world!",
    })
}
