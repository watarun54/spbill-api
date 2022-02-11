module github.com/watarun54/serverless-skill-manager/server

go 1.13

require (
	github.com/aws/aws-lambda-go v0.0.0-20190129190457-dcf76fe64fb6
	github.com/awslabs/aws-lambda-go-api-proxy v0.6.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/line/line-bot-sdk-go v7.3.0+incompatible
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca
	github.com/valyala/fasttemplate v1.1.0 // indirect
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	google.golang.org/appengine v1.4.0
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0
