package radarr

import "errors"

var ErrUnauthorized = errors.New("Unauthorized")
var ErrInvalidID = errors.New("Invalid ID supplied")
var ErrInvalidAPI = errors.New("Invalid API Key")
var ErrNotFound = errors.New("Movie not found")
var ErrValidation = errors.New("Validation exception")
