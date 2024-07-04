package requests

type NotificationRequest struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
	Provider  string `json:"provider"`
}

type BulkNotificationRequest struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Recipient []string `json:"recipient"`
	Provider  []string `json:"provider"`
}
