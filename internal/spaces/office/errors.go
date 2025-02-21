package office

import "errors"

var ErrInvalidOfficeNumber = errors.New("el número de la oficina debe ser mayor a 0")
var ErrInvalidOfficeLeasePeriod = errors.New("el período de alquiler de la oficina debe ser mayor a 0")
var ErrOfficeAlreadyExists = errors.New("la oficina ya existe")
