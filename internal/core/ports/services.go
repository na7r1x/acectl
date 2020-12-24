package ports

//go:generate mockgen -destination=../../../mocks/mock_services.go -package=mocks github.com/na7r1x/acectl/internal/core/ports BrokerService,ExecutorService

import "github.com/na7r1x/acectl/internal/core/domain"

type BrokerService interface {
	Register(domain.Broker) error
	Unregister(id string) error
	List() ([]domain.Broker, error)
	Get(id string) (domain.Broker, error)
	Status(id string, timeout int) (string, string, error)
	Start(id string, timeout int) (string, string, error)
	Stop(id string, timeout int) (string, string, error)
}

type ExecutorService interface {
	Exec(brk domain.Broker, cmd []string, timeout int) (string, string, error)
}
