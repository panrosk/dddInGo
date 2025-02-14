package errors

import "errors"

var ErrInvalidHotDeskNumber = errors.New("el número del hotdesk debe ser mayor a 0")
var ErrHotDeskAlreadyExists = errors.New("el hotdesk ya existe")
