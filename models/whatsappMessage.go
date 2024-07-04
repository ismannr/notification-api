package models

import "gorm.io/gorm"

type Message struct {
	//id nya menyesuaikan
	gorm.Model
	Title     string `json:"title"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
	Provider  string `json:"provider"`
	Status    string `json:"status"`
	SessionId string `json:"session_id"`
	IsBulk    bool
	BulkId    string
}
