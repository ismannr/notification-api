package services

import (
	"api-notification/requests"
	"api-notification/responses"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func SendOtpBySMS(w http.ResponseWriter, r *http.Request) {
	var request requests.SmsOtpRequest
	url := "https://api.verihubs.com/v1/otp/send"
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responses.GlobalResponse("Invalid request payload", http.StatusBadRequest, nil, w)
		log.Println(err.Error())
		return
	}

	request.CallbackUrl = os.Getenv("CALLBACK_SMS_URL")

	payloadBytes, err := json.Marshal(request)
	if err != nil {
		responses.GlobalResponse("Error encoding request payload", http.StatusInternalServerError, nil, w)
		log.Println(err.Error())
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		responses.GlobalResponse("Error creating request", http.StatusInternalServerError, nil, w)
		log.Println(err.Error())
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Add("App-ID", os.Getenv("APP_ID"))
	req.Header.Add("API-Key", os.Getenv("API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		responses.GlobalResponse("Error sending request", http.StatusInternalServerError, nil, w)
		log.Println(err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		responses.GlobalResponse("Error reading response", http.StatusInternalServerError, nil, w)
		log.Println(err.Error())
		return
	}
	var apiResponse map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		responses.GlobalResponse("Error decoding API response", http.StatusInternalServerError, nil, w)
		log.Println(err.Error())
		return
	}
	responses.GlobalResponse("Successfully connecting to the API", res.StatusCode, apiResponse, w)
}

func VerifySmsOtp(w http.ResponseWriter, r *http.Request) {
	var request requests.VerifySmsOtpRequest
	url := "https://api.verihubs.com/v1/otp/verify"
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responses.GlobalResponse("Invalid request payload", http.StatusBadRequest, nil, w)
		log.Println(err.Error())
		return
	}

	payloadBytes, err := json.Marshal(request)
	if err != nil {
		responses.GlobalResponse("Error encoding request payload", http.StatusInternalServerError, nil, w)
		log.Println(err.Error())
		return
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))

	req.Header.Add("accept", "application/json")
	req.Header.Add("App-ID", os.Getenv("APP_ID"))
	req.Header.Add("API-Key", os.Getenv("API_KEY"))
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		responses.GlobalResponse("Error decoding API response", http.StatusInternalServerError, nil, w)
		log.Println(err.Error())
		return
	}
	responses.GlobalResponse("Successfully connecting to the API", res.StatusCode, apiResponse, w)
}
