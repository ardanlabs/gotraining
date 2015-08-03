// Package models contains data structures and associated behavior
package models

import "gopkg.in/bluesuncorp/validator.v6"

var validate *validator.Validate

func init() {
	config := validator.Config{
		TagName:         "validate",
		ValidationFuncs: validator.BakedInValidators,
	}

	validate = validator.New(config)
}
