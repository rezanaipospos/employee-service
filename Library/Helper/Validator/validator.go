package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func Validate(mystruct interface{}) string {
	validate = validator.New()
	err := validate.Struct(mystruct)
	if err != nil {
		var errorValidate = "error when validation, error:"
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errorValidate += fmt.Sprintf("%s ,", err.Error())
		}
		errValidator := err.(validator.ValidationErrors)
		for index, err := range errValidator {
			if index+1 == len(errValidator) {
				errorValidate += err.Field() + " " + err.Tag()
			} else {
				errorValidate += err.Field() + " " + err.Tag() + ","
			}
		}
		return errorValidate
	}
	return ""
}

func ValidateSortColumn(allowedFields map[string]string, sortColumn string, defaultSortColumn string) string {
	column := allowedFields[sortColumn]
	if column == "" {
		column = defaultSortColumn
	}
	return column
}

func ValidateSortOrder(SortOrder string, defaultSortOrder string) string {
	var orders = []string{"ASC", "DESC"}
	for _, ordering := range orders {
		if strings.EqualFold(SortOrder, ordering) {
			return SortOrder
		}
	}
	return defaultSortOrder
}

func ValidateExtension(fileName string) error {
	var allowedExtension = []string{".jpeg", ".png", ".jpg"}
	for _, ext := range allowedExtension {
		if strings.Contains(ext, fileName) {
			return nil
		}
	}
	return fmt.Errorf("the provided file format is not allowed. please upload a %s file", strings.Join(allowedExtension, ", "))
}
