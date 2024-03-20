package models

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type User struct {
	gorm.Model
	Username string
	Address  string
	Email    string
}

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
	)
	v.Check(user.Address,
		revel.Required{},
	)
	v.Email(user.Email)
}