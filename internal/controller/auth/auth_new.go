// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"backend/api/auth"
)

type ControllerV1 struct{}

type ControllerV1Protected struct{}

func NewV1() auth.IAuthV1 {
	return &ControllerV1{}
}

func NewV1Protected() auth.IAuthV1Protected {
	return &ControllerV1Protected{}
}
