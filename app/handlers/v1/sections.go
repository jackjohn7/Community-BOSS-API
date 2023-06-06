package v1

import (
	"github.com/JingusJohn/Community-BOSS-API/app/storage"
	"github.com/JingusJohn/Community-BOSS-API/app/utilities"
	"github.com/gin-gonic/gin"
)

func GetSectionsByCID(group *gin.RouterGroup) {
	group.GET("/sections-by-cid/:cid", func(ctx *gin.Context) {
		cid := ctx.Param("cid")
		// validate that cid is a valid UUID
		ok := utilities.ValidateUUID(cid)
		if !ok {
			ctx.JSON(400, &gin.H{
				"message": "Invalid CourseID provided. Expected UUID.",
			})
			return
		}
		// query the database for all sections with the given cid
		rows, err := storage.BossGorm.Raw("select * from sections where course_id = ?", cid).Rows()
		if err != nil {
			// handle database error
			ctx.JSON(500, &gin.H{
				"message": "Unable to query database for sections with given course_id",
				"err":     err,
			})
			return
		}
		// create a slice of sections to store the results
		sections := []Section{}
		// iterate through the rows returned by the query
		for rows.Next() {
			// create a section to store the current row
			var section Section
			// scan the current row into the section
			storage.BossGorm.ScanRows(rows, &section)
			// append the section to the slice of sections
			sections = append(sections, section)
		}
		// validate that there are sections in the slice
		if len(sections) == 0 {
			ctx.JSON(404, &gin.H{
				"message": "No sections found with given course_id",
			})
			return
		}
		// return the slice of sections
		ctx.JSON(200, sections)
	})
}
