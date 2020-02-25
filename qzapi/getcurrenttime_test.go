package qzapi

import (
	"testing"
)

func Test_getCurrentTime(t *testing.T) {
	type args struct {
		school   string
		token    string
		currDate string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetCurrentTimeRespBody
		wantErr bool
	}{
		{
			name: "get current time",
			args:args{
				school:   School,
				token:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1ODI2NzkxNjUsImF1ZCI6IjIwMTgxMDAwMDQzMSJ9.2pG_nvLXScOng_2W5jxETbz4I3EVVTUg6t_ruem30Xw",
				currDate: "2020-02-25",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCurrentTime(tt.args.school, tt.args.token, tt.args.currDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCurrentTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("getCurrentTime() got = %v, want %v", got, tt.want)
			//}
		})
	}
}