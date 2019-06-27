package utils

import (
	"testing"
)

func TestSimpePost(t *testing.T) {
	urlEntry := "http://v3810.cc.mobimon.com.tw/uzu/entry"
	url := "http://v3810.cc.mobimon.com.tw/uzu/result"
	type args struct {
		url    string
		fields map[string]string
		sid    string
	}

	myFieldEntry := map[string]string{
		"fid":   "1965350",
		"htype": "0",
		"uzid":  "7",
		"scid":  "42",
		"pt":    "0",
		"st":    "10",
	}

	myField := map[string]string{
		"res":     "1",
		"uzid":    "7",
		"wvt":     "[]",
		"mission": `{"cid":[7585,6203,6202,6201,6003,6272],"sid":[4010,248,5227,9214,8877,7206],"fid":[6202],"hrid":[],"ms":0,"md":0,"sc":{"0":0,"1":0,"2":0,"3":0,"4":0},"es":0,"at":0,"he":0,"da":0,"ba":0,"bu":0,"job":{"0":5,"1":0,"2":2,"3":0,"4":0},"weapon":{"0":5,"1":0,"2":0,"3":0,"4":0,"5":0,"8":2,"9":0,"10":0},"box":0,"um":{"1":0,"2":0,"3":0},"fj":-1,"fw":-1,"fo":0,"mlv":80,"mbl":436,"udj":0,"sdmg":0,"tp":0,"gma":4,"gmr":0,"gmp":0,"stp":1,"auto":0,"uh":{"10":7},"cc":1,"bf_atk":0,"bf_hp":0,"bf_spd":0}`,
	}
	tests := []struct {
		name        string
		args        args
		wantRespMap map[string]interface{}
		wantErr     bool
	}{
		{name: "test01", args: args{url: urlEntry, fields: myFieldEntry, sid: "157af3ac3707eeaac9795fee0b9cea73"}, wantRespMap: nil, wantErr: false},
		{name: "test02", args: args{url: url, fields: myField, sid: "157af3ac3707eeaac9795fee0b9cea73"}, wantRespMap: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, _ := SimpePost(tt.args.url, tt.args.fields, tt.args.sid)
			t.Log(resp)
		})
	}
}

// func TestPostV2(t *testing.T) {
// 	type args struct {
// 		requestUrl string
// 		rawPayload string
// 		body       map[string]interface{}
// 		sid        string
// 	}
// 	tests := []struct {
// 		name        string
// 		args        args
// 		wantRespMap map[string]interface{}
// 		wantErr     bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotRespMap, err := PostV2(tt.args.requestUrl, tt.args.rawPayload, tt.args.body, tt.args.sid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("PostV2() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotRespMap, tt.wantRespMap) {
// 				t.Errorf("PostV2() = %v, want %v", gotRespMap, tt.wantRespMap)
// 			}
// 		})
// 	}
// }

// func TestStructToMap(t *testing.T) {
// 	type args struct {
// 		i interface{}
// 	}
// 	tests := []struct {
// 		name       string
// 		args       args
// 		wantValues url.Values
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotValues := StructToMap(tt.args.i); !reflect.DeepEqual(gotValues, tt.wantValues) {
// 				t.Errorf("StructToMap() = %v, want %v", gotValues, tt.wantValues)
// 			}
// 		})
// 	}
// }

// func TestDecodeResponse(t *testing.T) {
// 	type args struct {
// 		raw []byte
// 	}
// 	tests := []struct {
// 		name       string
// 		args       args
// 		wantResult map[string]interface{}
// 		wantErr    bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotResult, err := DecodeResponse(tt.args.raw)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("DecodeResponse() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotResult, tt.wantResult) {
// 				t.Errorf("DecodeResponse() = %v, want %v", gotResult, tt.wantResult)
// 			}
// 		})
// 	}
// }

// func TestDecodeResponseV2(t *testing.T) {
// 	type args struct {
// 		raw []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantB   []byte
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotB, err := DecodeResponseV2(tt.args.raw)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("DecodeResponseV2() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotB, tt.wantB) {
// 				t.Errorf("DecodeResponseV2() = %v, want %v", gotB, tt.wantB)
// 			}
// 		})
// 	}
// }

// func TestStruct2JsonString(t *testing.T) {
// 	type args struct {
// 		s interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Struct2JsonString(tt.args.s); got != tt.want {
// 				t.Errorf("Struct2JsonString() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestMap2JsonString(t *testing.T) {
// 	type args struct {
// 		m map[string]interface{}
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantRet string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotRet := Map2JsonString(tt.args.m); gotRet != tt.wantRet {
// 				t.Errorf("Map2JsonString() = %v, want %v", gotRet, tt.wantRet)
// 			}
// 		})
// 	}
// }

// func TestMap2Struct(t *testing.T) {
// 	type args struct {
// 		m map[string]interface{}
// 		v interface{}
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := Map2Struct(tt.args.m, tt.args.v); (err != nil) != tt.wantErr {
// 				t.Errorf("Map2Struct() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestGetLogger(t *testing.T) {
// 	type args struct {
// 		f *os.File
// 	}
// 	tests := []struct {
// 		name       string
// 		args       args
// 		wantLogger *log.Logger
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotLogger := GetLogger(tt.args.f); !reflect.DeepEqual(gotLogger, tt.wantLogger) {
// 				t.Errorf("GetLogger() = %v, want %v", gotLogger, tt.wantLogger)
// 			}
// 		})
// 	}
// }

// func TestGetLoggerEx(t *testing.T) {
// 	type args struct {
// 		f *os.File
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantLog *logging.Logger
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotLog := GetLoggerEx(tt.args.f); !reflect.DeepEqual(gotLog, tt.wantLog) {
// 				t.Errorf("GetLoggerEx() = %v, want %v", gotLog, tt.wantLog)
// 			}
// 		})
// 	}
// }

// func TestDumpConfig(t *testing.T) {
// 	type args struct {
// 		c *config.Config
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			DumpConfig(tt.args.c)
// 		})
// 	}
// }

// func TestParseConfig2Struct(t *testing.T) {
// 	type args struct {
// 		conf    *config.Config
// 		section string
// 		data    interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ParseConfig2Struct(tt.args.conf, tt.args.section, tt.args.data)
// 		})
// 	}
// }

// func TestInArray(t *testing.T) {
// 	type args struct {
// 		val   interface{}
// 		array interface{}
// 	}
// 	tests := []struct {
// 		name       string
// 		args       args
// 		wantExists bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotExists := InArray(tt.args.val, tt.args.array); gotExists != tt.wantExists {
// 				t.Errorf("InArray() = %v, want %v", gotExists, tt.wantExists)
// 			}
// 		})
// 	}
// }
