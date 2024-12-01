package domain

import "time"

type History struct {
	HistoryId int       `json:"history_id"`
	Image     string    `json:"image"`
	Result    string    `json:"result"`
	User_id   uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
