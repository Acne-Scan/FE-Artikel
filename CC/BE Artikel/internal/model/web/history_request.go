package web

type HistoryRequest struct {
	HistoryId int    `json:"history_id" validate:"required"`
	Image     string `json:"image"`
	Result    string `json:"result"`
	User_id   uint   `json:"user_id"`
}