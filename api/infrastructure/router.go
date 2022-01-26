package infrastructure

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"

	"github.com/watarun54/serverless-skill-manager/server/interfaces/controllers"
)

var echoLambda *echoadapter.EchoLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func Init() {
	e := echo.New()

	scrapeController := controllers.NewScrapeController(NewScrapeHandler())
	linebotController := controllers.NewLinebotController(NewSqlHandler(), NewScrapeHandler())
	authController := controllers.NewAuthController(NewSqlHandler())
	userController := controllers.NewUserController(NewSqlHandler())
	paperController := controllers.NewPaperController(NewSqlHandler(), NewScrapeHandler())
	roomController := controllers.NewRoomController(NewSqlHandler())
	commentController := controllers.NewCommentController(NewSqlHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Health Check
	e.GET("/", func(c echo.Context) error { return c.String(200, "OK") })

	e.POST("/login", func(c echo.Context) error { return authController.Login(c) })
	e.POST("/signup", func(c echo.Context) error { return userController.Create(c) })

	e.POST("/scrape/title", func(c echo.Context) error { return scrapeController.GetPaperTitle(c) })

	line := e.Group("/linebot")
	line.GET("", func(c echo.Context) error { return linebotController.GetTest(c) })
	line.POST("", func(c echo.Context) error { return linebotController.Post(c) })

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(controllers.NewJWTConfig()))
	api.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	api.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	api.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	api.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })

	api.GET("/papers", func(c echo.Context) error { return paperController.Index(c) })
	api.GET("/papers/:id", func(c echo.Context) error { return paperController.Show(c) })
	api.POST("/papers", func(c echo.Context) error { return paperController.Create(c) })
	api.PUT("/papers/:id", func(c echo.Context) error { return paperController.Update(c) })
	api.DELETE("/papers/:id", func(c echo.Context) error { return paperController.DeleteLogically(c) })
	api.DELETE("/papers/:id/complete", func(c echo.Context) error { return paperController.Delete(c) })

	api.GET("/comments", func(c echo.Context) error { return commentController.Index(c) })
	api.GET("/comments/:id", func(c echo.Context) error { return commentController.Show(c) })
	api.POST("/comments", func(c echo.Context) error { return commentController.Create(c) })
	api.PUT("/comments/:id", func(c echo.Context) error { return commentController.Update(c) })
	api.DELETE("/comments/:id", func(c echo.Context) error { return commentController.Delete(c) })

	api.GET("/rooms", func(c echo.Context) error { return roomController.Index(c) })
	api.GET("/rooms/:id", func(c echo.Context) error { return roomController.Show(c) })
	api.POST("/rooms", func(c echo.Context) error { return roomController.Create(c) })
	api.PUT("/rooms/:id", func(c echo.Context) error { return roomController.Update(c) })
	api.DELETE("/rooms/:id", func(c echo.Context) error { return roomController.Delete(c) })

	// Start server
	isLambda := os.Getenv("LAMBDA")
	if isLambda == "TRUE" {
		echoLambda = echoadapter.New(e)
		lambda.Start(Handler)
	} else {
		e.Logger.Fatal(e.Start(":8000"))
	}
}
