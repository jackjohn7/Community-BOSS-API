package v1

import (
	// "github.com/JingusJohn/Community-BOSS-API/storage"
	"github.com/JingusJohn/Community-BOSS-API/storage"
	"github.com/JingusJohn/Community-BOSS-API/utilities"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCoursesBySID(group *gin.RouterGroup) {
  group.GET("/courses/by-subject-id/:subject_id", func(ctx *gin.Context) {
    subjectId := ctx.Param("subject_id")
    ok := utilities.ValidateUUID(subjectId)
    if !ok {
      ctx.JSON(400, &gin.H{
        "message": "Invalid value provided for subject id. Subject ID is a UUIDv4",
        "received": subjectId,
      })
    }

    rows, err := storage.BossGorm.Raw("select course_id, name, subject_id from courses where subject_id = ?", subjectId).Rows()
    utilities.HandleDBError(ctx, "courses", err)
    courses := []Course{}
    for rows.Next() {
      var course Course
      storage.BossGorm.ScanRows(rows, &course)
      courses = append(courses, course)
    }

    ctx.JSON(200, courses)
  })
}

func GetCoursesBySYS(group *gin.RouterGroup) {
	group.GET("/courses/:season/:year/:subject", func(ctx *gin.Context) {
		season := ctx.Param("season")
		ok := utilities.ValidateSeason(season)
		if !ok {
			ctx.JSON(400, &gin.H{
				"message": "Invalid 'season' provided. Check docs for valid seasons",
			})
		}
		year, err := strconv.Atoi(ctx.Param("year"))
		if err != nil {
			// handle parsing error (invalid year provided)
			ctx.JSON(400, &gin.H{
				"message": "Invalid 'year' provided. Year must be an integer 2023-Present",
			})
		}
		subject := ctx.Param("subject")
		qRows, err := storage.BossGorm.Raw("select id, year, season, date_updated from quarters where year = ? AND season = ? order by date_updated desc", year, season).Rows()
    utilities.HandleDBError(ctx, "quarters", err)
		defer qRows.Close()

		quarters := []Quarter{}
		for qRows.Next() {
			// ScanRows scan a row into a quarter
			var quarter Quarter
			storage.BossGorm.ScanRows(qRows, &quarter)
			quarters = append(quarters, quarter)
		}

		latestQuarter := quarters[0]

		// search for courses matching this quarter ID

		sRows, err := storage.BossGorm.Raw("select subject_id, name, quarter_id from subjects where quarter_id = ? and name = ?", latestQuarter.Id, subject).Rows()
    utilities.HandleDBError(ctx, "subjects", err)
    defer sRows.Close()

		subjects := []Subject{}
		for sRows.Next() {
			var subjectScanned Subject
			storage.BossGorm.ScanRows(sRows, &subjectScanned)
			subjects = append(subjects, subjectScanned)
		}
		// should only be one subject found
		if len(subjects) > 1 {
			ctx.JSON(500, &gin.H{
        "message": "Too many subjects found",
        "raw_err": err,
      })
		}
    if len(subjects) == 0 {
      ctx.JSON(400, &gin.H{
        "message": "Nonexistent subject",
      })
    }

		cRows, err := storage.BossGorm.Raw("select course_id, name, subject_id from courses where subject_id = ?", subjects[0].SubjectId).Rows()
    utilities.HandleDBError(ctx, "courses", err)
    defer cRows.Close()

    courses := []Course{}
    for cRows.Next() {
      var course Course
      storage.BossGorm.ScanRows(cRows, &course)
      courses = append(courses, course)
    }

    ctx.JSON(200, courses)
	})
}
