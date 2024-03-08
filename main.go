package main

import (
	"fmt"
	"log"
	"time"

	"diserve-expl/auth"
	"diserve-expl/cache"
	"diserve-expl/cmd"
	"diserve-expl/controller"
	"diserve-expl/repository"
	"diserve-expl/service"

	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Load() string {
	jsonBytes, err := os.ReadFile("user.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return ""
	}
	return string(jsonBytes)
}

func init() {

	docTemplate := Load()
	var SwaggerInfo = &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  docTemplate,
	}
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

func DocsRoute(e *echo.Echo) {
	e.GET("/docs/*", echoSwagger.WrapHandler)
}

func ConnectDB() *gorm.DB {
	godotenv.Load()
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&cache.Post{})
	return db
}

func task() {
	log.Println("halow")
}

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	//shedulling in go
	// gocron.Every(1).Second().Do(task)
	// <-gocron.Start()
	// rootCmd.AddCommand(reverseCmd)

	cmd.Execute()

	DocsRoute(e)
	db := ConnectDB()
	rd := cache.NewRedisChace("redis-11884.c1.asia-northeast1-1.gce.cloud.redislabs.com: 11884", 12167185, time.Hour*1)
	rp := repository.NewRepo(db)
	svc := service.NewSvc(rp)
	h := controller.NewController(svc, rd)

	e.POST("/login", h.Login)
	e.POST("/datas", h.CreatePost)
	e.POST("/refresh", h.RefreshToken)

	auth := e.Group("/auth", auth.Middleware())
	auth.GET("/datas/:id", h.FindByID)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${p ath} ${latency_human}` + "\n",
	}))

	e.Logger.Fatal(e.Start(":8080"))
}
