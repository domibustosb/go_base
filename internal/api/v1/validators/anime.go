package validators

import "github.com/go-playground/validator"

type AnimeValidator struct {
	ID   *string `validate:"required,min=3,max=32"`
	Name *string `json:"name"`
}
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(animeValidator AnimeValidator) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(animeValidator)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
