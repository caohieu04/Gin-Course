package main

import (
	"io"
	"os"

	"github.com/caohieu04/Gin-Course/api"
	"github.com/caohieu04/Gin-Course/controller"
	"github.com/caohieu04/Gin-Course/docs"
	"github.com/caohieu04/Gin-Course/middlewares"
	"github.com/caohieu04/Gin-Course/repository"
	"github.com/caohieu04/Gin-Course/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Pragmatic Reviews - Video API"
	docs.SwaggerInfo.Description = "Pragmatic Reviews - Youtube Video API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	server := gin.Default()
	server.Use(cors.Default())

	videoAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// We can setup this env variable from the EB console
	port := os.Getenv("PORT")

	// Elastic Beanstalk forwards requests to port 5000
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
