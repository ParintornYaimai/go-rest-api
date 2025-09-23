package handler

import (
	"database/sql"
	"strconv"

	"github.com/ParintornYaimai/go-rest-api/api/model"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func GetAllBooks(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		row, err := db.Query("SELECT * FROM Book")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// หลังทำทุกอย่างเสร็จปิดการเชื่อมต่อ
		defer row.Close()

		//สร้าง slice ไว้รอ
		books := []model.BookModel{}

		//วนลูปอ่านเเถวต่อไปถ้ามีต่อไปจะเป็น true ไม่มีจะเป็น false
		for row.Next() {
			//สร้าง struct b
			var b model.BookModel
			row.Scan(&b.ID, &b.Name, &b.Category)

			//บันทึกลง slice
			books = append(books, b)
		}

		return c.JSON(books)
	}
}

func GetById(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("bookId")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid bookId")
		}

		row := db.QueryRow("SELECT * FROM Book WHERE if = $1", id)

		var b model.BookModel
		if err := row.Scan(&b.ID, &b.Name, &b.Category); err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).SendString("Book not found")
			}

			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(b)
	}
}

func AddBook(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

	}
}

type BookInput struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func Update(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("bookId")

		var input BookInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		var updated model.BookModel
		err := db.QueryRow("UPDATE book SET Name=$1, Category=$2 WHERE id=$3", input.Name, input.Category, idParam).Scan(&updated.ID, &updated.Name, &updated.Category)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.JSON(updated)
	}
}

func Delete(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

	}
}
