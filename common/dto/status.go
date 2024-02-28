package dto

import (
	"fmt"
)

const (
	StatusSuccess    string = "Success"
	StatusFailed     string = "Failed"
	StatusBadRequest string = "Bad Request"
	StatusNotFound   string = "Not Found"
)

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

func (response *DefaultResponse) PrepareStatusSuccess(message string) {
	response.Status = StatusSuccess
	response.Message = message
}

func (response *DefaultResponse) PrepareStatusFailed(message string) {
	response.Status = StatusFailed
	response.Message = message
}