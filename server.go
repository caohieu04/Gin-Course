package main

import (
	"io"
	"os"

	"github.com/caohieu04/Gin-Course/api"
	"github.com/caohieu04/Gin-Course/controller"
	"github.com/caohieu04/Gin-Course/docs"
	"github.com/caohieu04/Gin-Course/repository"
	"github.com/caohieu04/Gin-Course/service"
	"github.com/gin-gonic/gin"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	// gin.SetMode(gin.ReleaseMode)

	setupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	// server.Use(gin.Recovery(), middlewares.Logger(),
	// 	middlewares.BasicAuth(), middlewares.AuthorizeJWT(), gindump.Dump())

	server.Use(gin.Recovery(), gin.Logger())

	videoAPI := api.NewVideoApi(loginController, videoController)
	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}
	}

	// server.POST("/login", func(ctx *gin.Context) {
	// 	token := loginController.Login(ctx)
	// 	if token != "" {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"token": token,
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusUnauthorized, nil)
	// 	}
	// })

	// apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	// {

	// 	apiRoutes.GET("/videos", func(ctx *gin.Context) {
	// 		ctx.JSON(200, videoController.FindAll())
	// 	})
	// 	apiRoutes.POST("/videos", func(ctx *gin.Context) {
	// 		err := videoController.Save(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
	// 		}
	// 	})
	// 	apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
	// 		err := videoController.Update(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "PUT Success!"})
	// 		}
	// 	})
	// 	apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
	// 		err := videoController.Delete(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "DELETE Success!"})
	// 		}
	// 	})
	// }

	// viewRoutes := server.Group("/view")
	// {
	// 	viewRoutes.GET("/videos", videoController.ShowAll)
	// }
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	// server.Run("127.0.0.1:" + port)
	server.Run(":" + port)
	// server.Run(":8080")
}
