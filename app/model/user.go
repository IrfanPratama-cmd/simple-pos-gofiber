package model

import (
	"api/app/lib"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model
type User struct {
	Base
	DataOwner
	UserAPI
	Asset *Asset `json:"asset" gorm:"foreignkey:AssetID"`
}

// UserAPI model
type UserAPI struct {
	AssetID                   *uuid.UUID       `json:"asset_id,omitempty" gorm:"type:varchar(36)"`
	IsOwner                   *bool            `json:"is_owner,omitempty"`
	Fullname                  *string          `json:"fullname,omitempty" gorm:"not null"`
	Username                  *string          `json:"username,omitempty" gorm:"type:varchar(191);index:idx_users_username_unique,unique,where:deleted_at is null;not null"`
	Email                     *string          `json:"email,omitempty" gorm:"type:varchar(191);index:idx_users_email_unique,unique,where:deleted_at is null;not null"`
	Mobile                    *string          `json:"mobile,omitempty"`
	Password                  *string          `json:"-" gorm:"not null"`
	Salt                      *string          `json:"-,omitempty"`
	IsPasswordSystemGenerated *bool            `json:"is_password_system_generated,omitempty"`
	PasswordLastChange        *strfmt.DateTime `json:"password_last_change,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	PasswordExpiration        *strfmt.DateTime `json:"password_expiration,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	IsActivated               *bool            `json:"is_activated,omitempty"`
	ActivatedAt               *strfmt.DateTime `json:"activated_at,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	ResetPasswordCode         *string          `json:"-"`
	ResetPasswordExpiration   *strfmt.DateTime `json:"reset_password_expiration,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	OTPEnabled                *bool            `json:"otp_enabled,omitempty"`
	OTPCode                   *string          `json:"-"`
	OTPExpiration             *strfmt.DateTime `json:"otp_expiration,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	VerificationCode          *string          `json:"-"`
	VerificationExpiration    *strfmt.DateTime `json:"verification_expiration,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	LastLogin                 *strfmt.DateTime `json:"last_login,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
}

// BeforeCreate Data
func (b *User) BeforeCreate(tx *gorm.DB) error {
	if nil != b.ID {
		return nil
	}
	id, e := uuid.NewRandom()
	now := strfmt.DateTime(time.Now())
	b.ID = &id
	b.CreatedAt = &now
	b.UpdatedAt = &now

	// Generate username if it's null
	if b.Username == nil {
		emailParts := strings.Split(*b.Email, "@")
		username := emailParts[0]
		b.Username = &username
	}

	return e
}

// UserData
type UserData struct {
	Username *string `json:"username,omitempty" example:"john.robert.doe" gorm:"type:varchar(191);index:idx_users_username_unique,unique,where:deleted_at is null;not null"`
	Password *string `json:"password" example:"password"  gorm:"not null"`
}

type UpdateProfile struct {
	AssetID         *uuid.UUID   `json:"asset_id,omitempty" swaggertype:"string" format:"uuid"`                                                  // Asset ID
	EmployeeCode    *string      `json:"employee_code,omitempty" example:"EMP-000001" gorm:"type:varchar(191);not null"`                         // Employee Code
	FirstName       *string      `json:"first_name,omitempty" example:"John" gorm:"type:varchar(191);not null"`                                  // First Name
	MiddleName      *string      `json:"middle_name,omitempty" example:"Robert" gorm:"type:varchar(191)"`                                        // Middle Name
	LastName        *string      `json:"last_name,omitempty" example:"Doe" gorm:"type:varchar(191)"`                                             // Last Name
	Gender          *string      `json:"gender,omitempty" example:"male" gorm:"type:varchar(191)"`                                               // Gender (e.g : 'male','female','rather_not_say')
	MaritalStatus   *string      `json:"marital_status,omitempty" example:"married" gorm:"type:varchar(191)"`                                    // Marital Status (e.g : 'married','unmarried','divorced')
	DateOfBirth     *strfmt.Date `json:"date_of_birth,omitempty" format:"date" swaggertype:"string" gorm:"type:date"`                            // Date Of Birth
	Mobile          *string      `json:"mobile,omitempty" example:"08123456789" gorm:"type:varchar(191)" validate:"omitempty,phone"`             // Mobile
	AlternateNumber *string      `json:"alternate_number,omitempty" example:"08123456789" gorm:"type:varchar(191)" validate:"omitempty,phone"`   // Alternate Number
	Email           *string      `json:"email,omitempty" example:"walk-in-customer@mail.com" gorm:"type:varchar(191)" validate:"required,email"` // Email
	ProvinceID      *uuid.UUID   `json:"province_id,omitempty" swaggertype:"string" format:"uuid"`                                               // Province ID
	CityID          *uuid.UUID   `json:"city_id,omitempty" swaggertype:"string" format:"uuid"`                                                   // City ID
	SubdistrictID   *uuid.UUID   `json:"subdistrict_id,omitempty" swaggertype:"string" format:"uuid"`                                            // Subdistrict ID
	Address         *string      `json:"address,omitempty" example:"Surakarta" gorm:"type:varchar(255)"`                                         // Address
	ZipCode         *string      `json:"zip_code,omitempty" example:"57422" gorm:"type:varchar(255)"`                                            // Zip Code
	User            *UserData    `json:"user,omitempty" gorm:"foreignKey:EmployeeID"`
}

// func (s User) Seed() {
// 	db := services.DB

// 	pass := "password"
// 	hashedPassword, _ := lib.HashPassword(pass)
// 	now := strfmt.DateTime(time.Now())

// 	var data User
// 	data.ID = lib.GenUUID()
// 	data.Fullname = lib.Strptr("Administrator")
// 	data.Username = lib.Strptr("admin")
// 	data.Email = lib.Strptr("admin@gmail.com")
// 	data.Password = &hashedPassword
// 	data.ActivatedAt = &now
// 	db.Create(&data)
// }

func (s User) Seed() *[]User {
	data := []User{}
	items := []string{
		"Administrator|admin|admin@gmail.com",
	}

	pass := "password"
	hashedPassword, _ := lib.HashPassword(pass)
	now := strfmt.DateTime(time.Now())

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		fullname := c[0]
		username := c[1]
		email := c[2]

		data = append(data, User{
			UserAPI: UserAPI{
				Fullname:    &fullname,
				Username:    &username,
				Email:       &email,
				Password:    &hashedPassword,
				IsActivated: lib.Boolptr(true),
				ActivatedAt: &now,
			},
		})
	}
	return &data
}
