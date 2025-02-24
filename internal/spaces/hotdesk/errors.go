package hotdesk

import "errors"

var ErrInvalidHotDeskNumber = errors.New("el número del hotdesk debe ser mayor a 0")
var ErrHotDeskAlreadyExists = errors.New("el hotdesk ya existe")
var ErrHotDeskAlredyReserved = errors.New("el hotdesk ya está reservado for that date for that user")
