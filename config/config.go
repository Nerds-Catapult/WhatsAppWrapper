package config

import "net/http"

const (
	APIUrl      string = "https://api.whatsapp-cloud.com/v1"
	APIEndpoint string = "api"
	ApiVersion         = "v1"
)

type Config struct {
	HttpClient        *http.Client
	Url               *string
	ApiVersion        string
	AccessToken       string
	PhoneNumberID     string
	BusinessAccountID string
}

type Credentials struct {
}
