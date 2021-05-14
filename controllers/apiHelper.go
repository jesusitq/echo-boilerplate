package controllers

import (
	utils "api-echo-template/libs"
	"api-echo-template/models"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-validator/validator.v2"
)

// BodyToStruct :
func BodyToStruct(c echo.Context) (models.BoostrapSample, error) {
	// Read body
	boostrapSample := new(models.BoostrapSample)
	err := c.Bind(boostrapSample)
	return *boostrapSample, err
}

// ValidateInputModel :
func ValidateInputModel(boostrapSample models.BoostrapSample) (errToOut models.ErrorsResponse) {

	errValidator := validator.NewValidator().Validate(boostrapSample)
	if errValidator != nil {
		errToOut.Errors = utils.AppendErrors(errValidator)
	}
	return errToOut
}
