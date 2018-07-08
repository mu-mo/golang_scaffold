package main

import (
	"config"
	"fmt"
	"net/http"
	"view"
	"view/admin"

	mid "middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	validator "gopkg.in/go-playground/validator.v9"
)

func healthy(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(mid.DefaultJWTConfig))
	e.Use(middleware.CSRFWithConfig(mid.DefaultCSRFConfig))
	//  against cross-site scripting (XSS) attack, content type sniffing, clickjacking, insecure connection and other code injection attacks
	e.Use(middleware.Secure())
	// 参数验证器
	e.Validator = &mid.DefaultValidator{Validator: validator.New()}

	e.GET("/healthy", healthy)

	view.InitIndexView(e)

	v1 := e.Group("/api/v1")
	user := v1.Group("/user")
	view.InitUserView(user)
	adminGroup := v1.Group("/admin")
	admin.InitAdminMemView(adminGroup)

	// 启动
	e.Logger.Fatal(e.Start(config.Conf.Addr))
	fmt.Println("启动")
}