package domain

import "time"

type Broker struct {
	Id       string    `json:"id"`
	Created  time.Time `json:"created"`
	Host     string    `json:"host"`
	Port     string    `json:"port,omitempty"`
	Username string    `json:"username"`
	Password string    `json:"-"`
}

func NewBroker(id string, created time.Time, host string, port string, username string, password string) Broker {
	return Broker{
		Id:       id,
		Created:  created,
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}
