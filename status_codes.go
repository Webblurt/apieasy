package apieasy

type StatusCode int

const (
	// 2xx Success
	OK                        StatusCode = 200
	Created                   StatusCode = 201
	Accepted                  StatusCode = 202
	NonAuthritativeInformaton StatusCode = 203
	NoContent                 StatusCode = 204
	ResetContent              StatusCode = 205
	PartialContent            StatusCode = 206
	MultiStatus               StatusCode = 207
	AlreadyReported           StatusCode = 208
	IMUsed                    StatusCode = 226

	// 3xx Redirection

	MultipleChoices   StatusCode = 300
	MovedPermanently  StatusCode = 301
	Found             StatusCode = 302
	SeeOther          StatusCode = 303
	NotModified       StatusCode = 304
	UseProxy          StatusCode = 305
	TemporaryRedirect StatusCode = 307
	PermanentRedirect StatusCode = 308

	// 4xx Client Error
	// 400 - 499

	BadRequest                   StatusCode = 400
	Unauthorized                 StatusCode = 401
	PaymentRequired              StatusCode = 402
	Forbidden                    StatusCode = 403
	NotFound                     StatusCode = 404
	MethodNotAllowed             StatusCode = 405
	NotAcceptable                StatusCode = 406
	ProxyAuthRequired            StatusCode = 407
	RequestTimeout               StatusCode = 408
	Conflict                     StatusCode = 409
	Gone                         StatusCode = 410
	LengthRequired               StatusCode = 411
	PreconditionFailed           StatusCode = 412
	RequestEntityTooLarge        StatusCode = 413
	RequestURITooLong            StatusCode = 414
	UnsupportedMediaType         StatusCode = 415
	RequestedRangeNotSatisfiable StatusCode = 416
	ExpectationFailed            StatusCode = 417
	Teapot                       StatusCode = 418
	MisdirectedRequest           StatusCode = 421
	UnprocessableEntity          StatusCode = 422
	Locked                       StatusCode = 423
	FailedDependency             StatusCode = 424
	UpgradeRequired              StatusCode = 426
	PreconditionRequired         StatusCode = 428
	TooManyRequests              StatusCode = 429
	RequestHeaderFieldsTooLarge  StatusCode = 431
	UnavailableForLegalReasons   StatusCode = 451
	ClientClosedRequest          StatusCode = 499

	// 5xx Server Error
	// 500 - 599

	InternalServerError           StatusCode = 500
	NotImplemented                StatusCode = 501
	BadGateway                    StatusCode = 502
	ServiceUnavailable            StatusCode = 503
	GatewayTimeout                StatusCode = 504
	HttpVersionNotSupported       StatusCode = 505
	VariantAlsoNegotiates         StatusCode = 506
	InsufficientStorage           StatusCode = 507
	LoopDetected                  StatusCode = 508
	NotExtended                   StatusCode = 510
	NetworkAuthenticationRequired StatusCode = 511
	NetworkConnectTimeout         StatusCode = 599
)

var StatusText = map[StatusCode]string{
	// 2xx Success
	OK:                        "OK",
	Created:                   "Created",
	Accepted:                  "Accepted",
	NonAuthritativeInformaton: "Non Authritative Informaton",
	NoContent:                 "No Content",
	ResetContent:              "Reset Content",
	PartialContent:            "Partial Content",
	MultiStatus:               "Multi Status",
	AlreadyReported:           "Already Reported",
	IMUsed:                    "IM Used",

	// 3xx Redirection

	MultipleChoices:   "Multiple Choices",
	MovedPermanently:  "Moved Permanently",
	Found:             "Found",
	SeeOther:          "See Other",
	NotModified:       "Not Modified",
	UseProxy:          "Use Proxy",
	TemporaryRedirect: "Temporary Redirect",
	PermanentRedirect: "Permanent Redirect",

	// 4xx Client Error
	// 400 - 499

	BadRequest:                   "Bad Request",
	Unauthorized:                 "Unauthorized",
	PaymentRequired:              "Payment Required",
	Forbidden:                    "Forbidden",
	NotFound:                     "Not Found",
	MethodNotAllowed:             "Method Not Allowed",
	NotAcceptable:                "Not Acceptable",
	ProxyAuthRequired:            "Proxy Auth Required",
	RequestTimeout:               "Request Timeout",
	Conflict:                     "Conflict",
	Gone:                         "Gone",
	LengthRequired:               "Length Required",
	PreconditionFailed:           "Precondition Failed",
	RequestEntityTooLarge:        "Request Entity Too Large",
	RequestURITooLong:            "Request URI Too Long",
	UnsupportedMediaType:         "Unsupported Media Type",
	RequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
	ExpectationFailed:            "Expectation Failed",
	Teapot:                       "Teapot",
	MisdirectedRequest:           "Misdirected Request",
	UnprocessableEntity:          "Unprocessable Entity",
	Locked:                       "Locked",
	FailedDependency:             "Failed Dependency",
	UpgradeRequired:              "Upgrade Required",
	PreconditionRequired:         "Precondition Required",
	TooManyRequests:              "Too Many Requests",
	RequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
	UnavailableForLegalReasons:   "Unavailable For Legal Reasons",
	ClientClosedRequest:          "Client Closed Request",

	// 5xx Server Error
	// 500 - 599

	InternalServerError:           "Internal Server Error",
	NotImplemented:                "Not Implemented",
	BadGateway:                    "Bad Gateway",
	ServiceUnavailable:            "Service Unavailable",
	GatewayTimeout:                "Gateway Timeout",
	HttpVersionNotSupported:       "HTTP Version Not Supported",
	VariantAlsoNegotiates:         "Variant Also Negotiates",
	InsufficientStorage:           "Insufficient Storage",
	LoopDetected:                  "Loop Detected",
	NotExtended:                   "Not Extended",
	NetworkAuthenticationRequired: "Network Authentication Required",
	NetworkConnectTimeout:         "Network Connect Timeout",
}

