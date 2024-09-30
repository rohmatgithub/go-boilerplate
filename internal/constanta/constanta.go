package constanta

type key string

const (
	TokenHeaderNameConstanta        = "Authorization"
	TokenNexposNameConstanta        = "X-AUTH-TOKEN-NEXPOS"
	RequestIDConstanta              = "X-Request-ID"
	IPAddressConstanta              = "X-Forwarded-For"
	SourceConstanta                 = "X-Source"
	ApplicationContextConstanta key = "application_context"
)
