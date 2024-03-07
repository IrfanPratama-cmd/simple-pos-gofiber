package account

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"errors"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PostVerifyAccount godoc
// @Summary Verify account
// @Description Verify account
// @Param data body model.VerificationAccountAPI true "Payload"
// @Security TokenKey
// @Success 200 {object} model.UserAPI "Account verified"
// @Failure 401 {object} lib.Response "Unauthorized"
// @Failure 400 {object} lib.Response "Bad Request"
// @Failure 204 {object} lib.Response "No Content"
// @Failure 404 {object} lib.Response "Not Found"
// @Failure 409 {object} lib.Response "Conflict"
// @Failure 500 {object} lib.Response "Internal Server Error"
// @Router /account/verify-account [post]
// @Tags Account
func PostVerifyAccount(c *fiber.Ctx) error {
	api := new(model.VerificationAccountAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	// Check if verification code exists
	user := new(model.User)
	if err := db.Where("email = ? AND verification_code = ?", api.Email, api.VerificationCode).Take(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lib.ErrorBadRequest(c, "Invalid verification code")
		}
		return lib.ErrorInternal(c, err.Error())
	}

	// Check if verification code has expired
	expiration := user.VerificationExpiration
	if expiration != nil {
		expirationTime, err := time.Parse(time.RFC3339Nano, expiration.String())
		if err != nil {
			return lib.ErrorInternal(c, err.Error())
		}
		if time.Now().After(expirationTime) {
			return lib.ErrorBadRequest(c, "Verification code has expired")
		}
	}

	// Update user activation status and activated_at
	user.IsActivated = lib.Boolptr(true)
	now := time.Now()
	user.ActivatedAt = (*strfmt.DateTime)(&now)
	if err := db.Save(user).Error; err != nil {
		return lib.ErrorInternal(c, err.Error())
	}

	return lib.OK(c, user)
}
