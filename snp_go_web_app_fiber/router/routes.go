package router

import (
	"snp_go_web_app_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	setupUserRoutes(app)
}

func setupUserRoutes(app *fiber.App) {
	// Group for general routes
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

	// Group for SampleProgram routes
	sampleProgramGroup := app.Group("/sampleprogram")
	sampleProgramGroup.Get("/search", controllers.SearchAllSamplePrograms)
	sampleProgramGroup.Get("/all", controllers.GetAllSamplePrograms)
	sampleProgramGroup.Get("/count", controllers.GetSampleProgramCount)
	sampleProgramGroup.Post("/create", controllers.CreateSampleProgram)
	sampleProgramGroup.Get("/details/:id", controllers.GetSampleProgramDetailsById)
	sampleProgramGroup.Get("/item/:id", controllers.GetSampleProgramById)
	sampleProgramGroup.Put("/:id", controllers.EditSampleProgram)
	sampleProgramGroup.Delete("/:id", controllers.DeleteSampleProgram)
	sampleProgramGroup.Get("", controllers.GetSampleProgramPage)

	// Group for Sample routes
	sampleGroup := app.Group("/sample")
	sampleGroup.Get("/search", controllers.SearchAllSamples)
	sampleGroup.Get("/all", controllers.GetAllSamples)
	sampleGroup.Get("/count", controllers.GetSampleCount)
	sampleGroup.Post("/create", controllers.CreateSample)
	sampleGroup.Get("/details/:id", controllers.GetSampleDetailsById)
	sampleGroup.Get("/item/:id", controllers.GetSampleById)
	sampleGroup.Put("/:id", controllers.EditSample)
	sampleGroup.Delete("/:id", controllers.DeleteSample)
	sampleGroup.Get("", controllers.GetSamplePage)
}
