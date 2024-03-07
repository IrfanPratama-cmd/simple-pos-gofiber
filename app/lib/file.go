package lib

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(c *fiber.Ctx) error {
	file, errFile := c.FormFile("thumbnail")
	if errFile != nil {
		log.Println("Error file = ", errFile)
	}

	var fileName string

	if file != nil {
		fileName = file.Filename
		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/thumbnail/%s", fileName))
		if errSaveFile != nil {
			log.Println("Upload Failed ", errFile)
		}
	} else {
		log.Println("Nothing file to uploading.")
	}

	c.Locals("thumbnail", fileName)

	return c.Next()
}
