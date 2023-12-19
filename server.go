package main

import (
	"gonic_api_2/config"
	"gonic_api_2/controller"
	"gonic_api_2/middleware"
	"gonic_api_2/repository"
	"gonic_api_2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoute := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoute.GET("/profile", userController.Profile)
		userRoute.PUT("/update", userController.Update)
	}

	r.Run()
}
