package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
	Get(string) interface{}
	Request() *http.Request
	Response() *echo.Response
}
