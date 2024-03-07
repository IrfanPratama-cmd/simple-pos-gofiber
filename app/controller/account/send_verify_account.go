package account

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
)

// PostSendVerifyAccountCode godoc
// @Summary Resend Code Verify account
// @Description Resend Code Verify account
// @Param data body model.SendVerificationAccountAPI true "Payload"
// @Security TokenKey
// @Success 200 {object} model.UserAPI "Account verified"
// @Failure 401 {object} lib.Response "Unauthorized"
// @Failure 400 {object} lib.Response "Bad Request"
// @Failure 204 {object} lib.Response "No Content"
// @Failure 404 {object} lib.Response "Not Found"
// @Failure 409 {object} lib.Response "Conflict"
// @Failure 500 {object} lib.Response "Internal Server Error"
// @Router /account/send-verify-account-code [post]
// @Tags Account
func PostSendVerifyAccountCode(c *fiber.Ctx) error {
	api := new(model.SendVerificationAccountAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	// Check if verification code exists
	user := new(model.User)
	if err := db.Where("email = ?", api.Email).First(user).Error; err != nil {
		return lib.ErrorNotFound(c, "Email not found")
	}

	// Generate unique verification code
	verificationCode := lib.RandomNumber(6)

	// Set verification expiration time
	verificationExpiration := time.Now().Add(5 * time.Minute)

	// Update user's verification code and expiration
	user.VerificationCode = lib.Strptr(verificationCode)
	user.VerificationExpiration = lib.DateTimeptr(strfmt.DateTime(verificationExpiration))

	// Save the changes
	if err := db.Save(user).Error; err != nil {
		return lib.ErrorInternal(c, err.Error())
	}

	sendVerificationEmail(*api.Email, verificationCode)

	return lib.OK(c)
}
