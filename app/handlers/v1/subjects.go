package v1

import (
	// "github.com/JingusJohn/Community-BOSS-API/storage"
	"fmt"
	"strconv"
	"strings"

	"github.com/JingusJohn/Community-BOSS-API/app/storage"
	"github.com/JingusJohn/Community-BOSS-API/app/utilities"
	"github.com/gin-gonic/gin"
)

func GetSubjectsByQID(group *gin.RouterGroup) {
	group.GET("subjects/by-quarter-id/:quarter_id", func(ctx *gin.Context) {
		quarterId, err := strconv.Atoi(ctx.Param("quarter_id"))
		if err != nil || quarterId < 0 {
			ctx.JSON(400, &gin.H{
				"message": "quarter_id must be a positive integer",
				"err":     err,
			})
			return
		}
		// if quarterId < 0 {
		//   ctx.JSON(400, &gin.H{
		//     "message": "quarter_id"
		//   })
		// }
		rows, err := storage.BossGorm.Raw("select subject_id, name, quarter_id from subjects where quarter_id = ?", quarterId).Rows()
		utilities.HandleDBError(ctx, "subjects", err)
		subjects := []Subject{}
		for rows.Next() {
			var subject Subject
			storage.BossGorm.ScanRows(rows, &subject)
			subjects = append(subjects, subject)
		}

		ctx.JSON(200, subjects)
	})
}

func GetSubjectsBySeasonAndYear(group *gin.RouterGroup) {
	group.GET("subjects/:season/:year", func(ctx *gin.Context) {
		season := ctx.Param("season")
		ok := utilities.ValidateSeason(season)
		if !ok {
			ctx.JSON(400, &gin.H{
				"message": "Invalid 'season' provided. Check docs for valid seasons",
			})
			return
		}
		year, err := strconv.Atoi(ctx.Param("year"))
		if err != nil {
			// handle parsing error (invalid year provided)
			ctx.JSON(400, &gin.H{
				"message": "Invalid 'year' provided. Year must be an integer 2023-Present",
			})
			return
		}
		if year < 2023 && year >= 100 {
			ctx.JSON(400, &gin.H{
				"message": "Invalid 'year' provided. Year must be an integer 2023-Present",
			})
			return
		}
		if year < 100 {
			year += 2000
		}
		qRows, err := storage.BossGorm.Raw("select id, year, season, date_updated from quarters where year = ? AND UPPER(season) = ? order by date_updated desc", year, strings.ToUpper(season)).Rows()
		utilities.HandleDBError(ctx, "quarters", err)
		defer qRows.Close()

		quarters := []Quarter{}
		for qRows.Next() {
			// ScanRows scan a row into a quarter
			var quarter Quarter
			storage.BossGorm.ScanRows(qRows, &quarter)
			quarters = append(quarters, quarter)
		}

		// validate that we have at least one quarter
		if len(quarters) == 0 {
			ctx.JSON(400, &gin.H{
				"message": "No quarters found for the given season and year",
			})
			return
		}

		latestQuarter := quarters[0]
		fmt.Println(latestQuarter.Id)

		// search for courses matching this quarter ID

		sRows, err := storage.BossGorm.Raw("select subject_id, name, quarter_id from subjects where quarter_id = ?", latestQuarter.Id).Rows()
		utilities.HandleDBError(ctx, "subjects", err)
		defer sRows.Close()

		subjects := []Subject{}
		for sRows.Next() {
			// scan a row into a subject
			var subject Subject
			storage.BossGorm.ScanRows(sRows, &subject)
			subjects = append(subjects, subject)
		}

		ctx.JSON(200, subjects)
	})
}
