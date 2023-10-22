package errors

import "net/http"

const (
	BadRequest          = "BAD_REQUEST"
	NotFound            = "NOT_FOUND"
	InternalServerError = "INTERNAL_SERVER_ERROR"
)

func NewBadRequest(err string) HTTPError {
	return NewHTTPError(
		http.StatusBadRequest,
		err,
		BadRequest,
	)
}

func NewNotFound(err string) HTTPError {
	return NewHTTPError(
		http.StatusNotFound,
		err,
		NotFound,
	)
}

func NewInternalServerError(err string) HTTPError {
	return NewHTTPError(
		http.StatusInternalServerError,
		err,
		InternalServerError,
	)
}
