package router

import (
	"project/vnexpress/api/controllers"
	repository "project/vnexpress/api/repositories"
	"project/vnexpress/config/driver/database"

	"github.com/gofiber/fiber/v2"
)

// thiết lập thông tin route
func SetupRouter(app *fiber.App) {
	db := database.DBConn
	articleRepo := repository.NewArticleRepository(db)
	articleController := controllers.NewArticleController(articleRepo)
	// api := app.Group("/api")
	app.Get("/api/vnexpress", articleController.ArticleList)
	app.Get("/api/vnexpress/:id", articleController.AritcleDetail)
	app.Post("/api/vnexpress/insert", articleController.ArticleCreate)
	app.Put("/api/vnexpress/:id", articleController.ArticleUpdate)
	app.Delete("/api/vnexpress/:id", articleController.ArticleDelete)
	app.Post("/register", controllers.Signup)
	app.Post("/login", controllers.Login)

}
