/*
This package contains the errors that can be raised by the WhatsApp Cloud API or incoming errors from the webhook.
*/

package whatsapperrors

import (
	"fmt"
	"net/http"
)

// WhatsAppError is the base exception for all WhatsApp errors.
//
// - `'Cloud API Error Codes' on developers.facebook.com <https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes>`_.
// - `'Flow Error Codes' on developers.facebook.com <https://developers.facebook.com/docs/whatsapp/flows/reference/error-codes>`_.
type WhatsAppError struct {
	ErrorCode   int         `json:"error_code"`
	Message     string      `json:"message"`
	Details     string      `json:"details,omitempty"`
	FBTraceID   string      `json:"fbtrace_id,omitempty"`
	Href        string      `json:"href,omitempty"`
	RawResponse *http.Response
}

func (e *WhatsAppError) Error() string {
	return fmt.Sprintf("%T(message=%q, details=%q, code=%d)", e, e.Message, e.Details, e.ErrorCode)
}

// StatusCode returns the status code (if the error was raised from an API call, otherwise returns 0).
func (e *WhatsAppError) StatusCode() int {
	if e.RawResponse != nil {
		return e.RawResponse.StatusCode
	}
	return 0
}

// AuthorizationError is the base exception for all authorization errors.
type AuthorizationError struct {
	WhatsAppError
}

// AuthException represents an error where we were unable to authenticate the app user.
type AuthException struct {
	AuthorizationError
}

// APIMethod represents a capability or permissions issue.
type APIMethod struct {
	AuthorizationError
}

// PermissionDenied represents an error where the permission is either not granted or has been removed.
type PermissionDenied struct {
	AuthorizationError
}

// ExpiredAccessToken represents an error where the access token has expired.
type ExpiredAccessToken struct {
	AuthorizationError
}

// APIPermission represents an error where the permission is either not granted or has been removed.
type APIPermission struct {
	AuthorizationError
}

// ThrottlingError is the base exception for all rate limit errors.
type ThrottlingError struct {
	WhatsAppError
}

// ToManyAPICalls represents an error where the app has reached its API call rate limit.
type ToManyAPICalls struct {
	ThrottlingError
}

// RateLimitIssues represents an error where the WhatsApp Business Account has reached its rate limit.
type RateLimitIssues struct {
	ThrottlingError
}

// RateLimitHit represents an error where the Cloud API message throughput has been reached.
type RateLimitHit struct {
	ThrottlingError
}

// SpamRateLimitHit represents an error where the message failed to send because there are restrictions on how many messages can be sent from this phone number.
type SpamRateLimitHit struct {
	ThrottlingError
}

// TooManyMessages represents an error where too many messages were sent from the sender phone number to the same recipient phone number in a short period of time.
type TooManyMessages struct {
	ThrottlingError
}

// ToManyMessages is deprecated, use TooManyMessages instead.
type ToManyMessages = TooManyMessages

// IntegrityError is the base exception for all integrity errors.
type IntegrityError struct {
	WhatsAppError
}

// TemporarilyBlocked represents an error where the WhatsApp Business Account associated with the app has been restricted or disabled for violating a platform policy.
type TemporarilyBlocked struct {
	IntegrityError
}

// AccountLocked represents an error where the WhatsApp Business Account associated with the app has been restricted or disabled for violating a platform policy, or we were unable to verify data included in the request against data set on the WhatsApp Business Account (e.g, the two-step pin included in the request is incorrect).
type AccountLocked struct {
	IntegrityError
}

// SendMessageError is the base exception for all message errors.
type SendMessageError struct {
	WhatsAppError
}

// MessageUndeliverable represents an error where the message was unable to be delivered.
type MessageUndeliverable struct {
	SendMessageError
}

// ReEngagementMessage represents an error where more than 24 hours have passed since the recipient last replied to the sender number.
type ReEngagementMessage struct {
	SendMessageError
}

// UnsupportedMessageType represents an error where the message type is not supported.
type UnsupportedMessageType struct {
	SendMessageError
}

// MediaDownloadError represents an error where the media sent by the user could not be downloaded.
type MediaDownloadError struct {
	SendMessageError
}

// MediaUploadError represents an error where the media used in the message could not be uploaded.
type MediaUploadError struct {
	SendMessageError
}

// RecipientNotInAllowedList represents an error where the recipient is not in the allowed list when using test numbers.
type RecipientNotInAllowedList struct {
	SendMessageError
}

// InvalidParameter represents an error where the parameter passed is invalid.
type InvalidParameter struct {
	SendMessageError
}

// MissingRequiredParameter represents an error where a required parameter is missing.
type MissingRequiredParameter struct {
	SendMessageError
}

// TemplateParamCountMismatch represents an error where the number of variable parameter values included in the request did not match the number of variable parameters defined in the template.
type TemplateParamCountMismatch struct {
	SendMessageError
}

// TemplateNotExists represents an error where the template does not exist in the specified language or the template has not been approved.
type TemplateNotExists struct {
	SendMessageError
}

// TemplateTextTooLong represents an error where the template text is too long.
type TemplateTextTooLong struct {
	SendMessageError
}

// TemplateContentPolicyViolation represents an error where the template content violates a WhatsApp policy.
type TemplateContentPolicyViolation struct {
	SendMessageError
}

// TemplateParamValueInvalid represents an error where the variable parameter values are formatted incorrectly.
type TemplateParamValueInvalid struct {
	SendMessageError
}

// TemplatePaused represents an error where the template is paused due to low quality so it cannot be sent in a template message.
type TemplatePaused struct {
	SendMessageError
}

// TemplateDisabled represents an error where the template has been paused too many times due to low quality and is now permanently disabled.
type TemplateDisabled struct {
	SendMessageError
}

// FlowBlocked represents an error where the flow is in a blocked state.
type FlowBlocked struct {
	SendMessageError
}

// FlowThrottled represents an error where the flow is in a throttled state and 10 messages using this flow were already sent in the last hour.
type FlowThrottled struct {
	SendMessageError
}

// GenericError represents a generic error.
type GenericError struct {
	SendMessageError
}

// UnknownError represents an error where the message failed to send due to an unknown error.
type UnknownError struct {
	SendMessageError
}

// AccessDenied represents an error where the permission is either not granted or has been removed.
type AccessDenied struct {
	SendMessageError
}

// ServiceUnavailable represents an error where a service is temporarily unavailable.
type ServiceUnavailable struct {
	SendMessageError
}

// RecipientCannotBeSender represents an error where the sender and recipient phone numbers are the same.
type RecipientCannotBeSender struct {
	SendMessageError
}

// BusinessPaymentIssue represents an error where the message failed to send because there were one or more errors related to your payment method.
type BusinessPaymentIssue struct {
	SendMessageError
}

// FlowError is the base exception for all flow errors.
//
// Read more at `developers.facebook.com <https://developers.facebook.com/docs/whatsapp/flows/reference/error-codes>`_.
type FlowError struct {
	WhatsAppError
}

// FlowBlockedByIntegrity represents an error where we've identified an issue with integrity in your account and have prevented you from creating or publishing your Flow.
type FlowBlockedByIntegrity struct {
	FlowError
}

// FlowUpdatingError represents an error where the flow failed to update.
type FlowUpdatingError struct {
	FlowError
}

// FlowPublishingError represents an error where the flow failed to publish.
type FlowPublishingError struct {
	FlowError
}

// FlowDeprecatingError represents an error where the flow failed to deprecate.
type FlowDeprecatingError struct {
	FlowError
}

// FlowDeletingError represents an error where the flow failed to delete.
type FlowDeletingError struct {
	FlowError
}