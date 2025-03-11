package membership

import "time"

type PackageSubscribedPayload struct {
	PackageID string
	Credits   int
	StartedAt time.Time
	EndedAt   time.Time
}
