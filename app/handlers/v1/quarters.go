package v1

import (
	"fmt"
	"time"

	"github.com/JingusJohn/Community-BOSS-API/storage"
	"github.com/gin-gonic/gin"
)

var (
	seasonMap = map[string]int{
		"Winter": 0,
		"Spring": 1,
		"Summer": 2,
		"Fall":   3,
	}
)

func GetAllQuarters(group *gin.RouterGroup) {
	group.GET("/quarters-all", func(ctx *gin.Context) {
		rows, err := storage.BossGorm.Raw("select id, year, season, date_updated from quarters").Rows()
		if err != nil {
			// return some error
		}
		defer rows.Close()

		quarters := []Quarter{}
		for rows.Next() {
			// ScanRows scan a row into a quarter
			var quarter Quarter
			storage.BossGorm.ScanRows(rows, &quarter)
			quarters = append(quarters, quarter)
		}
		fmt.Println(quarters)
		ctx.JSON(200, quarters)
	})
}

func GetQuarters(group *gin.RouterGroup) {
	group.GET("/quarters", func(ctx *gin.Context) {
		latestYear := time.Now().Year()
		rows, err := storage.BossGorm.Raw("select id, year, season, date_updated from quarters where year >= ? order by date_updated desc", latestYear).Rows()
		if err != nil {
			// return some error
      ctx.JSON(500, err)
		}
		defer rows.Close()

		quarters := []Quarter{}
		for rows.Next() {
			// ScanRows scan a row into a quarter
			var quarter Quarter
			storage.BossGorm.ScanRows(rows, &quarter)
			quarters = append(quarters, quarter)
		}

    quarterMap := map[string]Quarter{}
    fmt.Println(quarterMap)
    for _, q := range quarters {
      key := fmt.Sprintf("%s %d", q.Season, q.Year)
      current, ok := quarterMap[key]
      if ok {
        // if in map, check if q is newer than current
        if q.DateUpdated.UnixMilli() > current.DateUpdated.UnixMilli() {
          quarterMap[key] = q
        }
      } else {
        quarterMap[key] = q
      }
    }
    latestQuarters := make([]Quarter, 0)
    fmt.Println(latestQuarters)
    for _, q := range quarterMap {
      fmt.Println(q)
      latestQuarters = append(latestQuarters, q)
    }
    fmt.Println(latestQuarters)
    ctx.JSON(200, latestQuarters)

	})
}

func GetLatestQuarter(group *gin.RouterGroup) {
	group.GET("/latest-quarter", func(ctx *gin.Context) {
		latestYear := time.Now().Year()
		rows, err := storage.BossGorm.Raw("select id, year, season, date_updated from quarters where year >= ? order by year desc, date_updated desc", latestYear).Rows()
		if err != nil {
			// return some error
		}
		defer rows.Close()

		quarters := []Quarter{}
		for rows.Next() {
			// ScanRows scan a row into a quarter
			var quarter Quarter
			storage.BossGorm.ScanRows(rows, &quarter)
			quarters = append(quarters, quarter)
		}

		// iterate through quarters
		// start by assuming it's the first in the slice
		latest := 0
		for i := 1; i < len(quarters); i++ {
			if (quarters[i].Year > quarters[latest].Year) || (quarters[i].Year == quarters[latest].Year && seasonMap[quarters[i].Season] > seasonMap[quarters[latest].Season]) {
				latest = i
			}
		}

		ctx.JSON(200, quarters[latest])
	})
}

