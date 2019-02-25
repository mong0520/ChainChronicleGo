package quest

import (
	"reflect"
	"testing"

	"github.com/mong0520/ChainChronicleGo/clients"
)

func Test_quest_StartQeust(t *testing.T) {
	type fields struct {
		AutoSell        bool
		AutoBuy         bool
		AutoRaid        bool
		AutoRaidRecover bool
		Count           int
		Type            int
		QuestId         int
		QuestIds        string
		Fid             int
		Pt              int
		Htype           int
		Lv              int
		Hcid            int
		Version         int
		Res             int
		Bt              int
		Wc              int
		Wn              int
		Time            string
		D               int
		S               int
		Cc              int
		Oc              int
	}
	type args struct {
		u *clients.Metadata
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp map[string]interface{}
		wantRes  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quest{
				AutoSell:        tt.fields.AutoSell,
				AutoBuy:         tt.fields.AutoBuy,
				AutoRaid:        tt.fields.AutoRaid,
				AutoRaidRecover: tt.fields.AutoRaidRecover,
				Count:           tt.fields.Count,
				Type:            tt.fields.Type,
				QuestId:         tt.fields.QuestId,
				QuestIds:        tt.fields.QuestIds,
				Fid:             tt.fields.Fid,
				Pt:              tt.fields.Pt,
				Htype:           tt.fields.Htype,
				Lv:              tt.fields.Lv,
				Hcid:            tt.fields.Hcid,
				Version:         tt.fields.Version,
				Res:             tt.fields.Res,
				Bt:              tt.fields.Bt,
				Wc:              tt.fields.Wc,
				Wn:              tt.fields.Wn,
				Time:            tt.fields.Time,
				D:               tt.fields.D,
				S:               tt.fields.S,
				Cc:              tt.fields.Cc,
				Oc:              tt.fields.Oc,
			}
			gotResp, gotRes := q.StartQeust(tt.args.u)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("quest.StartQeust() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotRes != tt.wantRes {
				t.Errorf("quest.StartQeust() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_quest_EndQeustV2(t *testing.T) {
	q := NewQuest()
	q.EndQeustV2(nil)

}

func Test_quest_EndQeust(t *testing.T) {
	type fields struct {
		AutoSell        bool
		AutoBuy         bool
		AutoRaid        bool
		AutoRaidRecover bool
		Count           int
		Type            int
		QuestId         int
		QuestIds        string
		Fid             int
		Pt              int
		Htype           int
		Lv              int
		Hcid            int
		Version         int
		Res             int
		Bt              int
		Wc              int
		Wn              int
		Time            string
		D               int
		S               int
		Cc              int
		Oc              int
	}
	type args struct {
		u *clients.Metadata
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp map[string]interface{}
		wantRes  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quest{
				AutoSell:        tt.fields.AutoSell,
				AutoBuy:         tt.fields.AutoBuy,
				AutoRaid:        tt.fields.AutoRaid,
				AutoRaidRecover: tt.fields.AutoRaidRecover,
				Count:           tt.fields.Count,
				Type:            tt.fields.Type,
				QuestId:         tt.fields.QuestId,
				QuestIds:        tt.fields.QuestIds,
				Fid:             tt.fields.Fid,
				Pt:              tt.fields.Pt,
				Htype:           tt.fields.Htype,
				Lv:              tt.fields.Lv,
				Hcid:            tt.fields.Hcid,
				Version:         tt.fields.Version,
				Res:             tt.fields.Res,
				Bt:              tt.fields.Bt,
				Wc:              tt.fields.Wc,
				Wn:              tt.fields.Wn,
				Time:            tt.fields.Time,
				D:               tt.fields.D,
				S:               tt.fields.S,
				Cc:              tt.fields.Cc,
				Oc:              tt.fields.Oc,
			}
			gotResp, gotRes := q.EndQeust(tt.args.u)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("quest.EndQeust() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotRes != tt.wantRes {
				t.Errorf("quest.EndQeust() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_quest_GetTreasure(t *testing.T) {
	type fields struct {
		AutoSell        bool
		AutoBuy         bool
		AutoRaid        bool
		AutoRaidRecover bool
		Count           int
		Type            int
		QuestId         int
		QuestIds        string
		Fid             int
		Pt              int
		Htype           int
		Lv              int
		Hcid            int
		Version         int
		Res             int
		Bt              int
		Wc              int
		Wn              int
		Time            string
		D               int
		S               int
		Cc              int
		Oc              int
	}
	type args struct {
		u *clients.Metadata
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp map[string]interface{}
		wantRes  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quest{
				AutoSell:        tt.fields.AutoSell,
				AutoBuy:         tt.fields.AutoBuy,
				AutoRaid:        tt.fields.AutoRaid,
				AutoRaidRecover: tt.fields.AutoRaidRecover,
				Count:           tt.fields.Count,
				Type:            tt.fields.Type,
				QuestId:         tt.fields.QuestId,
				QuestIds:        tt.fields.QuestIds,
				Fid:             tt.fields.Fid,
				Pt:              tt.fields.Pt,
				Htype:           tt.fields.Htype,
				Lv:              tt.fields.Lv,
				Hcid:            tt.fields.Hcid,
				Version:         tt.fields.Version,
				Res:             tt.fields.Res,
				Bt:              tt.fields.Bt,
				Wc:              tt.fields.Wc,
				Wn:              tt.fields.Wn,
				Time:            tt.fields.Time,
				D:               tt.fields.D,
				S:               tt.fields.S,
				Cc:              tt.fields.Cc,
				Oc:              tt.fields.Oc,
			}
			gotResp, gotRes := q.GetTreasure(tt.args.u)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("quest.GetTreasure() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotRes != tt.wantRes {
				t.Errorf("quest.GetTreasure() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestGetEndpoint(t *testing.T) {
	type args struct {
		version int
	}
	tests := []struct {
		name         string
		args         args
		wantEndpoint string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEndpoint := GetEndpoint(tt.args.version); gotEndpoint != tt.wantEndpoint {
				t.Errorf("GetEndpoint() = %v, want %v", gotEndpoint, tt.wantEndpoint)
			}
		})
	}
}

func TestGetResultEndpoint(t *testing.T) {
	type args struct {
		version int
	}
	tests := []struct {
		name         string
		args         args
		wantEndpoint string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEndpoint := GetResultEndpoint(tt.args.version); gotEndpoint != tt.wantEndpoint {
				t.Errorf("GetResultEndpoint() = %v, want %v", gotEndpoint, tt.wantEndpoint)
			}
		})
	}
}

func Test_quest_getEndPostBody(t *testing.T) {
	type fields struct {
		AutoSell        bool
		AutoBuy         bool
		AutoRaid        bool
		AutoRaidRecover bool
		Count           int
		Type            int
		QuestId         int
		QuestIds        string
		Fid             int
		Pt              int
		Htype           int
		Lv              int
		Hcid            int
		Version         int
		Res             int
		Bt              int
		Wc              int
		Wn              int
		Time            string
		D               int
		S               int
		Cc              int
		Oc              int
	}
	tests := []struct {
		name     string
		fields   fields
		wantBody map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quest{
				AutoSell:        tt.fields.AutoSell,
				AutoBuy:         tt.fields.AutoBuy,
				AutoRaid:        tt.fields.AutoRaid,
				AutoRaidRecover: tt.fields.AutoRaidRecover,
				Count:           tt.fields.Count,
				Type:            tt.fields.Type,
				QuestId:         tt.fields.QuestId,
				QuestIds:        tt.fields.QuestIds,
				Fid:             tt.fields.Fid,
				Pt:              tt.fields.Pt,
				Htype:           tt.fields.Htype,
				Lv:              tt.fields.Lv,
				Hcid:            tt.fields.Hcid,
				Version:         tt.fields.Version,
				Res:             tt.fields.Res,
				Bt:              tt.fields.Bt,
				Wc:              tt.fields.Wc,
				Wn:              tt.fields.Wn,
				Time:            tt.fields.Time,
				D:               tt.fields.D,
				S:               tt.fields.S,
				Cc:              tt.fields.Cc,
				Oc:              tt.fields.Oc,
			}
			if gotBody := q.getEndPostBody(); !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("quest.getEndPostBody() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

func Test_quest_getPostBody(t *testing.T) {
	type fields struct {
		AutoSell        bool
		AutoBuy         bool
		AutoRaid        bool
		AutoRaidRecover bool
		Count           int
		Type            int
		QuestId         int
		QuestIds        string
		Fid             int
		Pt              int
		Htype           int
		Lv              int
		Hcid            int
		Version         int
		Res             int
		Bt              int
		Wc              int
		Wn              int
		Time            string
		D               int
		S               int
		Cc              int
		Oc              int
	}
	tests := []struct {
		name     string
		fields   fields
		wantBody map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &quest{
				AutoSell:        tt.fields.AutoSell,
				AutoBuy:         tt.fields.AutoBuy,
				AutoRaid:        tt.fields.AutoRaid,
				AutoRaidRecover: tt.fields.AutoRaidRecover,
				Count:           tt.fields.Count,
				Type:            tt.fields.Type,
				QuestId:         tt.fields.QuestId,
				QuestIds:        tt.fields.QuestIds,
				Fid:             tt.fields.Fid,
				Pt:              tt.fields.Pt,
				Htype:           tt.fields.Htype,
				Lv:              tt.fields.Lv,
				Hcid:            tt.fields.Hcid,
				Version:         tt.fields.Version,
				Res:             tt.fields.Res,
				Bt:              tt.fields.Bt,
				Wc:              tt.fields.Wc,
				Wn:              tt.fields.Wn,
				Time:            tt.fields.Time,
				D:               tt.fields.D,
				S:               tt.fields.S,
				Cc:              tt.fields.Cc,
				Oc:              tt.fields.Oc,
			}
			if gotBody := q.getPostBody(); !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("quest.getPostBody() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

func TestNewQuest(t *testing.T) {
	tests := []struct {
		name  string
		wantQ *quest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotQ := NewQuest(); !reflect.DeepEqual(gotQ, tt.wantQ) {
				t.Errorf("NewQuest() = %v, want %v", gotQ, tt.wantQ)
			}
		})
	}
}
