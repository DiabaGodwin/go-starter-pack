package service

import "errors"

var ErrValidation = errors.New("validation error")
var NotFound = errors.New("not found")
var ErrInvalidCredentials = errors.New("invalid credentials")
