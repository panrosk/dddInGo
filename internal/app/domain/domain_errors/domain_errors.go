package domain_errors

import "errors"

var ErrInvalidHotDeskNumber = errors.New("el número del hotdesk debe ser mayor a 0")
var ErrHotDeskAlreadyExists = errors.New("el hotdesk ya existe")
var ErrInvalidMeetingRoomName = errors.New("el nombre de la sala de reuniones no puede estar vacío")
var ErrMeetingRoomAlreadyExists = errors.New("la sala de reuniones ya existe")
var ErrInvalidMeetingRoomCapacity = errors.New("la capacidad de la sala de reuniones debe ser mayor a 0")
