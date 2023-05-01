package model

import "time"


type Detail struct {
	Id       int64  `json:"id" gorm:"primary_key"`
	Title   string `json:"title" gorm:"type:varchar(100); not null"`
	MyPassword string `json:"mypassword" gorm:"type:varchar(100); not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"update_at"`
	User	User `json:"user" gorm:"foreignkey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId int64 `json:"user_id" gorm:"not null"`
}

type DetailResponse struct {
	Id       int64  `json:"id" gorm:"primary_key"`
	Title   string `json:"title" gorm:"type:varchar(100); not null"`
	MyPassword string `json:"mypassword" gorm:"type:varchar(100); not null"`
}