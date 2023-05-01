package model

import "time"

type User struct {
	Id       int64  `json:"id" gorm:"primary_key"`
	Email	string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UserResponse struct {
	Id       int64  `json:"id" gorm:"primary_key"`
	Email	string `json:"email" gorm:"type:varchar(100);unique_index"`
}
