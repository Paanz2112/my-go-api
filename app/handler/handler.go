package handler

import (
	"database/sql"
	"log"
	"strconv"
	"workspace/model"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetAllRows(c *fiber.Ctx, db *sql.DB) error {
	// Get All Mobile data from postgres db

	// declare variable
	results := model.Mobiles{}
	var countRow int

	// Get row count
	err := db.QueryRow("SELECT count(*) FROM public.mobile").Scan(&countRow)
	if err != nil {
		log.Fatalln(err)
	}

	// If some data return then get all data by condition
	if countRow > 0 {
		// Get all data
		rows, err := db.Query("SELECT * FROM public.mobile")
		if err != nil {
			log.Println("-----error from GetAllRows-----")
			log.Fatalln(err)
			return c.Status(500).JSON(&fiber.Map{"message": "An error occured"})
		}

		// Close db connection when done
		defer rows.Close()

		for rows.Next() {
			m := model.Mobile{}
			err := rows.Scan(&m.Manufacturer, &m.Model, &m.Form, &m.Smartphone, &m.Year_, &m.Units_sold_m, &m.Ids)
			if err != nil {
				return c.Status(500).JSON(&fiber.Map{
					"success": false,
					"error":   err,
				})
			}
			results.Mobiles = append(results.Mobiles, m)
		}
		return c.Status(200).JSON(results)
	} else {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   "no rows return",
		})
	}

}

func GetByManufac(c *fiber.Ctx, db *sql.DB) error {
	// Get data by Manufacturer

	// Get user params
	var manuName string = cases.Title(language.Und).String(c.Params("manufac"))
	log.Println("GET param from user ->> ", manuName)

	// declare variable
	results := model.Mobiles{}
	var countRow int

	// Get row count
	err := db.QueryRow("SELECT count(*) FROM public.mobile").Scan(&countRow)
	if err != nil {
		log.Fatalln(err)
	}

	// If some data return then get all data by condition
	if countRow > 0 {
		// Get all data
		rows, err := db.Query("SELECT * FROM public.mobile WHERE manufacturer = $1", manuName)
		if err != nil {
			log.Println("-----error from GetByManufac-----")
			log.Fatalln(err)
			return c.Status(500).JSON(&fiber.Map{"message": "An error occured"})
		}

		defer rows.Close()

		for rows.Next() {
			m := model.Mobile{}
			err := rows.Scan(&m.Manufacturer, &m.Model, &m.Form, &m.Smartphone, &m.Year_, &m.Units_sold_m, &m.Ids)
			if err != nil {
				return c.Status(500).JSON(&fiber.Map{
					"success": false,
					"error":   err,
				})
			}
			results.Mobiles = append(results.Mobiles, m)
		}
		return c.Status(200).JSON(results)
	} else {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   "no rows return",
		})
	}
}

func AddNewRow(c *fiber.Ctx, db *sql.DB) error {
	m := new(model.Mobile)
	// fmt.Println(m)
	if err := c.BodyParser(&m); err != nil {
		log.Println("-----error from AddNewRow-----")
		log.Fatalln(err)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}
	if m.Manufacturer != "" && m.Model != "" && m.Units_sold_m != 0 && m.Year_ != 0 {
		res, err := db.Query("INSERT INTO public.mobile(manufacturer, model, form, smartphone, year_, units_sold_m) VALUES ($1, $2, $3, $4, $5, $6)", m.Manufacturer, m.Model, m.Form, m.Smartphone, m.Year_, m.Units_sold_m)
		log.Println(res)
		defer res.Close()
		// error handle
		if err != nil {
			log.Println("-----error from AddNewRow-----")
			log.Fatalln(err)
			return c.Status(500).JSON(&fiber.Map{"message": "An error occured"})
		} else {
			return c.JSON(fiber.Map{
				"success": true,
				"mobile":  m,
				"message": "new data has been added.",
			})

		}
	} else {
		return c.JSON(fiber.Map{
			"error":   "please input data in correct format",
			"example": "{Manufacturer:Nokia,Model:1110,Form:Bar,Smartphone:No,Year_:2005,Units_sold_m:247,Ids:2}",
		})
	}

}

func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "please input id in interger format.",
		})
	}
	res, err := db.Query("DELETE FROM public.mobile WHERE ids = $1", id)
	log.Println(res)
	if err != nil {
		log.Println("-----error from AddNewRow-----")
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{"message": "An error occured"})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Data successfully deleted.",
		})
	}

}
