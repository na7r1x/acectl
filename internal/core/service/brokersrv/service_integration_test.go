package brokersrv

import (
	"log"
	"strings"
	"testing"
	"time"

	"github.com/na7r1x/acectl/internal/core/domain"
	"github.com/na7r1x/acectl/internal/repositories/brokerrepo"
)

func Test_service_Stop(t *testing.T) {
	testDb := "./testing.db"
	brokerRepo := brokerrepo.NewSqliteRepo(testDb)
	broker := domain.Broker{
		Id:       "BRK.IVO",
		Created:  time.Time{},
		Host:     "10.0.2.3",
		Port:     "",
		Username: "unity",
		Password: "responsiv",
	}
	brokerRepo.Set(broker)
	defer func() {
		if err := brokerRepo.Destroy(); err != nil {
			log.Fatal(err)
		}
	}()

	type args struct {
		id      string
		timeout int
	}
	tests := []struct {
		name    string
		srv     *service
		args    args
		want    string
		wantErr bool
	}{
		{"stop a broker", New(brokerRepo), args{id: broker.Id, timeout: 60}, "Successful", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.Stop(tt.args.id, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Stop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.Contains(got, tt.want) {
				t.Errorf("service.Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_ForceStop(t *testing.T) {
	testDb := "./testing.db"
	brokerRepo := brokerrepo.NewSqliteRepo(testDb)
	broker := domain.Broker{
		Id:       "BRK.IVO",
		Created:  time.Time{},
		Host:     "10.0.2.3",
		Port:     "",
		Username: "unity",
		Password: "responsiv",
	}
	brokerRepo.Set(broker)
	defer func() {
		if err := brokerRepo.Destroy(); err != nil {
			log.Fatal(err)
		}
	}()

	type args struct {
		id      string
		timeout int
	}
	tests := []struct {
		name    string
		srv     *service
		args    args
		want    string
		wantErr bool
	}{
		{"force stop a broker", New(brokerRepo), args{id: broker.Id, timeout: 60}, "Successful", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.ForceStop(tt.args.id, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ForceStop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.Contains(got, tt.want) {
				t.Errorf("service.ForceStop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Start(t *testing.T) {
	testDb := "./testing.db"
	brokerRepo := brokerrepo.NewSqliteRepo(testDb)
	broker := domain.Broker{
		Id:       "BRK.IVO",
		Created:  time.Time{},
		Host:     "10.0.2.3",
		Port:     "",
		Username: "unity",
		Password: "responsiv",
	}
	brokerRepo.Set(broker)
	defer func() {
		if err := brokerRepo.Destroy(); err != nil {
			log.Fatal(err)
		}
	}()

	type args struct {
		id      string
		timeout int
	}
	tests := []struct {
		name    string
		srv     *service
		args    args
		want    string
		wantErr bool
	}{
		{"start a broker", New(brokerRepo), args{id: broker.Id, timeout: 60}, "Successful", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.Start(tt.args.id, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.Contains(got, tt.want) {
				t.Errorf("service.Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Status(t *testing.T) {
	testDb := "./testing.db"
	brokerRepo := brokerrepo.NewSqliteRepo(testDb)
	broker := domain.Broker{
		Id:       "BRK.IVO",
		Created:  time.Time{},
		Host:     "10.0.2.3",
		Port:     "",
		Username: "unity",
		Password: "responsiv",
	}
	brokerRepo.Set(broker)
	defer func() {
		if err := brokerRepo.Destroy(); err != nil {
			log.Fatal(err)
		}
	}()

	type args struct {
		id      string
		timeout int
	}
	tests := []struct {
		name    string
		srv     *service
		args    args
		want    string
		wantErr bool
	}{
		{"get status", New(brokerRepo), args{id: "BRK.IVO", timeout: 60}, "", false},
		{"get status (non-existing broker)", New(brokerRepo), args{id: "NON_EXISTING", timeout: 60}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.Status(tt.args.id, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Status() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("service.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
