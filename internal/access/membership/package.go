package membership

import "time"

type Package struct {
	packageID string
	credits   int
	startedAt time.Time
	endedAt   time.Time
}
