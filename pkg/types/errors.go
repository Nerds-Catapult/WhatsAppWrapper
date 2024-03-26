package types

type WhatsAppError struct {
	ErrorCode int
	Message   string
	Details   string
	FbtraceID string
	Href string
	RawResponse interface{}
}


func NewWhatsAppError(errorcode int, message, details, fbtraceID, href string, rawResponse interface{}) *WhatsAppError {
	return &WhatsAppError{
		ErrorCode: errorcode,
		Message: message,
		Details: details,
		FbtraceID: fbtraceID,
		Href: href,
		RawResponse: rawResponse,
	}
}

// AuthorizationError represents the base exception for all authorization errors.
type AuthorizationError WhatsAppError

// AuthException represents an exception for authentication failures.
type AuthException AuthorizationError

// APIMethod represents an exception for capability or permissions issues.
type APIMethod AuthorizationError

// PermissionDenied represents an exception for denied permissions.
type PermissionDenied AuthorizationError

// ExpiredAccessToken represents an exception for expired access tokens.
type ExpiredAccessToken AuthorizationError

// APIPermission represents an exception for API permissions.
type APIPermission AuthorizationError

// ThrottlingError represents the base exception for all rate limit errors.
type ThrottlingError WhatsAppError

// ToManyAPICalls represents an exception for reaching API call rate limits.
type ToManyAPICalls ThrottlingError

// RateLimitIssues represents an exception for rate limit issues.
type RateLimitIssues ThrottlingError

// RateLimitHit represents an exception for reaching Cloud API message throughput limits.
type RateLimitHit ThrottlingError

// SpamRateLimitHit represents an exception for reaching spam rate limits.
type SpamRateLimitHit ThrottlingError

// TooManyMessages represents an exception for sending too many messages.
type TooManyMessages ThrottlingError

// IntegrityError represents the base exception for all integrity errors.
type IntegrityError WhatsAppError

// TemporarilyBlocked represents an exception for temporarily blocked WhatsApp Business Accounts.
type TemporarilyBlocked IntegrityError

// AccountLocked represents an exception for locked WhatsApp Business Accounts.
type AccountLocked IntegrityError

// SendMessageError represents the base exception for all message errors.
type SendMessageError WhatsAppError

// MessageUndeliverable represents an exception for undeliverable messages.
type MessageUndeliverable SendMessageError

// ReEngagementMessage represents an exception for re-engagement message limitations.
type ReEngagementMessage SendMessageError

// UnsupportedMessageType represents an exception for unsupported message types.
type UnsupportedMessageType SendMessageError

// MediaDownloadError represents an exception for media download failures.
type MediaDownloadError SendMessageError

// MediaUploadError represents an exception for media upload failures.
type MediaUploadError SendMessageError

// RecipientNotInAllowedList represents an exception for recipients not in the allowed list.
type RecipientNotInAllowedList SendMessageError

// InvalidParameter represents an exception for invalid parameters.
type InvalidParameter SendMessageError

// MissingRequiredParameter represents an exception for missing required parameters.
type MissingRequiredParameter SendMessageError

// TemplateParamCountMismatch represents an exception for template parameter count mismatches.
type TemplateParamCountMismatch SendMessageError

// TemplateNotExists represents an exception for non-existent templates.
type TemplateNotExists SendMessageError

// TemplateTextTooLong represents an exception for overly long template text.
type TemplateTextTooLong SendMessageError

// TemplateContentPolicyViolation represents an exception for template content policy violations.
type TemplateContentPolicyViolation SendMessageError

// TemplateParamValueInvalid represents an exception for invalid template parameter values.
type TemplateParamValueInvalid SendMessageError

// TemplatePaused represents an exception for paused templates.
type TemplatePaused SendMessageError

// TemplateDisabled represents an exception for disabled templates.
type TemplateDisabled SendMessageError

// FlowError represents the base exception for all flow errors.
type FlowError WhatsAppError

// FlowBlockedByIntegrity represents an exception for flow blockages due to integrity issues.
type FlowBlockedByIntegrity FlowError

// FlowUpdatingError represents an exception for flow updating failures.
type FlowUpdatingError FlowError

// FlowPublishingError represents an exception for flow publishing failures.
type FlowPublishingError FlowError

// FlowDeprecatingError represents an exception for flow deprecation failures.
type FlowDeprecatingError FlowError

// FlowDeletingError represents an exception for flow deletion failures.
type FlowDeletingError FlowError
