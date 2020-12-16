package config

import (
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		environment string
	}
	tests := []struct {
		name     string
		args     args
		wantPort string
	}{
		{
			name: "Testcase 1",
			args: args{
				environment: "test",
			},
			wantPort: "9000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(tt.args.environment); got != nil && got.Server.Port != tt.wantPort {
				t.Errorf("Init() = %v, want %v", got, tt.wantPort)
			}
		})
	}
}
