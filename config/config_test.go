package config

import (
	"simple_chat_server/internal/model"
	"testing"
)

func TestConfig_ReadConfig(t *testing.T) {

	cfg := &Config{}

	type args struct {
		environment string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Testcase 1",
			args: args{
				environment: "test",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := cfg.ReadConfig(model.ConfigPath[tt.args.environment]); got != tt.want {
				t.Errorf("Config.ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
