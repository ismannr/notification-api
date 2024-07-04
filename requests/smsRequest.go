package requests

type SmsOtpRequest struct {
	Msisdn      string `json:"msisdn"`
	Template    string `json:"template,omitempty"`
	Otp         string `json:"otp,omitempty"`
	TimeLimit   string `json:"time_limit,omitempty"`
	Challenge   string `json:"challenge,omitempty"`
	CallbackUrl string `json:"callback_url,omitempty"`
}

type VerifySmsOtpRequest struct {
	Msisdn    string `json:"msisdn"`
	Otp       string `json:"otp"`
	Challenge string `json:"challenge,omitempty"`
}
