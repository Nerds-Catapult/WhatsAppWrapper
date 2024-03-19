package whatsAppCloud_go

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}
