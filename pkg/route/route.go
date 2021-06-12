package route

import (
	"github.com/gin-gonic/gin"
	service "github.com/lion-devs/ilife-api/pkg/auth"
	"github.com/lion-devs/ilife-api/pkg/controller"
	"net/http"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	router := gin.Default()

	baseGroup := router.Group("/api/v1")
	{
		baseGroup.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello world",
				"interface": []interface{}{
					"str1",
					"str2",
				},
				"map": map[string]string{
					"k1": "v1",
					"k2": "v2",
				},
			})
		})

		baseGroup.POST("/login", func(ctx *gin.Context) {
			token := loginController.Login(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"data": map[string]string{
						"token": token,
					},
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		})
	}

	userGroup := router.Group("/api/v1/user")
	{
		userGroup.GET("/", controller.GetUsers)
		userGroup.POST("/", controller.CreateUser)
		userGroup.GET("/:id", controller.GetUserByID)
		userGroup.PUT("/:id", controller.UpdateUser)
		userGroup.DELETE("/:id", controller.DeleteUser)
	}
	return router
}
