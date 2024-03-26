package whatsAppCloud_go

import "context"

type Endpoint interface {
	Authenticate(ctx context.Context, credentials AuthCredentials) error
	LoginWithQRCode(ctx context.Context) error
	LoginWithSessionCredentials(ctx context.Context, sessionCredentials ...string) error

	SendMessage(ctx context.Context, messageType string, content interface{}) error
	ReceiveMessage(ctx context.Context) (Message, error)
	ManageMedia(ctx context.Context, mediaType string, action string, media interface{}) error
	SetupWebhook(ctx context.Context, url string) error
	HandleError(ctx context.Context, err error) error
	RespectRateLimit(ctx context.Context) error
	RetrieveAnalytics(ctx context.Context, metrics string) (Analytics, error)
	AddReceivers(ctx context.Context, contacts ...string) error
}
