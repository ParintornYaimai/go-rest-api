package handler

import (
	"database/sql"

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

func getById(c *fiber.Ctx) {

}

func update(c *fiber.Ctx) {

}

func delete(c *fiber.Ctx) {

}
