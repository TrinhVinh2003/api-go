package controllers

import (
	"log"
	"project/vnexpress/api/models"
	repository "project/vnexpress/api/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	Repo repository.ArticleRepository
}

func NewArticleController(repo repository.ArticleRepository) *ArticleController {
	return &ArticleController{Repo: repo}
}

func (a *ArticleController) ArticleList(c *fiber.Ctx) error {
	articles, err := a.Repo.FindAll()
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"msg":        "Faith",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "Ok",
		"msg":        "Article List",
		"records":    articles,
	})

}
func (a *ArticleController) AritcleDetail(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	article, err := a.Repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"msg":        "Not found id",
		})

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"StatusText": "Ok",
		"msg":        "Article detail",
		"records":    article,
	})

}

func (a *ArticleController) ArticleCreate(c *fiber.Ctx) error {
	var article models.Article
	if err := c.BodyParser(&article); err != nil {
		log.Println("Error in parsing request:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"msg":        "Invalid request data.",
		})
	}
	if err := a.Repo.Create(&article); err != nil {
		log.Println("Error in saving data:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"msg":        "Failed to save the record.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "Ok",
		"msg":        "Record saved successfully.",
		"data":       article,
	})
}
func (ac *ArticleController) ArticleUpdate(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"msg":        "Invalid ID format.",
		})
	}

	var article models.Article
	if err := c.BodyParser(&article); err != nil {
		log.Println("Error in parsing request:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"msg":        "Invalid request data.",
		})
	}
	article.ID = uint(id)

	if err := ac.Repo.Update(&article); err != nil {
		log.Println("Error in updating data:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"msg":        "Failed to update the record.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "Ok",
		"msg":        "Record updated successfully.",
	})
}

func (ac *ArticleController) ArticleDelete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusText": "Error",
			"msg":        "Invalid ID format.",
		})
	}

	if err := ac.Repo.Delete(uint(id)); err != nil {
		log.Println("Error in deleting data:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusText": "Error",
			"msg":        "Failed to delete the record.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusText": "Ok",
		"msg":        "Record deleted successfully.",
	})
}

// package controllers

// import (
// 	"log"
// 	model "project/vnexpress/api/models"
// 	"project/vnexpress/internal/database"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// )

// // list articles
// func ArticleList(c *fiber.Ctx) error {

// 	context := fiber.Map{
// 		"statusText": "Ok",
// 		"msg":        "Article List",
// 	}

// 	time.Sleep(time.Millisecond * 1500)

// 	db := database.DBConn

// 	var records []model.Article

// 	db.Find(&records)

// 	context["article_records"] = records

// 	c.Status(200)
// 	return c.JSON(context)
// }

// // xem chi tiết 1 bài báo
// func ArticleDetail(c *fiber.Ctx) error {
// 	c.Status(400)
// 	context := fiber.Map{
// 		"statusText": "",
// 		"msg":        "",
// 	}

// 	id := c.Params("id")

// 	var record model.Article

// 	database.DBConn.First(&record, id)

// 	if record.ID == 0 {
// 		log.Println("Record not Found.")
// 		context["msg"] = "Record not Found."

// 		c.Status(404)
// 		return c.JSON(context)
// 	}

// 	context["record"] = record
// 	context["statusText"] = "Ok"
// 	context["msg"] = "Article Detail"
// 	c.Status(200)
// 	return c.JSON(context)
// }

// // tạo 1 bài báo
// func ArticleCreate(c *fiber.Ctx) error {
// 	context := fiber.Map{
// 		"statusText": "Ok",
// 		"msg":        "Add a Article",
// 	}

// 	record := new(model.Article)

// 	if err := c.BodyParser(record); err != nil {
// 		log.Println("Error in parsing request.")
// 		context["statusText"] = ""
// 		context["msg"] = "Something went wrong."
// 	}

// 	result := database.DBConn.Create(record)

// 	if result.Error != nil {
// 		log.Println("Error in saving data.")
// 		context["statusText"] = ""
// 		context["msg"] = "Something went wrong."
// 	}

// 	context["msg"] = "Record is saved successully."
// 	context["data"] = record

// 	c.Status(201)
// 	return c.JSON(context)
// }

// // cập nhật 1 bài báo
// func ArticleUpdate(c *fiber.Ctx) error {

// 	context := fiber.Map{
// 		"statusText": "Ok",
// 		"msg":        "Update Article",
// 	}
// 	id := c.Params("id")

// 	var record model.Article

// 	database.DBConn.First(&record, id)

// 	if record.ID == 0 {
// 		log.Println("Record not Found.")

// 		context["statusText"] = ""
// 		context["msg"] = "Record not Found."
// 		c.Status(400)
// 		return c.JSON(context)
// 	}

// 	if err := c.BodyParser(&record); err != nil {
// 		log.Println("Error in parsing request.")

// 		context["msg"] = "Something went wrong."
// 		c.Status(400)
// 		return c.JSON(context)

// 	result := database.DBConn.Save(record)

// 	if result.Error != nil {
// 		log.Println("Error in saving data.")

// 		context["msg"] = "Error in saving data."
// 		c.Status(400)
// 		return c.JSON(context)
// 	}

// 	context["msg"] = "Record updated successfully."
// 	context["data"] = record

// 	c.Status(200)
// 	return c.JSON(context)
// }
// func ArticleDelete(c *fiber.Ctx) error {

// 	c.Status(400)
// 	context := fiber.Map{
// 		"statusText": "",
// 		"msg":        "",
// 	}

// 	id := c.Params("id")

// 	var record model.Article

// 	database.DBConn.First(&record, id)

// 	if record.ID == 0 {
// 		log.Println("Record not Found.")
// 		context["msg"] = "Record not Found."

// 		return c.JSON(context)
// 	}

// 	result := database.DBConn.Delete(record)

// 	if result.Error != nil {
// 		context["msg"] = "Something went wrong."
// 		return c.JSON(context)
// 	}

// 	context["statusText"] = "Ok"
// 	context["msg"] = "Record deleted successfully."
// 	c.Status(200)
// 	return c.JSON(context)
// }
