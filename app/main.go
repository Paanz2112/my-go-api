package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"workspace/handler"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// type mobile struct {
// 	Manufacturer string `json: "manufacturer`
// 	Model        string `json: "model"`
// 	Form         string `json: "form"`
// 	Smartphone   string `json: "smartphone"`
// 	Year_        int16  `json: "year"`
// 	Units_sold_m int16  `json: "unit_sold"`
// 	Ids          int16  `json: "ids"`
// }
// type mobiles struct {
// 	Mobiles []mobile `json: "mobiles"`
// }

func main() {

	connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	log.Println(db.Ping())
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	// app.Use(middleware.Logger())

	app.Get("/", func(c *fiber.Ctx) error {
		// Get all data
		return handler.GetAllRows(c, db)
	})

	app.Get("/:manufac", func(c *fiber.Ctx) error {
		// Get by manufacturer
		return handler.GetByManufac(c, db)
	})

	app.Post("/post", func(c *fiber.Ctx) error {
		return handler.AddNewRow(c, db)
	})

	// app.Put("/update", func(c *fiber.Ctx) error {
	// 	return putHandler(c, db)
	// })

	app.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handler.DeleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

// func getAllRows(c *fiber.Ctx, db *sql.DB) error {
// 	results := mobiles{}
// 	rows, err := db.Query("SELECT * FROM public.mobile")
// 	if err != nil {
// 		log.Fatalln(err)
// 		c.JSON("An error occured")
// 	}

// 	defer rows.Close()
// 	if rows.Next() {
// 		for rows.Next() {
// 			m := mobile{}
// 			err := rows.Scan(&m.Manufacturer, &m.Model, &m.Form, &m.Smartphone, &m.Year_, &m.Units_sold_m, &m.Ids)
// 			if err != nil {
// 				return c.Status(500).JSON(&fiber.Map{
// 					"success": false,
// 					"error":   err,
// 				})
// 			}
// 			results.Mobiles = append(results.Mobiles, m)
// 		}
// 		return c.Status(200).JSON(results)
// 	} else {
// 		return c.Status(500).JSON(&fiber.Map{
// 			"success": false,
// 			"error":   "no rows return",
// 		})
// 	}

// }

// func getByManufac(c *fiber.Ctx, db *sql.DB) error {
// 	var manuName string = cases.Title(language.Und).String(c.Params("manufac"))
// 	log.Println(manuName)
// 	results := mobiles{}
// 	rows, err := db.Query("SELECT * FROM public.mobile WHERE manufacturer = $1", manuName)
// 	log.Println(rows)
// 	if err != nil {
// 		log.Fatalln(err)
// 		// c.JSON("An error occured")
// 	}
// 	defer rows.Close()
// 	if rows.Next() {
// 		for rows.Next() {
// 			m := mobile{}
// 			err := rows.Scan(&m.Manufacturer, &m.Model, &m.Form, &m.Smartphone, &m.Year_, &m.Units_sold_m, &m.Ids)
// 			if err != nil {
// 				return c.Status(500).JSON(&fiber.Map{
// 					"success": false,
// 					"error":   err,
// 				})
// 			}
// 			log.Println(m)
// 			results.Mobiles = append(results.Mobiles, m)
// 		}

// 		return c.Status(200).JSON(results)
// 	} else {
// 		return c.Status(500).JSON(&fiber.Map{
// 			"success": false,
// 			"error":   "no rows return",
// 		})
// 	}
// }

// func addNewRow(c *fiber.Ctx, db *sql.DB) error {
// 	m := new(mobile)
// 	// fmt.Println(m)
// 	if err := c.BodyParser(&m); err != nil {
// 		log.Fatalln(err)
// 		return c.JSON(fiber.Map{
// 			"error": err,
// 		})
// 	}
// 	if m.Manufacturer != "" && m.Model != "" && m.Units_sold_m != 0 && m.Year_ != 0 {
// 		res, err := db.Query("INSERT INTO public.mobile(manufacturer, model, form, smartphone, year_, units_sold_m) VALUES ($1, $2, $3, $4, $5, $6)", m.Manufacturer, m.Model, m.Form, m.Smartphone, m.Year_, m.Units_sold_m)
// 		log.Println(res)
// 		defer res.Close()
// 		// error handle
// 		if err != nil {
// 			log.Fatalln(err)
// 			return c.JSON(err)
// 		} else {
// 			return c.JSON(fiber.Map{
// 				"success": true,
// 				"mobile":  m,
// 				"message": "new data has been added.",
// 			})

// 		}
// 	} else {
// 		return c.JSON(fiber.Map{
// 			"error":   "please input data in correct format",
// 			"example": "{Manufacturer:Nokia,Model:1110,Form:Bar,Smartphone:No,Year_:2005,Units_sold_m:247,Ids:2}",
// 		})
// 	}

// }

// func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		log.Println(err)
// 		return c.Status(500).JSON(fiber.Map{
// 			"success": false,
// 			"message": "please input id in interger format.",
// 		})
// 	}
// 	log.Println(id)
// 	res, err := db.Query("DELETE FROM public.mobile WHERE ids = $1", id)
// 	log.Println(res)
// 	if err != nil {
// 		log.Println(err)
// 		return c.Status(500).JSON(fiber.Map{
// 			"success": false,
// 			"message": err,
// 		})
// 	} else {
// 		return c.Status(200).JSON(fiber.Map{
// 			"success": true,
// 			"message": "Data successfully deleted.",
// 		})
// 	}

// }
