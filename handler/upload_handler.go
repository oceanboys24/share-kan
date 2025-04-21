package handler

import (
	"share-kan/utils"

	"github.com/gofiber/fiber/v2"
)




func UploadHandler(c *fiber.Ctx) error {
	file ,err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message" : "cannot read file",
		})
	}


	srcFile , err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "cannot open file",
		})
	}	
	
	resultDownload , err := utils.UploadFile("share-kan",file.Filename, srcFile, file.Size, file.Header.Get("content-type"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message " : "Upload Failed",
		})
	}


	return c.Status(200).JSON(fiber.Map{
		"message " : "Success Upload",
		"link" : resultDownload.String(),
	})
}