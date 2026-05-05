package Common

type ApiResponse[T any] struct {
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
	Code    int    `json:"code"`
	TraceID string `json:"traceId,omitempty"`
}
