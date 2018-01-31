package raid

import (
	"fmt"
	"github.com/icza/dyno"
	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/general"
	"github.com/mong0520/ChainChronicleGo/utils"
	"strings"
	"time"
)

type BossInfo struct {
	BossID    int `json:"boss_id"`
	BossParam struct {
		AngerFlg  bool `json:"anger_flg"`
		Bg        int  `json:"bg"`
		BossEneid int  `json:"boss_eneid"`
		FieldType int  `json:"field_type"`
		GachaID   int  `json:"gacha_id"`
		Hp        int  `json:"hp"`
		HpMax     int  `json:"hpMax"`
		Lv        int  `json:"lv"`
		Strength  int
	} `json:"boss_param"`
	Discoverer int
	Dname      string `json:"discoverer_name"`
	Own        bool
	Validtime  int
}

func RaidList(sid string) (resp map[string]interface{}, res int) {
	api := "raid/list"
	param := map[string]interface{}{}
	return general.GeneralAction(api, sid, param)
}

func GetRaidBossInfo(sid string) (bossInfo *BossInfo) {

	if ret, err := RaidList(sid); err != 0 {
		fmt.Println("Unable to get raid list", err)
		return nil
	} else {
		//fmt.Println(utils.Map2JsonString(ret))
		if dataArray, err := dyno.GetSlice(ret, "body", 0, "data"); err != nil {
			fmt.Println("err", err)
			return nil
		} else {
			for _, d := range dataArray {
				//fmt.Println(utils.Map2JsonString(d.(map[string]interface{})))
				if _, ok := d.(map[string]interface{})["own"]; ok && d.(map[string]interface{})["own"].(bool) == true {
					bossInfo := &BossInfo{}
					utils.Map2Struct(d.(map[string]interface{}), bossInfo)
					return bossInfo
				}
			}
			return nil
		}
	}
}

func (bossInfo *BossInfo) StartQuest(metadata *clients.Metadata, param map[string]interface{}) (resp map[string]interface{}, res int) {

	api := "raid/entry"
	param["bid"] = bossInfo.BossID
	param["fid"] = 1965350
	param["pt"] = 0
	param["use"] = 1
	return general.GeneralAction(api, metadata.Sid, param)
}

func (bossInfo *BossInfo) EndQuest(metadata *clients.Metadata, param map[string]interface{}) (resp map[string]interface{}, res int) {

	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, "raid/result")
	param["bid"] = bossInfo.BossID
	param["res"] = 1
	param["damage"] = 9994500
	param["t"] = 15
	param["ch"] = 1
	param["eh"] = 1

	nowHex := fmt.Sprintf("%x", int(time.Now().UnixNano()))
	payloadString := "mission=%7b%22cid%22%3a%5b1032%2c57%2c7505%2c3022%2c1021%2c38%5d%2c%22fid%22%3a43%2c%22ms%22%3a0%2c%22md%22%3a198601%2c%22sc%22%3a%7b%221%22%3a0%2c%222%22%3a0%2c%223%22%3a0%2c%224%22%3a0%7d%2c%22es%22%3a0%2c%22at%22%3a0%2c%22he%22%3a0%2c%22da%22%3a0%2c%22ba%22%3a0%2c%22bu%22%3a0%2c%22job%22%3a%7b%220%22%3a3%2c%221%22%3a1%2c%222%22%3a1%2c%223%22%3a1%2c%224%22%3a1%7d%2c%22weapon%22%3a%7b%220%22%3a2%2c%221%22%3a0%2c%222%22%3a1%2c%223%22%3a1%2c%224%22%3a1%2c%225%22%3a1%2c%228%22%3a1%2c%229%22%3a0%2c%2210%22%3a0%7d%2c%22box%22%3a1%2c%22um%22%3a%7b%221%22%3a0%2c%222%22%3a0%2c%223%22%3a0%7d%2c%22fj%22%3a1%2c%22fw%22%3a3%2c%22fo%22%3a0%2c%22cc%22%3a1%7d&nature=bid%3dBID%26cnt%3dCNT%26damage%3d994500%26mission%3d%257b%2522cid%2522%253a%255b1032%252c57%252c7505%252c3022%252c1021%252c38%255d%252c%2522fid%2522%253a43%252c%2522ms%2522%253a0%252c%2522md%2522%253a198601%252c%2522sc%2522%253a%257b%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a0%252c%25224%2522%253a0%257d%252c%2522es%2522%253a0%252c%2522at%2522%253a0%252c%2522he%2522%253a0%252c%2522da%2522%253a0%252c%2522ba%2522%253a0%252c%2522bu%2522%253a0%252c%2522job%2522%253a%257b%25220%2522%253a3%252c%25221%2522%253a1%252c%25222%2522%253a1%252c%25223%2522%253a1%252c%25224%2522%253a1%257d%252c%2522weapon%2522%253a%257b%25220%2522%253a2%252c%25221%2522%253a0%252c%25222%2522%253a1%252c%25223%2522%253a1%252c%25224%2522%253a1%252c%25225%2522%253a1%252c%25228%2522%253a1%252c%25229%2522%253a0%252c%252210%2522%253a0%257d%252c%2522box%2522%253a1%252c%2522um%2522%253a%257b%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a0%257d%252c%2522fj%2522%253a1%252c%2522fw%2522%253a3%252c%2522fo%2522%253a0%252c%2522cc%2522%253a1%257d%26res%3d1%26t%3d15"
	payloadString = strings.Replace(payloadString, "BID", string(param["bid"].(int)), 1)
	payloadString = strings.Replace(payloadString, "CNT", nowHex, 1)
	resp, _ = utils.PostV2(requestUrl, payloadString, param, metadata.Sid)
	res = int(resp["res"].(float64))
	return resp, res
}

func (bossInfo *BossInfo) GetBonus(metadata *clients.Metadata, param map[string]interface{}) (resp map[string]interface{}, res int) {

	api := "raid/record"
	param["bid"] = bossInfo.BossID
	return general.GeneralAction(api, metadata.Sid, param)
}
