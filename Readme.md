# WhatsAppCloud-go

## Introduction

WhatsAppCloud-go is a Golang wrapper library for the WhatsApp Cloud API. It simplifies the process of integrating WhatsApp messaging capabilities into Golang applications. This library aims to provide a straightforward and idiomatic way to interact with the WhatsApp Cloud API.

## Features

The library will cover the following functionalities of the WhatsApp Cloud API:

- **Message Sending**: Functions to send various message types, including text, images, and templates.
- **Message Receiving**: Handlers to receive and process incoming messages.
- **Media Management**: Methods to upload and manage media attachments.
- **Webhooks**: Support for setting up webhooks to receive event-driven notifications.
- **Authentication**: Utilities for handling authentication and token management.
- **Error Handling**: Comprehensive error handling that aligns with Golang's error philosophy.
- **Rate Limiting**: Tools to respect API rate limits and handle retries.
- **Conversation Analytics**: Functions to retrieve conversation metrics and analytics⁴.

## Getting Started

### Prerequisites

- Go 1.15 or higher
- A WhatsApp Business Account
- Access to WhatsApp Cloud API

### Installation

To install WhatsAppCloud-go, use the following command:

```bash
go get github.com/Nerds-Catapult/whatsappCloud-go
```

### Usage

Here's a quick example of how to send a text message using WhatsAppCloud-go:

```go
package main

import (
    "github.com/Nerds-Catapult/whatsappCloud-go/whatsapp"
    "log"
)

func main() {
    client := whatsapp.NewClient("your-access-token")
    message := whatsapp.TextMessage{
        To:   "recipient-number",
        Body: "Hello, World!",
    }

    if err := client.SendMessage(message); err != nil {
        log.Fatalf("Error sending message: %v", err)
    }
}
```
## Configuration

Before you can start using the WhatsAppCloud-go library, you need to obtain an Access Token, Phone number ID, and Business ID. These credentials are essential for authenticating and interacting with the WhatsApp Cloud API.

### Register as a Meta Developer

To get started, you'll need to register as a Meta Developer:

1. Visit the Meta Developer Registration page and follow the instructions to register.
2. Once registered, create an application by following the guide on Creating an App.
3. Configure your application to enable access to the WhatsApp Business Cloud API.

### Manage Your Applications

You can manage your applications and their settings at the Facebook App Dashboard.

### Testing Your Setup

It's a good idea to test your setup before integrating the API into your application:

1. Go to the WhatsApp Developer Dashboard.
2. Use the provided interface to send a test message to your phone number.
3. Ensure that you receive the test message, confirming that your setup is correct.

By completing these steps, you'll be ready to use the WhatsAppCloud-go library in your Golang applications.



## Contributing

Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
```

