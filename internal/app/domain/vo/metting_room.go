package vo

type MeetingRoomName struct {
	value string
}

func NewMeetingRoomName(value string) (MeetingRoomName, error) {
	if value == "" {
		return MeetingRoomName{}, ErrInvalidMeetingRoomName
	}
	return MeetingRoomName{value: value}, nil
}
