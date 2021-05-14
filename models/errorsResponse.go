package models

// ErrorsResponse :
type ErrorsResponse struct {
	Errors []ErrorResponse `json:"errors"`
}

// ErrorResponse :
type ErrorResponse struct {
	Code    string `json:"code" example:"ERROR.QS.ID"`
	Message string `json:"message" example:"id must be positive integer"`
}
