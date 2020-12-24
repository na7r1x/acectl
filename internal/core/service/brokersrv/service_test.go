package brokersrv

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/na7r1x/acectl/internal/core/domain"
	"github.com/na7r1x/acectl/mocks"
)

func Test_service_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	broker := domain.Broker{
		Id:       "BRK.IVO",
		Created:  time.Time{},
		Host:     "10.0.2.3",
		Port:     "",
		Username: "unity",
		Password: "responsiv",
	}

	mockRepo := mocks.NewMockBrokerRepository(ctrl)

	mockRepo.EXPECT().Set(gomock.Any()).Return(nil)
	mockRepo.EXPECT().Set(gomock.Any()).Return(errors.New("failed to insert broker"))

	type args struct {
		brk domain.Broker
	}
	tests := []struct {
		name    string
		srv     *service
		args    args
		wantErr bool
	}{
		{"register a broker", New(mockRepo), args{brk: broker}, false},
		{"fail register a broker", New(mockRepo), args{brk: broker}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.srv.Register(tt.args.brk); (err != nil) != tt.wantErr {
				t.Errorf("service.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Unregister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBrokerRepository(ctrl)

	mockRepo.EXPECT().Delete(gomock.Any()).Return(nil)
	mockRepo.EXPECT().Delete(gomock.Any()).Return(errors.New("Broker does not exist"))

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		srv     *service
		args    args
		wantErr bool
	}{
		{"unregister a broker", New(mockRepo), args{id: "BRK.IVO"}, false},
		{"unregister a non-existing broker", New(mockRepo), args{id: "BRK.IVO"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.srv.Unregister(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("service.Unregister() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func Test_service_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	broker := domain.Broker{
		Id:       "BRK.IVO",
		Created:  time.Time{},
		Host:     "10.0.2.3",
		Port:     "",
		Username: "unity",
		Password: "responsiv",
	}

	mockRepo := mocks.NewMockBrokerRepository(ctrl)

	mockRepo.EXPECT().GetAll().Return([]domain.Broker{broker}, nil)
	mockRepo.EXPECT().GetAll().Return(nil, errors.New("failed to fetch brokers"))

	tests := []struct {
		name    string
		srv     *service
		want    []domain.Broker
		wantErr bool
	}{
		{"list all brokers", New(mockRepo), []domain.Broker{broker}, false},
		{"fail listing all brokers", New(mockRepo), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	broker := domain.Broker{
		Id:       "BRK.IVO",
		Created:  time.Time{},
		Host:     "10.0.2.3",
		Port:     "",
		Username: "unity",
		Password: "responsiv",
	}

	mockRepo := mocks.NewMockBrokerRepository(ctrl)

	mockRepo.EXPECT().Get(gomock.Eq("BRK.IVO")).Return(broker, nil)
	mockRepo.EXPECT().Get(gomock.Any()).Return(domain.Broker{}, errors.New("failed to fetch broker"))

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		srv     *service
		args    args
		want    domain.Broker
		wantErr bool
	}{
		{"fetch broker", New(mockRepo), args{id: "BRK.IVO"}, broker, false},
		{"fetch non-existing broker", New(mockRepo), args{id: "NON_EXISTING"}, domain.Broker{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
