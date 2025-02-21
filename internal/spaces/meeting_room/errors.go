package meetingroom

import "errors"

var ErrInvalidMeetingRoomName = errors.New("el nombre de la sala de reuniones no puede estar vacío")
var ErrMeetingRoomAlreadyExists = errors.New("la sala de reuniones ya existe")
var ErrInvalidMeetingRoomCapacity = errors.New("la capacidad de la sala de reuniones debe ser mayor a 0")
