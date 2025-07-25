package handler

import (
	"booking-system/internal/backend/application/dto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerFunc func(*gin.Context) error

type ResponseHandler struct {
	logger *log.Logger
}

func NewResponseHandler() *ResponseHandler {
	return &ResponseHandler{
		logger: log.Default(),
	}
}

func NewResponseHandlerWithLogger(logger *log.Logger) *ResponseHandler {
	return &ResponseHandler{
		logger: logger,
	}
}

// WrapController converts controller method to gin.HandlerFunc
func (h *ResponseHandler) WrapController(controllerFunc ControllerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := controllerFunc(c); err != nil {
			h.handleError(c, err)
		}
	}
}

// handleError processes different types of errors and returns appropriate responses
func (h *ResponseHandler) handleError(c *gin.Context, err error) {
	h.logger.Printf("Controller error: %v", err)
	if c.Writer.Written() {
		return
	}

	switch e := err.(type) {
	case *ValidationError:
		c.JSON(http.StatusBadRequest, dto.ResponseAPI{
			Status:  dto.ApiStatusBadRequest.WithReason(e.Message),
			Message: e.Message,
		})
	case *AuthenticationError:
		c.JSON(http.StatusUnauthorized, dto.ResponseAPI{
			Status:  dto.ApiStatusUnauthorized.WithReason(e.Message),
			Message: e.Message,
		})
	case *AuthorizationError:
		c.JSON(http.StatusForbidden, dto.ResponseAPI{
			Status:  dto.ApiStatusForbidden.WithReason(e.Message),
			Message: e.Message,
		})
	case *NotFoundError:
		c.JSON(http.StatusNotFound, dto.ResponseAPI{
			Status:  dto.ApiStatusNotFound.WithReason(e.Message),
			Message: e.Message,
		})
	case *ConflictError:
		c.JSON(http.StatusConflict, dto.ResponseAPI{
			Status:  dto.ApiStatusConflict.WithReason(e.Message),
			Message: e.Message,
		})
	default:
		c.JSON(http.StatusInternalServerError, dto.ResponseAPI{
			Status:  dto.ApiStatusError.WithReason("Internal server error"),
			Message: "An unexpected error occurred",
		})
	}
}

// func (h *ResponseHandler) SendSuccess(c *gin.Context, message string, data interface{}) {
// 	c.JSON(http.StatusOK, dto.ResponseAPI{
// 		Status:  dto.ApiStatusOK,
// 		Message: message,
// 		Data:    data,
// 	})
// }

// SendCreated sends a created response
// func (h *ResponseHandler) SendCreated(c *gin.Context, message string, data interface{}) {
// 	c.JSON(http.StatusCreated, dto.ResponseAPI{
// 		Status:  dto.ApiStatusOK,
// 		Message: message,
// 		Data:    data,
// 	})

// }

// SendNoContent sends a no content response
// func (h *ResponseHandler) SendNoContent(c *gin.Context) {
// 	c.Status(http.StatusNoContent)
// }
