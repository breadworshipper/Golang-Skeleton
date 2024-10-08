package response

import (
	"github.com/gofiber/fiber/v2"
)

// Typed structures for common responses
type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ErrorValidation struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type ErrorResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message,omitempty"`
	Errors  []ErrorValidation `json:"errors"`
}

type PaginatedResponse struct {
	Status   string      `json:"status"`
	Message  string      `json:"message,omitempty"`
	Data     any         `json:"data,omitempty"`
	Paginate *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type ResponseParams struct {
	StatusCode int
	Message    string
	Paginate   *Pagination
	Data       any
	Errors     []ErrorValidation
}

func SendResponse(ctx *fiber.Ctx, params ResponseParams) error {
	var response interface{}
	var status string

	if params.StatusCode >= 200 && params.StatusCode <= 299 {
		status = "success"
	} else {
		status = "error"
	}

	if params.Data != nil && params.Paginate == nil {
		response = &SuccessResponse{
			Status:  status,
			Message: params.Message,
			Data:    params.Data,
		}
	} else if params.Data != nil && params.Paginate != nil {
		response = &PaginatedResponse{
			Status:   status,
			Message:  params.Message,
			Data:     params.Data,
			Paginate: params.Paginate,
		}
	} else if len(params.Errors) > 0 {
		response = &ErrorResponse{
			Status:  status,
			Message: params.Message,
			Errors:  params.Errors,
		}
	} else {
		response = map[string]interface{}{
			"status":  status,
			"message": params.Message,
		}
	}

	return ctx.Status(params.StatusCode).JSON(response)
}
