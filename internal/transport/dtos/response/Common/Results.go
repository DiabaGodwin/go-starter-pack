package Common

import "net/http"

const (
	internalServerErrorResponseMessage        = "An error occurred, please try again later."
	BadRequestErrorResponseMessage            = "One or more validation errors occurred, please check and try again"
	DefaultOkResponseMessage                  = "Retrieved successfully"
	DefaultAcceptedResponseMessage            = "Accepted"
	defaultNotFoundResponseMessage            = "Resource not found"
	defaultCreatedResponseMessage             = "Created successfully"
	defaultUnprocessableEntityResponseMessage = "Unable to process the contained instructions"
	defaultForbiddenMessage                   = "Access Denied"
	defaultNoContentMessage                   = "No content found"
	defaultUnauthorizedMessage                = "Not authorized to perform this operation"
)

// ---------- Error Responses ----------

func InternalServerError[T any](message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, internalServerErrorResponseMessage),
		Code:    http.StatusInternalServerError,
	}
}

func NotFound[T any](message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, defaultNotFoundResponseMessage),
		Code:    http.StatusNotFound,
	}
}

func UnprocessableEntity[T any](message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, defaultUnprocessableEntityResponseMessage),
		Code:    http.StatusUnprocessableEntity,
	}
}

func BadRequest[T any](message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, BadRequestErrorResponseMessage),
		Code:    http.StatusBadRequest,
	}
}

func Conflict[T any](message string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: message,
		Code:    http.StatusConflict,
	}
}

func Forbidden[T any](message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, defaultForbiddenMessage),
		Code:    http.StatusForbidden,
	}
}

func Unauthorized[T any](message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, defaultUnauthorizedMessage),
		Code:    http.StatusUnauthorized,
	}
}

// ---------- Success Responses ----------

func Created[T any](data T, message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, defaultCreatedResponseMessage),
		Code:    http.StatusCreated,
		Data:    &data,
	}
}

func Ok[T any](data T, message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, DefaultOkResponseMessage),
		Code:    http.StatusOK,
		Data:    &data,
	}
}

func Accepted[T any](data T, message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, DefaultAcceptedResponseMessage),
		Code:    http.StatusAccepted,
		Data:    &data,
	}
}

func NoContent[T any](message ...string) ApiResponse[T] {
	return ApiResponse[T]{
		Message: firstOrDefault(message, defaultNoContentMessage),
		Code:    http.StatusNoContent,
	}
}

// ---------- Helper ----------

func firstOrDefault(values []string, fallback string) string {
	if len(values) > 0 && values[0] != "" {
		return values[0]
	}
	return fallback
}
