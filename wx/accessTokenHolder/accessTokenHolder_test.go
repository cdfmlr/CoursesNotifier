package accessTokenHolder

import "testing"

func TestAccessTokenHolder_Get(t *testing.T) {
	type fields struct {
		accessToken string
		createTime  int64
		expiresIn   int64
	}
	type args struct {
		appID     string
		appSecret string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "test",
			fields: fields{
				accessToken: "",
				createTime:  0,
				expiresIn:   0,
			},
			args:   args{
				appID:     "wx63cf76ed67d69bb1",
				appSecret: "8a62c82aeac97ebf79b4617049499302",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &AccessTokenHolder{
				accessToken: tt.fields.accessToken,
				createTime:  tt.fields.createTime,
				expiresIn:   tt.fields.expiresIn,
			}
			if got := h.Get(tt.args.appID, tt.args.appSecret); got == "" {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			} else {
				t.Log(got)
			}
		})
	}
}