var (
	Success = struct {
		Default                   StatusCode
		OK                        StatusCode
		Created                   StatusCode
		Accepted                  StatusCode
		NonAuthritativeInformaton StatusCode
		NoContent                 StatusCode
		ResetContent              StatusCode
		PartialContent            StatusCode
		MultiStatus               StatusCode
		AlreadyReported           StatusCode
		IMUsed                    StatusCode
	}{
		Default:                   OK,
		OK:                        OK,
		Created:                   Created,
		Accepted:                  Accepted,
		NonAuthritativeInformaton: NonAuthritativeInformaton,
		NoContent:                 NoContent,
		ResetContent:              ResetContent,
		PartialContent:            PartialContent,
		MultiStatus:               MultiStatus,
		AlreadyReported:           AlreadyReported,
		IMUsed:                    IMUsed,
	}

	Redirection = struct {
		Default           StatusCode
		MultipleChoices   StatusCode
		MovedPermanently  StatusCode
		Found             StatusCode
		SeeOther          StatusCode
		NotModified       StatusCode
		UseProxy          StatusCode
		TemporaryRedirect StatusCode
		PermanentRedirect StatusCode
	}{
		Default:           MovedPermanently,
		MultipleChoices:   MultipleChoices,
		MovedPermanently:  MovedPermanently,
		Found:             Found,
		SeeOther:          SeeOther,
		NotModified:       NotModified,
		UseProxy:          UseProxy,
		TemporaryRedirect: TemporaryRedirect,
		PermanentRedirect: PermanentRedirect,
	}

	ClientError = struct {
		Default                      StatusCode
		BadRequest                   StatusCode
		Unauthorized                 StatusCode
		PaymentRequired              StatusCode
		Forbidden                    StatusCode
		NotFound                     StatusCode
		MethodNotAllowed             StatusCode
		NotAcceptable                StatusCode
		ProxyAuthRequired            StatusCode
		RequestTimeout               StatusCode
		Conflict                     StatusCode
		Gone                         StatusCode
		LengthRequired               StatusCode
		PreconditionFailed           StatusCode
		RequestEntityTooLarge        StatusCode
		RequestURITooLong            StatusCode
		UnsupportedMediaType         StatusCode
		RequestedRangeNotSatisfiable StatusCode
		ExpectationFailed            StatusCode
		Teapot                       StatusCode
		MisdirectedRequest           StatusCode
		UnprocessableEntity          StatusCode
		Locked                       StatusCode
		FailedDependency             StatusCode
		UpgradeRequired              StatusCode
		PreconditionRequired         StatusCode
		TooManyRequests              StatusCode
		RequestHeaderFieldsTooLarge  StatusCode
		UnavailableForLegalReasons   StatusCode
		ClientClosedRequest          StatusCode
	}{
		Default:                      BadRequest,
		BadRequest:                   BadRequest,
		Unauthorized:                 Unauthorized,
		PaymentRequired:              PaymentRequired,
		Forbidden:                    Forbidden,
		NotFound:                     NotFound,
		MethodNotAllowed:             MethodNotAllowed,
		NotAcceptable:                NotAcceptable,
		ProxyAuthRequired:            ProxyAuthRequired,
		RequestTimeout:               RequestTimeout,
		Conflict:                     Conflict,
		Gone:                         Gone,
		LengthRequired:               LengthRequired,
		PreconditionFailed:           PreconditionFailed,
		RequestEntityTooLarge:        RequestEntityTooLarge,
		RequestURITooLong:            RequestURITooLong,
		UnsupportedMediaType:         UnsupportedMediaType,
		RequestedRangeNotSatisfiable: RequestedRangeNotSatisfiable,
		ExpectationFailed:            ExpectationFailed,
		Teapot:                       Teapot,
		MisdirectedRequest:           MisdirectedRequest,
		UnprocessableEntity:          UnprocessableEntity,
		Locked:                       Locked,
		FailedDependency:             FailedDependency,
		UpgradeRequired:              UpgradeRequired,
		PreconditionRequired:         PreconditionRequired,
		TooManyRequests:              TooManyRequests,
		RequestHeaderFieldsTooLarge:  RequestHeaderFieldsTooLarge,
		UnavailableForLegalReasons:   UnavailableForLegalReasons,
		ClientClosedRequest:          ClientClosedRequest,
	}

	ServerError = struct {
		Default                       StatusCode
		InternalServerError           StatusCode
		NotImplemented                StatusCode
		BadGateway                    StatusCode
		ServiceUnavailable            StatusCode
		GatewayTimeout                StatusCode
		HttpVersionNotSupported       StatusCode
		VariantAlsoNegotiates         StatusCode
		InsufficientStorage           StatusCode
		LoopDetected                  StatusCode
		NotExtended                   StatusCode
		NetworkAuthenticationRequired StatusCode
		NetworkConnectTimeout         StatusCode
	}{
		Default:                       InternalServerError,
		InternalServerError:           InternalServerError,
		NotImplemented:                NotImplemented,
		BadGateway:                    BadGateway,
		ServiceUnavailable:            ServiceUnavailable,
		GatewayTimeout:                GatewayTimeout,
		HttpVersionNotSupported:       HttpVersionNotSupported,
		VariantAlsoNegotiates:         VariantAlsoNegotiates,
		InsufficientStorage:           InsufficientStorage,
		LoopDetected:                  LoopDetected,
		NotExtended:                   NotExtended,
		NetworkAuthenticationRequired: NetworkAuthenticationRequired,
		NetworkConnectTimeout:         NetworkConnectTimeout,
	}
)
