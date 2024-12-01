package domain

import "time"

type Users struct {
	User_id       uint   `gorm:"type:int;primarykey" json:"user_id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`	
}