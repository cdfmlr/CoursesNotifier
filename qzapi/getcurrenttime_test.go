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
				school:   SchoolNcepu,
				token:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1ODI2ODcyNjcsImF1ZCI6IjIwMTgxMDAwMDQzMSJ9.vm2jxI9ilvrmKMHnXd2dkGB2ERhfea4mNspE1ZK1VM8",
				currDate: "2020-02-25",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCurrentTime(tt.args.school, tt.args.token, tt.args.currDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchCurrentTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("FetchCurrentTime() got = %v, want %v", got, tt.want)
			//}
		})
	}
}