package brokerrepo

import (
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/na7r1x/acectl/internal/core/domain"
)

func TestSQLiteStore_GetAll(t *testing.T) {
	broker := domain.Broker{
		Id:       "test",
		Created:  time.Now(),
		Host:     "test",
		Port:     "",
		Username: "unity",
		Password: "responsiv",
	}
	tests := []struct {
		name string
		db   string
		data []domain.Broker
		want []domain.Broker
	}{
		{"Get all when empty", "./testing.db", []domain.Broker{}, []domain.Broker{}},
		{"Get all when one record", "./testing.db", []domain.Broker{broker}, []domain.Broker{broker}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSqliteRepo(tt.db)
			s.Init()

			for _, rin := range tt.data {
				s.Set(rin)
			}
			if got, _ := s.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLiteStore.GetAll() = %v, want %v", got, tt.want)
			}
			s.Destroy()
		})
	}
}
