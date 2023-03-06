package controllers

import (
	"DB-DB/database"
	"DB-DB/methods"
	"DB-DB/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"io"
	"log"
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

	files, err := os.ReadDir("./data/images")
	if err != nil {
		log.Print(err)
	}
	if len(tickets) != len(files) { // too many files, need to do some cleanup
		// tickets and files are both already sorted by uuid
		ticketIndex := 0
		for _, file := range files {
			if ticketIndex < len(tickets) && tickets[ticketIndex].ImageFilepath == file.Name() {
				ticketIndex++
			} else { //file not referenced, can delete
				_ = deleteImage(file.Name())
			}
		}
	}

	return c.JSON(tickets)
}

func GetPicture(c *fiber.Ctx) error {
	// return picture
	return c.SendFile(
		"./data/images/" + c.Params("fp"),
	)
}

func GetFailedChecks(c *fiber.Ctx) error {
	var checks []string

	database.DB.Model(&models.FailedChecks{}).
		Where("ticket_id = ? AND quantity = ?", c.Params("id"), c.Params("quantity")).
		Select("check_name").Find(&checks)

	return c.JSON(checks)
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

	// get id
	id := uuid.New().String()

	ticket := new(models.DeliveryTicket)

	for _, field := range data.ExtractedFields {
		if field == "total weight" {
			fieldData := data.TotalWeight
			ticket.WeightSuccess = fieldData.Success
			if fieldData.Success {
				float, err := strconv.ParseFloat(fieldData.Value, 64)
				if err != nil {
					return c.JSON(fiber.Map{
						"message": err.Error(),
					})
				}
				ticket.Weight = methods.RoundWithPrecision(float, 3)

				if !fieldData.Checks.DecimalPlaceCheck {
					check := new(models.FailedChecks)
					check.Id = id + "-DPC"
					check.TicketId = id
					check.Quantity = "weight"
					check.CheckName = "Decimal Place Check"
					result := database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(check)
					// nothing affected => conflict => show error
					if result.RowsAffected == 0 {
						c.Status(fiber.StatusForbidden)
						return c.JSON(fiber.Map{
							"message": "could not add check",
						})
					}
				}

				if !fieldData.Checks.ExtremeValueCheck {
					check := new(models.FailedChecks)
					check.Id = id + "-EVC"
					check.TicketId = id
					check.Quantity = "weight"
					check.CheckName = "Extreme Value Check"
					result := database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(check)
					// nothing affected => conflict => show error
					if result.RowsAffected == 0 {
						c.Status(fiber.StatusForbidden)
						return c.JSON(fiber.Map{
							"message": "could not add check",
						})
					}
				}
			} else {
				ticket.WeightErrorType = fieldData.ErrorInformation.Type
				ticket.WeightErrorDescription = fieldData.ErrorInformation.Description
			}
		}
	}

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

func DeleteTicket(c *fiber.Ctx) error {
	database.DB.Where("id = ?", c.Params("id")).Delete(&models.DeliveryTicket{})
	database.DB.Where("ticket_id = ?", c.Params("id")).Delete(&models.FailedChecks{})
	return c.JSON(fiber.Map{
		"message": "ticket " + c.Params("id") + " deleted",
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

func deleteImage(filename string) error {
	err := os.Remove(fmt.Sprintf("./data/images/%s", filename))
	if err != nil {
		return err
	}

	return nil
}
