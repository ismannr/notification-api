package responses

type SmsOtpResponse struct {
	Msisdn    string `json:"msisdn,omitempty"`
	Template  string `json:"template,omitempty"`
	Otp       string `json:"otp,omitempty"`
	TimeLimit string `json:"time_limit,omitempty"`
	Challenge string `json:"challenge,omitempty"`
}
