package handler

import (
	"fmt"
)

func ErrorParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create opening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

// Validating parameters
func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return ErrorParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return ErrorParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return ErrorParamIsRequired("location", "string")
	}
	if r.Remote == nil {
		return ErrorParamIsRequired("remote", "bool")
	}
	if r.Link == "" {
		return ErrorParamIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return ErrorParamIsRequired("salary", "int64")
	}
	return nil
}

// Update opening

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provied, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	// If none of the fields were proviided, return false
	return fmt.Errorf("as least one valid field must be provided")
}
