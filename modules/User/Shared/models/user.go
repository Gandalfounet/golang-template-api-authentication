package models

import (
	"github.com/jinzhu/gorm"
	"time"
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
	ResetTokenExpiry  time.Time `json:"ResetTokenExpiry"`
	ValidationToken 	string `json:"ValidationToken"`
}