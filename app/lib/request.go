package lib

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// VALIDATOR validate request body
var VALIDATOR *validator.Validate = validator.New()

// init Register custom validation function
func init() {
	VALIDATOR.RegisterValidation("phone", validatePhone)
	VALIDATOR.RegisterValidation("email", validateEmail)
	VALIDATOR.RegisterValidation("website", validateWebsite)
}

// Custom validation function for phone number format
func validatePhone(fl validator.FieldLevel) bool {
	phoneRegex := regexp.MustCompile(`^\d{10,12}$`)
	return phoneRegex.MatchString(fl.Field().String())
}

// Custom validation function for email format
func validateEmail(fl validator.FieldLevel) bool {
	emailRegex := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	return emailRegex.MatchString(fl.Field().String())
}

// Custom validation function for website format
func validateWebsite(fl validator.FieldLevel) bool {
	websiteRegex := regexp.MustCompile(`^(http|https):\/\/[^\s/$.?#].[^\s]*$`)
	return websiteRegex.MatchString(fl.Field().String())
}

// ClaimsJWT func
func ClaimsJWT(accesToken *string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(*accesToken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	timeNow := time.Now().Unix()
	timeSessions := int64(claims["exp"].(float64))
	if timeSessions < timeNow {
		return claims, err
	}
	return claims, nil
}

// GetXUserID provide user id from the authentication token
// func GetXUserID(c *fiber.Ctx) *uuid.UUID {
// 	authData, ok := c.Locals("auth").(model.ResponseAuthenticate)
// 	if ok && authData.UserID != nil {
// 		return StringToUUID(*authData.UserID)
// 	}
// 	return nil
// }

func GetXUserID(c *fiber.Ctx) *uuid.UUID {
	userID := c.Locals("userID")
	id := userID.(string)
	if id != "" {
		if current, err := uuid.Parse(id); nil == err {
			return &current
		}
	}

	return nil
}

// BodyParser with validation
func BodyParser(c *fiber.Ctx, payload interface{}) error {
	if err := c.BodyParser(payload); nil != err {
		return err
	}
	return VALIDATOR.Struct(payload)
}
