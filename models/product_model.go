package models

import "gorm.io/gorm"

type UserRequest struct {
	gorm.Model
	Command string `json:"command"`
	Result  int `json:"result"`
	Question string `json:"question"`
}
