package dto

type Payload struct {
	Status StatusPayload `json:"status"`
	Data   interface{}   `json:"data"`
	Other  interface{}   `json:"other"`
}

type StatusPayload struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}
