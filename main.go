package main

import (
	_ "embed"
	"net/http"

	"github/zjzjzjzj1874/my-gin/controllers"
	"github/zjzjzjzj1874/my-gin/docs"
	"github/zjzjzjzj1874/my-gin/middleware"

	"github.com/gin-gonic/gin"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @schemes http
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app := gin.Default()
	app.Use(middleware.Cors()) // 跨域设置

	app.GET("/swagger", func(c *gin.Context) {
		// 输出swagger文件
		c.Data(http.StatusOK, "text/plain; charset=utf-8", docs.SwaggerByte)
	})

	c := controllers.NewController()
	v1 := app.Group("/api/v1")
	{
		accounts := v1.Group("/account")
		{
			accounts.GET("", c.ListAccounts)
		}
	}
	app.GET("/echo", WSEchoHandler)

	if err := app.Run(":28888"); err != nil {
		panic(err)
	}
}
