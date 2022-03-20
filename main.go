package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.com/m/controller"
	"golang.com/m/docs"
	"golang.com/m/repository"
	"golang.com/m/service"
)

var (
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
	userRepository repository.UserRepository = repository.NewUserRepository()
	jwtService service.JWTService = service.NewJWTService()

	loginController controller.LoginController = controller.NewLoginController(jwtService)
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)

		if token != "" {
			ctx.JSON(200, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(401, gin.H{
				"message": "unauthorized",
			})
		}
	})

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/user")

		accounts.GET("/:id", func(c *gin.Context) {
			c.JSON(200, userController.GetAllUsers(c))
		})
		accounts.POST("", func(ctx *gin.Context) {
			ctx.JSON(200, userController.CreateUser(ctx))
		})

		//...
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
