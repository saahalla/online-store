package common

import (
	"fmt"
	"time"
)

const (
	StatusSuccess    string = "Success"
	StatusFailed     string = "Failed"
	StatusBadRequest string = "Bad Request"
	StatusNotFound   string = "Not Found"
)

type DefaultDate struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DefaultResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (response *DefaultResponse) PrepareStatusBadRequestRequired(requiredField string) {
	response.Status = StatusBadRequest
	response.Message = fmt.Sprintf("%v is required", requiredField)
}

func (response *DefaultResponse) PrepareStatusBadRequestInvalid(requiredField string) {
	response.Status = StatusBadRequest
	response.Message = fmt.Sprintf("Invalid %v", requiredField)
}

func (response *DefaultResponse) PrepareStatusNotFound(requiredField string) {
	response.Status = StatusNotFound
	response.Message = fmt.Sprintf("%v is not found", requiredField)
}

func (response *DefaultResponse) PrepareStatusSuccess() {
	response.Status = StatusSuccess
	response.Message = StatusSuccess
}
