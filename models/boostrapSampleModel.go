package models

// BoostrapSample :
type BoostrapSample struct {
	ID          int    `json:"id" example:"123547" `
	Description string `json:"description" example:"Set a description for a value" validate:"nonzero,max=50"`
}
