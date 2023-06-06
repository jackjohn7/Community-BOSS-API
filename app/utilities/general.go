/*
General utilities to assist in handling requests
*/
package utilities

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
Validates that a string matches the UUIDv4 format using
'github.com/google/uuid'
*/
func ValidateUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

/*
Handles errors in standard ways
*/
func HandleDBError(ctx *gin.Context, table string, err error) {
	if err != nil {
		message := fmt.Sprintf("Something went wrong fetching data from %s. Please inform adminis.", table)
		ctx.JSON(500, &gin.H{
			"message": message,
			"raw_err": err,
		})
	}
}
