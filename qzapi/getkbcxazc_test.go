package qzapi

import (
	"testing"
)

func Test_getKbcxAzc(t *testing.T) {
	type args struct {
		school string
		token  string
		xh     string
		xnxqid string
		zc     string
	}
	tests := []struct {
		name    string
		args    args
		want    []GetKbcxAzcRespBodyItem
		wantErr bool
	}{
		{
			name: "get kbcx at zc",
			args: args{
				school: SchoolNcepu,
				token:  "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1ODI2ODcyNjcsImF1ZCI6IjIwMTgxMDAwMDQzMSJ9.vm2jxI9ilvrmKMHnXd2dkGB2ERhfea4mNspE1ZK1VM8",
				xh:     "201810000431",
				xnxqid: "2019-2020-2",
				zc:     "2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetKbcxAzc(tt.args.school, tt.args.token, tt.args.xh, tt.args.xnxqid, tt.args.zc)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKbcxAzc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetKbcxAzc() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
