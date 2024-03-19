package Whatsapp

type WhatsApp struct {
	PhoneNumberID string
	BearerToken   string
	Version       string
}

func NewWhatsApp(phoneNumberID, bearerToken, version string) *WhatsApp {
	return &WhatsApp{
		PhoneNumberID: phoneNumberID,
		BearerToken:   bearerToken,
		Version:       version,
	}
}
