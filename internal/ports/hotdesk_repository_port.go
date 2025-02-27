package ports

import "coworking/internal/spaces/hotdesk"

type HotDeskRepositoryPort interface {
	RepositoryPort[*hotdesk.Hotdesk]
	FindHotdeskByNumber(hd *hotdesk.Number) (*hotdesk.Hotdesk, error)
}
