package membership

import "errors"

var ErrUserIDIsRequired = errors.New("userID is required")
var ErrMembershipAlreadyExists = errors.New("membership already exists")
var ErrInvalidMonth = errors.New("invalid month")
var ErrInvalidCredits = errors.New("invalid credits")
var ErrMembershipNotFound = errors.New("membership not found")
