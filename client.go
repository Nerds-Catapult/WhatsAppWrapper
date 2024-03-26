package whatsAppCloud_go

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Nerds-Catapult/whatsAppCloud-go/config"
	"net/http"
)

type Client struct {
	config        *config.Config
	httpClient    *http.Client
	W             Endpoint
	clientID      string
	requestHooks  []RequestHook
	responseHooks []ResponseHook
	credentials   AuthCredentials
}

func (c Client) Authenticate(ctx context.Context, credentials AuthCredentials) error {
	if credentials.AccessToken == "" || credentials.PhoneNumberID == "" || credentials.BusinessID == "" {
		return errors.New("invalid credentials provided")
	}
	c.credentials = credentials
	req, err := http.NewRequestWithContext(ctx, "POST", "https://graph.facebook.com/v13.0/"+credentials.PhoneNumberID+"/messages", nil)
	if err != nil {
		return err
	}

	// Set the necessary headers for authentication
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respError struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&respError); err != nil {
			return errors.New("authentication failed with status code: " + resp.Status)
		}
		return errors.New("authentication failed: " + respError.Error.Message)
	}

	return nil
}

func (c Client) LoginWithQRCode(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) LoginWithSessionCredentials(ctx context.Context, sessionCredentials ...string) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) SendMessage(ctx context.Context, messageType string, content interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) ReceiveMessage(ctx context.Context) (Message, error) {
	//TODO implement me
	panic("implement me")
}

func (c Client) ManageMedia(ctx context.Context, mediaType string, action string, media interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) SetupWebhook(ctx context.Context, url string) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) HandleError(ctx context.Context, err error) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) RespectRateLimit(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) RetrieveAnalytics(ctx context.Context, metrics string) (Analytics, error) {
	//TODO implement me
	panic("implement me")
}

func (c Client) AddReceivers(ctx context.Context, contacts ...string) error {
	//TODO implement me
	panic("implement me")
}

func NewClient(w Endpoint, httpClient *http.Client, config *config.Config, clientID string) Endpoint {
	return &Client{
		W:          w,
		httpClient: httpClient,
		config:     config,
		clientID:   clientID,
	}
}
