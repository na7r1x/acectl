package ports

//go:generate mockgen -destination=../../../mocks/mock_repositories.go -package=mocks github.com/na7r1x/acectl/internal/core/ports BrokerRepository

import "github.com/na7r1x/acectl/internal/core/domain"

type BrokerRepository interface {
	Get(string) (domain.Broker, error)
	Set(domain.Broker) error
	Delete(string) error
	GetAll() ([]domain.Broker, error)
}
