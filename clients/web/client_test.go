package web

import "testing"

func TestGetGachaInfo(t *testing.T) {
	type args struct {
		t   string
		sid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test", wantErr: false, args: args{t: "event", sid: "6884b67af0317ffc805aa74bce8bf2de"}},
		{name: "test", wantErr: false, args: args{t: "story", sid: "6884b67af0317ffc805aa74bce8bf2de"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gachasInfo, err := GetGachaInfo(tt.args.t, tt.args.sid)
			if err != nil {
				t.Error(err)
			} else {
				t.Logf("%+v", gachasInfo)
			}
		})
	}
}
