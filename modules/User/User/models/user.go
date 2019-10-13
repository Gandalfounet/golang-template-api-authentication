package models

import (
	"github.com/jinzhu/gorm"
)

//User struct declaration
type User struct {
	gorm.Model

	Name     			string
	Email    			string `gorm:"type:varchar(100);unique_index"`
	Gender   			string `json:"Gender"`
	Password 			string `json:"Password"`
	Role 	 			string `json:"Role"`
	Status 				string `json:"Status"`
	ResetToken 			string `json:"ResetToken"`
	ResetTokenExpiracy  string `json:"ResetTokenExpiracy"`
	ValidationToken 	string `json:"ValidationToken"`
}