package dto

import "net/http"

type ApiStatus struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (s *ApiStatus) Error() bool {
	panic("ApiStatus does not implement error interface, use WithReason to set message")
}

func (s *ApiStatus) WithReason(reason string) *ApiStatus {
	return &ApiStatus{
		Code:       s.Code,
		Message:    reason,
		StatusCode: s.StatusCode,
	}
}

var (
	ApiStatusOK                  = &ApiStatus{Code: 200, Message: "OK", StatusCode: http.StatusOK}
	ApiStatusCreated             = &ApiStatus{Code: 201, Message: "CREATED", StatusCode: http.StatusCreated}
	ApiStatusNotFound            = &ApiStatus{Code: 404, Message: "Not Found", StatusCode: http.StatusNotFound}
	ApiStatusError               = &ApiStatus{Code: 500, Message: "Internal Server Error", StatusCode: http.StatusInternalServerError}
	ApiStatusBadRequest          = &ApiStatus{Code: 400, Message: "Bad Request", StatusCode: http.StatusBadRequest}
	ApiStatusUnauthorized        = &ApiStatus{Code: 401, Message: "Unauthorized", StatusCode: http.StatusUnauthorized}
	ApiStatusForbidden           = &ApiStatus{Code: 403, Message: "Forbidden", StatusCode: http.StatusForbidden}
	ApiStatusConflict            = &ApiStatus{Code: 409, Message: "Conflict", StatusCode: http.StatusConflict}
	ApiStatusUnprocessableEntity = &ApiStatus{Code: 422, Message: "Unprocessable Entity", StatusCode: http.StatusUnprocessableEntity}
	ApiStatusServiceUnavailable  = &ApiStatus{Code: 503, Message: "Service Unavailable", StatusCode: http.StatusServiceUnavailable}
	ApiExternalError             = &ApiStatus{Code: 501, Message: "External Service Error", StatusCode: http.StatusInternalServerError}
)
