package models

import (
	"go-rest-crud/utils"
	"time"
)

type User struct {
	ID        uint32    `gorm:"column:id;primaryKey" json:"id,omitempty"`
	FistName  string    `gorm:"column:first_name;type:varchar(50)" json:"first_name"`
	LastName  string    `gorm:"column:last_name;type:varchar(50)" json:"last_name"`
	Gender    string    `gorm:"column:gender;type:char(10)" json:"gender"`
	BirthDate string    `gorm:"column:birth_date;type:date" json:"birth_date"`
	Email     string    `gorm:"column:email;unique;type:varchar(100)" json:"email"`
	Phone     string    `gorm:"column:phone;type:char(20)" json:"phone"`
	IsActive  *bool      `gorm:"column:is_active;default:true" json:"is_active,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

func (u *User) Validation() (s string) {

	if !utils.IsAlphabet(u.FistName) {
		s = "first name only contains alphabet and cannot be empty"
		return
	}

	if !utils.IsAlphabet(u.LastName) {
		s = "last name only contains alphabet and cannot be empty"
		return
	}

	if !utils.IsBirthDateValid(u.BirthDate) {
		s = "birth date format is not valid. valid format: yyyy-mm-dd"
		return
	}

	if !utils.IsEmailValid(u.Email) {
		s = "email address is not valid"
		return
	}

	if !utils.IsPhoneValid(u.Phone) {
		s = "phone number is not valid. phone number only contains numeric. min 7 digit, max 16 digit"
		return
	}

	return
}
