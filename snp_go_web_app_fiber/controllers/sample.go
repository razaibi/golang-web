package controllers

import (
	"fmt"
	"snp_go_web_app_fiber/db"
	"snp_go_web_app_fiber/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetSamplePage(ctx *fiber.Ctx) error {

	return ctx.Render("sample", fiber.Map{
		"title": "Sample",
	}, "layouts/main")
}

func GetSampleDetailsById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Invalid ID")
	}

	var sample models.Sample
	query := `SELECT s.*, sp.Name as SampleProgramName
	FROM Sample s
	JOIN SampleProgram sp ON s.SampleProgramId = sp.SampleProgramId
	WHERE s.SampleId = ?`
	err = db.GetDB().Get(&sample, query, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Sample not found")
	}
	return c.JSON(sample)
}

func CreateSample(c *fiber.Ctx) error {
	sample := new(models.Sample)
	if err := c.BodyParser(sample); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	sample.CreatedOn = time.Now()
	sample.LastModifiedOn = time.Now()
	_, err := db.GetDB().NamedExec(`INSERT INTO Sample (SampleProgramId, Title, Content, IsActive, CreatedOn, LastModifiedOn) 
		VALUES (:SampleProgramId, :Title, :Content, :IsActive, :CreatedOn, :LastModifiedOn)`, sample)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(sample)
}

func EditSample(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Invalid ID")
	}

	sample := new(models.Sample)
	if err := c.BodyParser(sample); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if sample.SampleId != int(id) {
		return c.Status(fiber.StatusNotFound).SendString("Sample ID mismatch")
	}

	sample.LastModifiedOn = time.Now()
	_, err = db.GetDB().NamedExec(`UPDATE Sample SET SampleProgramId = :SampleProgramId, Title = :Title, Content = :Content, IsActive = :IsActive, LastModifiedOn = :LastModifiedOn WHERE SampleId = :SampleId`, sample)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var result models.Sample
	err = db.GetDB().Get(&result, `SELECT * FROM Sample WHERE SampleId = ?`, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Sample not found")
	}

	return c.JSON(result)
}

func DeleteSample(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Invalid ID")
	}

	_, err = db.GetDB().Exec(`DELETE FROM Sample WHERE SampleId = ?`, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

func GetAllSamples(ctx *fiber.Ctx) error {
	var samples []models.SampleOutDto
	skip, err := strconv.Atoi(ctx.Query("skip", "0"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid skip parameter")
	}
	take := "7"

	query := `SELECT s.*, sp.Name as SampleProgramName
              FROM Sample s
              JOIN SampleProgram sp ON s.SampleProgramId = sp.SampleProgramId
              ORDER BY s.LastModifiedOn DESC
              LIMIT ? OFFSET ?`
	err = db.GetDB().Select(&samples, query, take, skip)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusNotFound).SendString("No samples found")
	}
	return ctx.JSON(fiber.Map{
		"result": samples,
	})
}

func SearchAllSamples(c *fiber.Ctx) error {
	term := c.Query("term")

	var samples []models.Sample
	if term != "" {
		query := `SELECT s.*, sp.name as SampleProgramName
		FROM Sample s
		JOIN SampleProgram sp ON s.SampleProgramId = sp.SampleProgramId
		WHERE s.Title LIKE ?
		LIMIT 7`
		err := db.GetDB().Select(&samples, query, "%"+term+"%")
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("No samples found")
		}
	} else {
		err := db.GetDB().Select(&samples, `SELECT * FROM Sample LIMIT 7`)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("No samples found")
		}
	}
	return c.JSON(samples)
}

func GetSampleCount(ctx *fiber.Ctx) error {
	var count int
	err := db.GetDB().Get(&count, `SELECT COUNT(*) FROM sample`)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.JSON(count)
}

func GetSampleById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Invalid ID")
	}

	var sample models.Sample
	query := `SELECT s.*, sp.name as SampleProgramName
	FROM Sample s
	JOIN SampleProgram sp ON s.SampleProgramId = sp.SampleProgramId
	WHERE s.SampleId = ?`
	err = db.GetDB().Get(&sample, query, id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Sample not found")
	}
	return ctx.JSON(sample)
}
