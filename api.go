package whatsAppCloud_go

import (
	"context"
	"net/http"
)

type AuthCredentials struct {
	AccessToken   string
	PhoneNumberID string
	BusinessID    string
}

type Message struct {
	ID      string
	Content interface{}
}

type Analytics struct {
	Metrics map[string]interface{}
}

type (
	RequestHook  func(ctx context.Context, request *http.Request) error
	ResponseHook func(ctx context.Context, response *http.Response) error
)
