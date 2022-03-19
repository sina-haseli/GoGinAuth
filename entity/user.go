package entity

import "time"

type User struct {
	ID       int    `json:"id" gorm:"primary_key;auto_increment"`
	Username string `json:"username" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"current_timestamp" json:"created_at"`
}
