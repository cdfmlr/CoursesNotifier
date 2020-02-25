package qzapi

import (
	"testing"
)

func TestAuthUser(t *testing.T) {
	type args struct {
		sid string
		pwd string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "auth user me",
			args:args{
				sid: "201810000431",
				pwd: "hd270516",
			},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := authUser(School, tt.args.sid, tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthUser() error = %v, wantErr %v", err, tt.wantErr)
				//return
			}
			t.Log(got)
		})
	}
}