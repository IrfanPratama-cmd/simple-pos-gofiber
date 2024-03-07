package account

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"sync"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
)

// PostRegister godoc
// @Summary Registration
// @Description Registration
// @Param data body model.RegistrationAPI true "Payload"
// @Security TokenKey
// @Success 200 {object} lib.Response "registered"
// @Failure 401 {object} lib.Response "Unauthorized"
// @Failure 400 {object} lib.Response "Bad Request"
// @Failure 204 {object} lib.Response "No Content"
// @Failure 404 {object} lib.Response "Not Found"
// @Failure 409 {object} lib.Response "Conflict"
// @Failure 500 {object} lib.Response "Internal Server Error"
// @Router /account/register [post]
// @Tags Account
func PostRegisterAccount(c *fiber.Ctx) error {
	api := &model.RegistrationAPI{}
	if err := lib.BodyParser(c, api); err != nil {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())

	// Generate unique verification code
	verificationCode := lib.RandomNumber(6)

	// Set verification expiration time
	verificationExpiration := time.Now().Add(5 * time.Minute)

	var email model.User
	checkEmail := db.Model(&model.User{}).Where(`email = ?`, api.Email).First(&email)

	if checkEmail.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is already used",
		})
	}

	// Create user
	user := &model.User{
		UserAPI: model.UserAPI{
			Fullname:               api.Fullname,
			IsOwner:                lib.Boolptr(true),
			VerificationCode:       lib.Strptr(verificationCode),
			VerificationExpiration: lib.DateTimeptr(strfmt.DateTime(verificationExpiration)),
			Email:                  api.Email,
			Password:               lib.Strptr(lib.RandomChars(10)),
		},
	}

	var passwordHash string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		passwordHash, _ = lib.HashPassword(*api.Password)
	}()

	wg.Wait()
	user.Password = lib.Strptr(passwordHash)
	if err := db.Save(user).Error; err != nil {
		return lib.ErrorInternal(c, err.Error())
	}

	sendVerificationEmail(*api.Email, verificationCode)

	var contact model.Customer
	contact.CustomerName = api.Fullname
	contact.Email = api.Email
	contact.UserID = user.ID
	db.Create(&contact)

	return lib.OK(c)
}

func sendVerificationEmail(toEmail, verificationCode string) {
	from := mail.Address{"Developer Irfan", os.Getenv("SMTP_USERNAME")}
	to := mail.Address{"Recipient Name", toEmail}
	subject := "Email Verification"

	body := fmt.Sprintf("Your verification code is: %s", verificationCode)

	message := "From: " + from.String() + "\r\n" +
		"To: " + to.String() + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))

	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("SMTP_USERNAME"), []string{to.Address}, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
