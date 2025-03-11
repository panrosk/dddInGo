package membership

type MembershipCreatedPayload struct {
	UserID string
}

func NewMembershipCreatedPayload(userID string) (MembershipCreatedPayload, error) {
	if userID == "" {
		return MembershipCreatedPayload{}, ErrUserIDIsRequired
	}

	return MembershipCreatedPayload{
		UserID: userID,
	}, nil

}
