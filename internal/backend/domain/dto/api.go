package dto

type ResponseAPI struct {
	Status  *ApiStatus  `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
