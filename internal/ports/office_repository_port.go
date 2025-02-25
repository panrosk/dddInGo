package ports

import "coworking/internal/spaces/office"

type OfficeRepositoryPort interface {
	RepositoryPort[*office.Office]
	FindByNumber(office *office.Office) (*office.Office, error)
}
