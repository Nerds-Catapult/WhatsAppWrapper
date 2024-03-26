package whatsapp

type Endpoint interface {
	SendMessage(messageType string, content interface{}) error
	ReceiveMessage() (Message, error)
	ManageMedia(mediaType string, action string, media interface{}) error
	SetupWebhook(url string) error
	Authenticate(credentials AuthCredentials) error
	HandleError(err error) error
	RespectRateLimit() error
	RetrieveAnalytics(metrics string) (Analytics, error)
}

// Message struct to represent a generic message.
//type Message struct {
//	ID      string
//	Content interface{}
//}
//
//// AuthCredentials struct to hold authentication details.
//type AuthCredentials struct {
//	ClientID     string
//	ClientSecret string
//	Token        string
//}
//
//// Analytics struct to represent conversation analytics.
//type Analytics struct {
//	Metrics map[string]interface{}
//}
