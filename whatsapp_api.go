package whatsAppCloud_go

import (
	"errors"
	"net/http"
)

type API struct {
	client *Client
}

// APIOptions are the options for creating an API struct

type APIOptions struct {
	ClientID   string
	HTTPClient *http.Client
}

func NewAPI(opts APIOptions) (*API, error) {
	if opts.ClientID == "" {
		return nil, errors.New("missing client ID")
	}

	if opts.HTTPClient == nil {
		opts.HTTPClient = http.DefaultClient
	}
	return &API{
		client: &Client{
			httpClient: opts.HTTPClient,
			W: &Client{
				httpClient: opts.HTTPClient,
			},
		},
	}, nil
}

// SetClientID -->sets the client ID
func (w *API) SetClientID(clientID string) {
	w.client.clientID = clientID
}

// ClientID  returns the client ID
func (w *API) ClientID() string {
	return w.client.clientID
}
