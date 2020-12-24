package execsrv

import (
	"strings"
	"testing"
	"time"

	"github.com/na7r1x/acectl/internal/core/domain"
)

func Test_service_Exec(t *testing.T) {
	testBroker := domain.NewBroker("test", time.Now(), "172.16.170.101", "", "unity", "responsiv")
	type args struct {
		brk domain.Broker
		cmd []string
	}
	tests := []struct {
		name    string
		srv     *service
		args    args
		want    string
		wantErr bool
	}{
		{
			"test running mqsilist",
			New(),
			args{brk: testBroker, cmd: []string{". .bash_profile && mqsilist"}},
			"BIP8071I: Successful command completion.",
			false,
		},
		{
			"test running mqsilist, no env",
			New(),
			args{brk: testBroker, cmd: []string{"mqsilist"}},
			"",
			true,
		},
		{
			"test running mqsilist, command warning",
			New(),
			args{brk: testBroker, cmd: []string{". .bash_profile && mqsilist BRK.TEST"}},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.Exec(tt.args.brk, tt.args.cmd, 60)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("service.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
