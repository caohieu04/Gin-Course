package main

import (
	"io"
	"net/http"
	"os"

	"github.com/caohieu04/Gin-Course/controller"
	"github.com/caohieu04/Gin-Course/middlewares"
	"github.com/caohieu04/Gin-Course/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

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

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.AuthorizeJWT(), gindump.Dump())

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api")
	{

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	// server.Run("127.0.0.1:" + port)
	server.Run(":" + port)
	// server.Run(":8080")
}
