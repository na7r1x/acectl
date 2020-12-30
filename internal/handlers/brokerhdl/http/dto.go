package brokerhdl

import "github.com/na7r1x/acectl/internal/core/domain"

// -------- Register ---------
type HttpBodyRegister struct {
	Id       string `json:"id" binding:"required"`
	Host     string `json:"host" binding:"required"`
	Port     string `json:"port"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// -------- List ---------
type HttpResponseList struct {
	Brokers []domain.Broker `json:"brokers"`
}

func BuildHttpResponseList(brokers []domain.Broker) HttpResponseList {
	return HttpResponseList{Brokers: brokers}
}

// -------- Status ---------
type HttpResponseStatus struct {
	Status         int    `json:"status"`
	StatusResponse string `json:"statusResponse"`
	ExecutionState string `json:"execState"`
}

func BuildHttpResponseStatus(status int, statusResponse string, exitState string) HttpResponseStatus {
	return HttpResponseStatus{
		Status:         status,
		StatusResponse: statusResponse,
		ExecutionState: exitState,
	}
}

// -------- Command ---------

type HttpResponseCommand struct {
	Success        bool   `json:"success"`
	Response       string `json:"response"`
	ExecutionState string `json:"execState"`
}

func BuildHttpResponseCommand(success bool, response string, exitState string) HttpResponseCommand {
	return HttpResponseCommand{
		Success:        success,
		Response:       response,
		ExecutionState: exitState,
	}
}
