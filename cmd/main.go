package main

import (
	"github.com/lion-devs/ilife-api/pkg/auth"
	"github.com/lion-devs/ilife-api/pkg/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
			"interface": []interface{}{
				"str1",
				"str2",
			},
			"map": map[string]string {
				"k1": "v1",
				"k2": "v2",
			},
		})
	})

	server.POST("/api/v1/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"data": map[string]string {
					"token": token,
				},
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	port := "8080"
	server.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
