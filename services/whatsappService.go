package services

import (
	"api-notification/requests"
	"api-notification/responses"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func SendNotification(w http.ResponseWriter, r *http.Request) {
	var request requests.NotificationRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responses.GlobalResponse("Invalid request payload", http.StatusBadRequest, nil, w)
		return
	}

	url := "https://api.verihubs.com/v1/whatsapp/message/send"

	payload := strings.NewReader("{\"msisdn\":\"%s\",\"content\":{\"body_params\":[\"string\"],\"header_params\":[\"string\"],\"button_param\":{\"url_param\":\"string\"}},\"lang_code\":\"en\",\"template_name\":\"%s\",\"callback_url\":\"https://your-url-callback.com\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func SendBulkNotification(w http.ResponseWriter, r *http.Request) {
	var request requests.BulkNotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responses.GlobalResponse("Invalid request payload", http.StatusBadRequest, nil, w)
		return
	}
	fmt.Println(request.Recipient)
	payload := map[string]interface{}{
		"msisdn": request.Recipient,
		"content": []map[string]interface{}{
			{
				"header_params": []string{"https://mystoragebucket.com/image.png"},
				"body_params":   []string{"John", "Our Website"},
				"button_param": map[string]string{
					"url_param": "content/promotion/244",
				},
			},
			{
				"header_params": []string{"https://mystoragebucket.com/image.png"},
				"body_params":   []string{"Doe", "Our Website"},
				"button_param": map[string]string{
					"url_param": "content/promotion/243",
				},
			},
		},
		"template_name": "transaction_daily_update",
		"lang_code":     "en",
		"callback_url":  "http://yourcallbackurl.com",
	}

	responses.GlobalResponse("test", 200, payload, w)
}
