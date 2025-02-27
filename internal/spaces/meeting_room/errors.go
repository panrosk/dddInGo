package meetingroom

import "errors"

var ErrInvalidMeetingRoomName = errors.New("el nombre de la sala de reuniones no puede estar vacío")
var ErrMeetingRoomAlreadyExists = errors.New("la sala de reuniones ya existe")
var ErrInvalidMeetingRoomCapacity = errors.New("la capacidad de la sala de reuniones debe ser mayor a 0")
var ErrInvalidHours = errors.New("las horas deben estar en el rango de 0 a 23")
var ErrInvalidDuration = errors.New("la duración debe ser mayor a 0 y menor de 12")
var ErrInvalidDate = errors.New("error de input en la fecha, esta mal formateada")
var ErrInvalidMeetingRoomId = errors.New("el id de la sala de reuniones no puede estar vacío")
var ErrInvdalidMeetingRoomUUID = errors.New("el id de la sala de reuniones no es un UUID válido")
var ErrMeetingRoomNotFound = errors.New("sala de reuniones no encontrada")
