package controllers

import (
	"DB-DB/database"
	"DB-DB/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"io"
	"net/http"
	"os"
	"strconv"
)

func GetDashboard(c *fiber.Ctx) error {
	return c.SendFile("./data/index.html")
}

func GetTickets(c *fiber.Ctx) error {
	var tickets []models.DeliveryTicket

	database.DB.Find(&tickets)

	return c.JSON(tickets)
}

func GetPicture(c *fiber.Ctx) error {
	// return picture
	return c.SendFile(
		"./data/images/" + c.Params("fp"),
	)
}

func AddTicket(c *fiber.Ctx) error {
	// Add a new ticket entry to the database
	// format of ticket not finalised, format of request is draft final

	c.Accepts("application/json")

	// create ticket and parse json
	data := new(models.JSONDeliveryTicket)
	if err := c.BodyParser(data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ticket := new(models.DeliveryTicket)

	for _, field := range data.ExtractedFields {
		if field == "total weight" {
			fieldData := data.TotalWeight
			if fieldData.Success {
				float, err := strconv.ParseFloat(fieldData.Value, 32)
				if err != nil {
					return c.JSON(fiber.Map{
						"message": err.Error(),
					})
				}
				ticket.Weight = float32(float)
			}
		}
	}

	// get id
	id := uuid.New().String()

	// add id to ticket
	ticket.Id = id
	ticket.ImageFilepath = id + ".jpg"

	// store ticket to database
	result := database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(ticket)
	// nothing affected => conflict => show error
	if result.RowsAffected == 0 {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "could not add ticket",
		})
	}

	err := downloadImage(data.ImageUrl, id+".jpg")
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "partial success, ticket was given id: " + id + ", image download failed: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success, ticket was given id: " + id + ", image download successful",
	})

}

func downloadImage(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("response code is %d instead of 200", response.StatusCode))
	}
	file, err := os.Create(fmt.Sprintf("./data/images/%s", filename))
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
