package main

import (
	"log"
	"os"

	"net/http"

	v1 "github.com/JingusJohn/Community-BOSS-API/app/handlers/v1"
	"github.com/JingusJohn/Community-BOSS-API/app/middleware"
	"github.com/JingusJohn/Community-BOSS-API/app/storage"
  docs "github.com/JingusJohn/Community-BOSS-API/app/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	// zlog "github.com/rs/zerolog/log"
  ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
  swaggerFiles "github.com/swaggo/files" // swagger embed files
)

func SetupRouter() *gin.Engine {
	if os.Getenv("BOSS_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// zlog.Print("Hello, world")
	connectionString := os.Getenv("BOSS_DB_CONNECTION")

	storage.InitBossDataDB(connectionString)

	router := gin.Default()

	// by convention, use /beta for incomplete versions
  docs.SwaggerInfo.BasePath = "/beta"
	v1Group := router.Group("/beta")
	v1.SetupGroup(v1Group)

	// use bundled rate limiter (at least for now)
	// app.Use(limiter.New(limiter.Config {
	//   Max: 20,
	//   Expiration: 30 * time.Second,
	// }))
	router.Use(middleware.Logger())

	// TODO: Implement custom logger (saves logs to database)
	// app.Use(logger)

	// serve up greeting page
	// router.StaticFS("/", "./public/pages/index.html")
	router.LoadHTMLGlob("public/pages/*.html")
	router.Static("/css", "public/css")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}
