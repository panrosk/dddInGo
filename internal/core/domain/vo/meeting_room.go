package vo

import "coworking/internal/app/domain/domain_errors"

type MeetingRoomName struct {
	value string
}

func NewMeetingRoomName(value string) (MeetingRoomName, error) {
	if value == "" {
		return MeetingRoomName{}, domain_errors.ErrInvalidMeetingRoomName
	}
	return MeetingRoomName{value: value}, nil
}

func (m MeetingRoomName) Value() string {
	return m.value
}

type MeetingRoomCapacity struct {
	value int
}

func NewMeetingRoomCapacity(value int) (MeetingRoomCapacity, error) {
	if value <= 0 {
		return MeetingRoomCapacity{}, domain_errors.ErrInvalidMeetingRoomCapacity
	}
	return MeetingRoomCapacity{value: value}, nil
}

func (m MeetingRoomCapacity) Value() int {
	return m.value
}
