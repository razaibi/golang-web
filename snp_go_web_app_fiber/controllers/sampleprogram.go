package controllers

import (
	"snp_go_web_app_fiber/db"
	"snp_go_web_app_fiber/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetSampleProgramPage(ctx *fiber.Ctx) error {

	return ctx.Render("sample-program", fiber.Map{
		"title": "Sample Program",
	}, "layouts/main")
}

func GetSampleProgramDetailsById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Invalid ID")
	}

	var sampleProgram models.SampleProgram
	err = db.GetDB().Get(&sampleProgram, "SELECT * FROM SampleProgram WHERE SampleProgramId = ?", id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Sample program not found")
	}
	return ctx.JSON(sampleProgram)
}

func CreateSampleProgram(ctx *fiber.Ctx) error {
	var sampleProgram models.SampleProgram
	if err := ctx.BodyParser(&sampleProgram); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	sampleProgram.CreatedOn = time.Now()
	sampleProgram.LastModifiedOn = time.Now()
	result, err := db.GetDB().NamedExec(`INSERT INTO SampleProgram (Name, IsActive, CreatedOn, LastModifiedOn) VALUES (:Name, :IsActive, :CreatedOn, :LastModifiedOn)`, &sampleProgram)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Retrieve the last inserted ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Assign the last inserted ID to the sampleProgram object
	sampleProgram.SampleProgramId = int(lastInsertID)
	return ctx.JSON(sampleProgram)
}

func EditSampleProgram(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Invalid ID")
	}

	var sampleProgram models.SampleProgram
	if err := ctx.BodyParser(&sampleProgram); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if id != sampleProgram.SampleProgramId {
		return ctx.Status(fiber.StatusNotFound).SendString("ID mismatch")
	}

	sampleProgram.LastModifiedOn = time.Now()
	_, err = db.GetDB().NamedExec(`UPDATE SampleProgram SET Name = :Name, IsActive = :IsActive, LastModifiedOn = :LastModifiedOn WHERE SampleProgramId = :SampleProgramId`, &sampleProgram)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(sampleProgram)
}

func DeleteSampleProgram(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Invalid ID")
	}

	_, err = db.GetDB().Exec("DELETE FROM SampleProgram WHERE SampleProgramId = ?", id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusAccepted)
}

func GetAllSamplePrograms(ctx *fiber.Ctx) error {
	skip := ctx.QueryInt("skip", 0)
	take := ctx.QueryInt("take", 7)

	var samplePrograms []models.SampleProgram
	err := db.GetDB().Select(&samplePrograms, "SELECT * FROM SampleProgram ORDER BY LastModifiedOn DESC LIMIT ? OFFSET ?", take, skip)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Sample programs not found")
	}
	return ctx.JSON(fiber.Map{
		"result": samplePrograms,
	})
}

func SearchAllSamplePrograms(ctx *fiber.Ctx) error {
	term := ctx.Query("term")
	if term == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Query parameter 'term' is required"})
	}
	var samplePrograms []models.SampleProgram
	// Correcting the SQL query and parameter usage
	query := "SELECT * FROM SampleProgram WHERE name LIKE ? LIMIT 7"
	err := db.GetDB().Select(&samplePrograms, query, "%"+term+"%")
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Sample programs not found")
	}
	return ctx.JSON(fiber.Map{
		"result": samplePrograms,
	})
}

func GetSampleProgramCount(ctx *fiber.Ctx) error {
	var count int
	err := db.GetDB().Get(&count, "SELECT COUNT(*) FROM SampleProgram")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(count)
}

func GetSampleProgramById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Invalid ID")
	}

	var sampleProgram models.SampleProgram
	err = db.GetDB().Get(&sampleProgram, "SELECT * FROM SampleProgram WHERE SampleProgramId = ?", id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Sample program not found")
	}
	return ctx.JSON(sampleProgram)
}
