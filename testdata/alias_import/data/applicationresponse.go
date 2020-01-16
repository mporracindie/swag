package data

import (
	typesapplication "github.com/mporracindie/swag/testdata/alias_import/types"
)

type ApplicationResponse struct {
	Application      typesapplication.Application   `json:"application"`
	ApplicationArray []typesapplication.Application `json:"application_array"`
	ApplicationTime  typesapplication.DateOnly      `json:"application_time"`
}
