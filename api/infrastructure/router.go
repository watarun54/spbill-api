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

	linebotController := controllers.NewLinebotController(NewSqlHandler())
	authController := controllers.NewAuthController(NewSqlHandler())
	userController := controllers.NewUserController(NewSqlHandler())
	roomController := controllers.NewRoomController(NewSqlHandler())
	billController := controllers.NewBillController(NewSqlHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Health Check
	e.GET("/", func(c echo.Context) error { return c.String(200, "OK") })

	e.POST("/login", func(c echo.Context) error { return authController.Login(c) })
	e.POST("/signup", func(c echo.Context) error { return userController.Create(c) })

	line := e.Group("/linebot")
	line.POST("", func(c echo.Context) error { return linebotController.GetTest(c) })
	line.POST("", func(c echo.Context) error { return linebotController.Post(c) })

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(controllers.NewJWTConfig()))
	api.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	api.GET("/users/me", func(c echo.Context) error { return userController.GetMe(c) })
	api.PUT("/users/me", func(c echo.Context) error { return userController.UpdateMe(c) })

	api.POST("/bills", func(c echo.Context) error { return billController.Create(c) })
	api.GET("/bills/:id", func(c echo.Context) error { return billController.Show(c) })
	api.PUT("/bills/:id", func(c echo.Context) error { return billController.Update(c) })
	api.DELETE("/bills/:id", func(c echo.Context) error { return billController.Delete(c) })

	api.GET("/rooms", func(c echo.Context) error { return roomController.Index(c) })
	api.POST("/rooms", func(c echo.Context) error { return roomController.Create(c) })
	api.GET("/rooms/:id", func(c echo.Context) error { return roomController.Show(c) })
	api.PUT("/rooms/:id", func(c echo.Context) error { return roomController.Update(c) })
	api.DELETE("/rooms/:id", func(c echo.Context) error { return roomController.Delete(c) })
	api.GET("/rooms/:id/bills", func(c echo.Context) error { return roomController.FetchBills(c) })
	api.GET("/rooms/:id/user_payments", func(c echo.Context) error { return roomController.UserPayments(c) })

	// Start server
	isLambda := os.Getenv("LAMBDA")
	if isLambda == "TRUE" {
		echoLambda = echoadapter.New(e)
		lambda.Start(Handler)
	} else {
		e.Logger.Fatal(e.Start(":8000"))
	}
}
