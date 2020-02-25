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
				school: School,
				token:  "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1ODI2ODI0NzQsImF1ZCI6IjIwMTgxMDAwMDQzMSJ9.yLIdapycLSamAVHGrm8AvZSnGS-rWR4Kjji-h2ZU_Mg",
				xh:     "201810000431",
				xnxqid: "2019-2020-2",
				zc:     "2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getKbcxAzc(tt.args.school, tt.args.token, tt.args.xh, tt.args.xnxqid, tt.args.zc)
			if (err != nil) != tt.wantErr {
				t.Errorf("getKbcxAzc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("getKbcxAzc() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
