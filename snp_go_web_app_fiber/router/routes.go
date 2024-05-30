package router

import (
	"snp_go_web_app_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	setupUserRoutes(app)
}

func setupUserRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Home",
		}, "layouts/main")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"title": "About",
		}, "layouts/main")
	})

	app.Get("/sampleprogram/search", controllers.SearchAllSamplePrograms)
	app.Get("/sampleprogram/all", controllers.GetAllSamplePrograms)
	app.Get("/sampleprogram/count", controllers.GetSampleProgramCount)
	app.Post("/sampleprogram/create", controllers.CreateSampleProgram)
	app.Get("/sampleprogram/item/:id", controllers.GetSampleProgramById)
	app.Get("/sampleprogram/details/:id", controllers.GetSampleProgramDetailsById)
	app.Put("/sampleprogram/:id", controllers.EditSampleProgram)
	app.Delete("/sampleprogram/:id", controllers.DeleteSampleProgram)
	app.Get("/sampleprogram", controllers.GetSampleProgramPage)

	app.Get("/sample/search", controllers.SearchAllSamples)
	app.Get("/sample/all", controllers.GetAllSamples)
	app.Get("/sample/count", controllers.GetSampleCount)
	app.Post("/sample/create", controllers.CreateSample)
	app.Get("/sample/item/:id", controllers.GetSampleProgramById)
	app.Get("/sample/details/:id", controllers.GetSampleDetailsById)
	app.Put("/sample/:id", controllers.EditSample)
	app.Delete("/sample/:id", controllers.DeleteSample)
	app.Get("/sample", controllers.GetSamplePage)
}
