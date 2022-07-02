package routes

import (
	"futuremap/controllers"
	"futuremap/middlewares"
	"futuremap/models"
	"time"
  	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	models.ConnectDatabase()
	r := gin.Default()
	//user cors
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	public := r.Group("/")
	public.GET("/", controllers.HomeList)
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.POST("/admin/register", controllers.RegisterAdmin)
	public.POST("/reset/password", controllers.ResetPassword)
	public.GET("/materials", controllers.LearningList)
	public.GET("/materials/:id", controllers.GetLearningID)
	public.GET("/materials/:id/discussion", controllers.ShowDiscussion)
	//=============================Middlewares for Client======================================================
	client := r.Group("/client")
	client.Use(middlewares.JwtAuthMiddleware())
	client.GET("/user", controllers.CurrentUser)
	client.PUT("/update/profile", controllers.UpdateProfile)
	client.GET("/materials/:id", controllers.GetLearningById)
	client.POST("/materials/:id", controllers.GetLearningById)
	client.POST("/materials/:id/discussion", controllers.MakeDiscussion)
	client.GET("/history", controllers.GetHistory)
	//=============================Middlewares for Admin======================================================
	admin := r.Group("/admin")
	admin.Use(middlewares.JwtAuthMiddlewareAdmin())
	admin.PUT("/materials/:id", controllers.UpdateLearning)
	admin.DELETE("/materials/:id", controllers.DeleteLearning)
	admin.GET("/user", controllers.CurrentUser)
	admin.PUT("/update/profile", controllers.UpdateProfile)
	admin.POST("/materials", controllers.Learning)
	admin.GET("/materials/:id", controllers.GetLearningById)
	admin.POST("/materials/:id", controllers.GetLearningById)
	admin.POST("/materials/:id/discussion", controllers.MakeDiscussion)
	admin.POST("/home", controllers.Home)
	admin.PUT("/home", controllers.UpdateHome)
	admin.DELETE("/home", controllers.DeleteHome)
	admin.GET("/history", controllers.GetHistory)
	return r
}
