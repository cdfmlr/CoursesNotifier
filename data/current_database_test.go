package data

import (
	"reflect"
	"testing"
	"time"
)

func TestCurrentDatabase_GetCurrentTermBeginDate(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    time.Time
		wantErr bool
	}{
		{
			name:    "GetCurrentTermBeginDate",
			fields:  fields{"c:000123@/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true"},
			want:    time.Date(2020, 02, 17, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crtdb := &CurrentDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			got, err := crtdb.GetCurrentTermBeginDate()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentTermBeginDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Unix(), tt.want.Unix()) {
				t.Errorf("GetCurrentTermBeginDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCurrentDatabase(t *testing.T) {
	type args struct {
		dataSourceName string
	}
	tests := []struct {
		name    string
		args    args
		wantDSN string
	}{
		{
			name:    "1",
			args:    args{"c:000123@/test?charset=utf8"},
			wantDSN: "c:000123@/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true",
		},
		{
			name:    "2",
			args:    args{"c:000123@/test"},
			wantDSN: "c:000123@/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true",
		},
		{
			name:    "3",
			args:    args{"c:000123@/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true"},
			wantDSN: "c:000123@/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true&loc=Asia%2FShanghai&parseTime=true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCurrentDatabase(tt.args.dataSourceName); !reflect.DeepEqual(got.dataSourceName, tt.wantDSN) {
				t.Errorf("NewCurrentDatabase().dataSourceName = %v, want %v", got, tt.wantDSN)
			} else {
				t.Log(got.GetCurrentTermBeginDate())
			}
		})
	}
}
