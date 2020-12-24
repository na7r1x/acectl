package brokersrv

import (
	"errors"

	"github.com/na7r1x/acectl/internal/core/domain"
	"github.com/na7r1x/acectl/internal/core/ports"
	"github.com/na7r1x/acectl/internal/core/service/execsrv"
)

type service struct {
	brokerRepository ports.BrokerRepository
	executorService  ports.ExecutorService
}

func New(brokerRepository ports.BrokerRepository) *service {
	return &service{
		brokerRepository: brokerRepository,
		executorService:  execsrv.New(),
	}
}

func (srv *service) Register(brk domain.Broker) error {
	return srv.brokerRepository.Set(brk)
}

func (srv *service) Unregister(id string) error {
	return srv.brokerRepository.Delete(id)
}

func (srv *service) List() ([]domain.Broker, error) {
	brokers, err := srv.brokerRepository.GetAll()
	if err != nil {
		return nil, errors.New("could not retrieve brokers list; " + err.Error())
	}

	return brokers, nil
}

func (srv *service) Get(id string) (domain.Broker, error) {
	broker, err := srv.brokerRepository.Get(id)
	if err != nil {
		return domain.Broker{}, errors.New("could not fetch broker [" + id + "]; " + err.Error())
	}

	return broker, nil
}

func (srv *service) Status(id string, timeout int) (string, string, error) {
	broker, err := srv.brokerRepository.Get(id)
	if err != nil {
		return "", "", errors.New("could not get broker; " + err.Error())
	}

	response, exitState, err := srv.executorService.Exec(broker, []string{". .bash_profile && mqsilist | grep " + broker.Id}, 60)
	if err != nil {
		return "", exitState, errors.New("could not execute command; " + err.Error())
	}

	if response == "" {
		return "", exitState, errors.New("broker not found;")
	}

	return response, exitState, nil
}

func (srv *service) Start(id string, timeout int) (string, string, error) {
	broker, err := srv.brokerRepository.Get(id)
	if err != nil {
		return "", "", errors.New("could not get broker; " + err.Error())
	}

	response, exitState, err := srv.executorService.Exec(broker, []string{". .bash_profile && mqsistart " + broker.Id}, 60)
	if err != nil {
		return "", exitState, errors.New("could not execute command; " + err.Error())
	}

	return response, exitState, nil
}

func (srv *service) Stop(id string, timeout int) (string, string, error) {
	broker, err := srv.brokerRepository.Get(id)
	if err != nil {
		return "", "", errors.New("could not get broker; " + err.Error())
	}

	response, exitState, err := srv.executorService.Exec(broker, []string{". .bash_profile && mqsistop " + broker.Id}, 60)
	if err != nil {
		return "", exitState, errors.New("could not execute command; " + err.Error())
	}

	return response, exitState, nil
}

func (srv *service) ForceStop(id string, timeout int) (string, string, error) {
	broker, err := srv.brokerRepository.Get(id)
	if err != nil {
		return "", "", errors.New("could not get broker; " + err.Error())
	}

	response, exitState, err := srv.executorService.Exec(broker, []string{". .bash_profile && mqsistop -i " + broker.Id}, 60)
	if err != nil {
		return "", exitState, errors.New("could not execute command; " + err.Error())
	}

	return response, exitState, nil
}
