package uzu

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/utils"
)

type UzuHistoryStruct []struct {
	ClearList      []int `json:"clear_list"`
	ClearedLv      int   `json:"cleared_lv"`
	Lap            int   `json:"lap"`
	LastScheduleID int   `json:"last_schedule_id"`
	Point          int   `json:"point"`
	Soul           int   `json:"soul"`
	SoulLap        int   `json:"soul_lap"`
	TotalLap       int   `json:"total_lap"`
	TotalSoul      int   `json:"total_soul"`
	TotalSoulLap   int   `json:"total_soul_lap"`
	UID            int   `json:"uid"`
	UzuID          int   `json:"uzu_id"`
}

type UzuDataStruct struct {
	Deceive int `json:"deceive"`
	Info    struct {
		UzuKeyInit        int `json:"uzu_key_init"`
		UzuKeyMax         int `json:"uzu_key_max"`
		UzuKeyRecover1Day int `json:"uzu_key_recover_1day"`
		UzuKeyRecoverItem int `json:"uzu_key_recover_item"`
		UzuKeyRecoverUse  int `json:"uzu_key_recover_use"`
		UzuKeyUpdateTime  int `json:"uzu_key_update_time"`
	} `json:"info"`
	Res   int `json:"res"`
	Stage []struct {
		StageID   int `json:"stage_id"`
		StageList []struct {
			BattlePoint int   `json:"battle_point"`
			Bg          int   `json:"bg"`
			Bgm         int   `json:"bgm"`
			BossFlag    int   `json:"boss_flag"`
			BtMana      int   `json:"bt_mana"`
			ClearedLv   int   `json:"cleared_lv"`
			DfIcon      int   `json:"df_icon"`
			Difficulty  int   `json:"difficulty"`
			Eneset      []int `json:"eneset"`
			HomeRate    int   `json:"home_rate"`
			Lv          int   `json:"lv"`
			Reward      struct {
				ID   int    `json:"id"`
				Type string `json:"type"`
				Val  int    `json:"val"`
			} `json:"reward"`
			Stage    int    `json:"stage"`
			Strength int    `json:"strength"`
			Title    string `json:"title"`
		} `json:"stage_list"`
		UseKey int `json:"use_key"`
	} `json:"stage"`
	Uzu []struct {
		BossBorder []int  `json:"boss_border"`
		BossID     int    `json:"boss_id"`
		Name       string `json:"name"`
		RewardMap  []struct {
			RewardList []struct {
				Border int    `json:"border"`
				ID     int    `json:"id"`
				Type   string `json:"type"`
				Val    int    `json:"val"`
			} `json:"reward_list"`
		} `json:"reward_map"`
		Schedule []struct {
			Data struct {
				Debuff struct {
					Atk   float64 `json:"atk"`
					Def   float64 `json:"def"`
					Hp    int     `json:"hp"`
					Speed float64 `json:"speed"`
				} `json:"debuff"`
				Home int `json:"home"`
				Pos  []struct {
					Area int `json:"area"`
					Posx int `json:"posx"`
					Posy int `json:"posy"`
				} `json:"pos"`
				StageID   int    `json:"stage_id"`
				UzuID     int    `json:"uzu_id"`
				WebBanner string `json:"web_banner"`
				WebText   string `json:"web_text"`
			} `json:"data"`
			End        int `json:"end"`
			ScheduleID int `json:"schedule_id"`
			Start      int `json:"start"`
		} `json:"schedule"`
		UzuID int `json:"uzu_id"`
	} `json:"uzu"`
}

var UzuData *UzuDataStruct

func init() {
	UzuData = &UzuDataStruct{}
}

func GetUzuInfo(sid string) (*UzuDataStruct, int) {
	api := "data/uzuinfo"
	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
	postBody := map[string]interface{}{}
	resp, _ := utils.PostV2(requestUrl, "", postBody, sid)
	res := int(resp["res"].(float64))
	jsonString, _ := json.Marshal(resp)
	json.Unmarshal([]byte(jsonString), UzuData)

	return UzuData, res
}

func (u *UzuDataStruct) GetCurrentScheduleID(uzuID int) int {

	now := int(time.Now().Unix())

	for _, schedule := range u.Uzu[uzuID-1].Schedule {
		if now >= schedule.Start && now <= schedule.End {
			return schedule.ScheduleID
		}
	}

	return -1
}
