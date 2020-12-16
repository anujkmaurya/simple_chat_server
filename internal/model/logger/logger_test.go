package logger

import "testing"

func TestInitLogger(t *testing.T) {
	type args struct {
		logfilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Testcase 1",
			args: args{
				logfilePath: "../../mock/files/testfile",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitLogger(tt.args.logfilePath); (err != nil) != tt.wantErr {
				t.Errorf("InitLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
