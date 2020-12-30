package domain

import "time"

type Broker struct {
	Id       string    `json:"id" header:"id"`
	Created  time.Time `json:"created" header:"created"`
	Host     string    `json:"host" header:"host"`
	Port     string    `json:"port,omitempty" header:"port"`
	Username string    `json:"username" header:"username"`
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
