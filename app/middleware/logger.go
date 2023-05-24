package middleware

import (
  "time"
  "github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

/*
Logs incoming requests along with the time required to fulfill it
*/
func Logger() gin.HandlerFunc {
  return func(c *gin.Context) {
    // before request (before c.Next())
    start := time.Now()


    // executes handler function
    c.Next()
    // after request

    timeSince := time.Since(start)

    log.Print(c.Request.URL.Path)
    log.Print(timeSince)
  }
}
