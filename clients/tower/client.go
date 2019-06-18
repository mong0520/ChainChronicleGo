package tower

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/general"
	"github.com/mong0520/ChainChronicleGo/utils"
)

// Type = 39 in alldata
type TowerInfo struct {
	Data struct {
		ArrivalExQuest  int           `json:"arrival_ex_quest"`
		ArrivalQuest    int           `json:"arrival_quest"`
		Cleared         int           `json:"cleared"`
		CloseTime       int           `json:"close_time"`
		ExCleared       int           `json:"ex_cleared"`
		ExFloor         int           `json:"ex_floor"`
		ExQuest         int           `json:"ex_quest"`
		ExStatus        int           `json:"ex_status"`
		ExUseCids       []interface{} `json:"ex_use_cids"`
		Floor           int           `json:"floor"`
		MaxExFloor      int           `json:"max_ex_floor"`
		MaxFloor        int           `json:"max_floor"`
		Point           int           `json:"point"`
		Quest           int           `json:"quest"`
		RecordExFloor   int           `json:"record_ex_floor"`
		RecordExQuest   int           `json:"record_ex_quest"`
		RecordExWave    int           `json:"record_ex_wave"`
		RecordFloor     int           `json:"record_floor"`
		RecordMaxExWave int           `json:"record_max_ex_wave"`
		RecordMaxWave   int           `json:"record_max_wave"`
		RecordQuest     int           `json:"record_quest"`
		RecordWave      int           `json:"record_wave"`
		Status          int           `json:"status"`
		TowerID         int           `json:"tower_id"`
		UID             int           `json:"uid"`
		UseCids         []interface{} `json:"use_cids"`
	} `json:"data"`
}

func EnterExTower(metadata *clients.Metadata, twid int, floor int, snum int, pt int) (resp map[string]interface{}, res int) {
	api := "ex_tower/entry"
	param := map[string]interface{}{
		"twid":  twid,
		"floor": floor,
		"snum":  snum,
		"fid":   1936248, //不能選自己
		"htype": 0,
		"pt":    pt,
	}
	return general.GeneralAction(api, metadata.Sid, param)
}

func ExitExTower(metadata *clients.Metadata, twid int, wc int) (resp map[string]interface{}, res int) {
	api := "ex_tower/result"
	param := map[string]interface{}{
		"res":  1,
		"twid": twid,
		"wc":   wc,
		"time": "10.0",
		"d":    1,
		"s":    0,
	}

	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
	resp, _ = utils.PostV2(requestUrl, "wvt=%5b%7b%22wave_num%22%3a1%2c%22time%22%3a1529%7d%2c%7b%22wave_num%22%3a2%2c%22time%22%3a3083%7d%5d&mission=%7b%22cid%22%3a%5b6280%2c6251%2c11216%2c6222%2c8159%2c4202%5d%2c%22sid%22%3a%5b6055%2c2043%2c11217%2c248%2c8158%2c58013%5d%2c%22fid%22%3a%5b8900%5d%2c%22hrid%22%3a%5b7206%5d%2c%22ms%22%3a0%2c%22md%22%3a11153%2c%22sc%22%3a%7b%220%22%3a0%2c%221%22%3a0%2c%222%22%3a0%2c%223%22%3a0%2c%224%22%3a0%7d%2c%22es%22%3a0%2c%22at%22%3a0%2c%22he%22%3a0%2c%22da%22%3a0%2c%22ba%22%3a0%2c%22bu%22%3a0%2c%22job%22%3a%7b%220%22%3a4%2c%221%22%3a1%2c%222%22%3a0%2c%223%22%3a3%2c%224%22%3a0%7d%2c%22weapon%22%3a%7b%220%22%3a4%2c%221%22%3a0%2c%222%22%3a0%2c%223%22%3a1%2c%224%22%3a0%2c%225%22%3a3%2c%228%22%3a0%2c%229%22%3a0%2c%2210%22%3a0%7d%2c%22box%22%3a1%2c%22um%22%3a%7b%221%22%3a0%2c%222%22%3a0%2c%223%22%3a0%7d%2c%22fj%22%3a0%2c%22fw%22%3a0%2c%22fo%22%3a0%2c%22mlv%22%3a80%2c%22mbl%22%3a214%2c%22udj%22%3a0%2c%22sdmg%22%3a70827%2c%22tp%22%3a0%2c%22gma%22%3a8%2c%22gmr%22%3a2%2c%22gmp%22%3a0%2c%22stp%22%3a0%2c%22auto%22%3a1%2c%22uh%22%3a%7b%2210%22%3a3%2c%2215%22%3a1%2c%223%22%3a2%2c%228%22%3a1%2c%221%22%3a1%7d%2c%22cc%22%3a1%2c%22bf_atk%22%3a0%2c%22bf_hp%22%3a0%2c%22bf_spd%22%3a0%7d&nature=cnt%3d16b6444a389%26d%3d1%26mission%3d%257b%2522cid%2522%253a%255b6280%252c6251%252c11216%252c6222%252c8159%252c4202%255d%252c%2522sid%2522%253a%255b6055%252c2043%252c11217%252c248%252c8158%252c58013%255d%252c%2522fid%2522%253a%255b8900%255d%252c%2522hrid%2522%253a%255b7206%255d%252c%2522ms%2522%253a0%252c%2522md%2522%253a11153%252c%2522sc%2522%253a%257b%25220%2522%253a0%252c%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a0%252c%25224%2522%253a0%257d%252c%2522es%2522%253a0%252c%2522at%2522%253a0%252c%2522he%2522%253a0%252c%2522da%2522%253a0%252c%2522ba%2522%253a0%252c%2522bu%2522%253a0%252c%2522job%2522%253a%257b%25220%2522%253a4%252c%25221%2522%253a1%252c%25222%2522%253a0%252c%25223%2522%253a3%252c%25224%2522%253a0%257d%252c%2522weapon%2522%253a%257b%25220%2522%253a4%252c%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a1%252c%25224%2522%253a0%252c%25225%2522%253a3%252c%25228%2522%253a0%252c%25229%2522%253a0%252c%252210%2522%253a0%257d%252c%2522box%2522%253a1%252c%2522um%2522%253a%257b%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a0%257d%252c%2522fj%2522%253a0%252c%2522fw%2522%253a0%252c%2522fo%2522%253a0%252c%2522mlv%2522%253a80%252c%2522mbl%2522%253a214%252c%2522udj%2522%253a0%252c%2522sdmg%2522%253a70827%252c%2522tp%2522%253a0%252c%2522gma%2522%253a8%252c%2522gmr%2522%253a2%252c%2522gmp%2522%253a0%252c%2522stp%2522%253a0%252c%2522auto%2522%253a1%252c%2522uh%2522%253a%257b%252210%2522%253a3%252c%252215%2522%253a1%252c%25223%2522%253a2%252c%25228%2522%253a1%252c%25221%2522%253a1%257d%252c%2522cc%2522%253a1%252c%2522bf_atk%2522%253a0%252c%2522bf_hp%2522%253a0%252c%2522bf_spd%2522%253a0%257d%26res%3d1%26s%3d0%26time%3d1.45%26twid%3d15%26wc%3d3%26wvt%3d%255b%257b%2522wave_num%2522%253a1%252c%2522time%2522%253a1529%257d%252c%257b%2522wave_num%2522%253a2%252c%2522time%2522%253a3083%257d%255d", param, metadata.Sid)
	res = int(resp["res"].(float64))

	return resp, res
}

func EnterTower(metadata *clients.Metadata, twid int, floor int, snum int, pt int) (resp map[string]interface{}, res int) {
	api := "tower/entry"
	param := map[string]interface{}{
		"twid":  twid,
		"floor": floor,
		"snum":  snum,
		"fid":   1936248, //不能選自己
		"htype": 0,
		"pt":    pt,
	}
	return general.GeneralAction(api, metadata.Sid, param)
}

func ExitTower(metadata *clients.Metadata, twid int, wc int) (resp map[string]interface{}, res int) {
	api := "tower/result"
	param := map[string]interface{}{
		"res":  1,
		"twid": twid,
		"wc":   wc,
		"time": "10.0",
		"d":    1,
		"s":    0,
	}

	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
	resp, _ = utils.PostV2(requestUrl, "wvt=%5b%7b%22wave_num%22%3a1%2c%22time%22%3a1529%7d%2c%7b%22wave_num%22%3a2%2c%22time%22%3a3083%7d%5d&mission=%7b%22cid%22%3a%5b6280%2c6251%2c11216%2c6222%2c8159%2c4202%5d%2c%22sid%22%3a%5b6055%2c2043%2c11217%2c248%2c8158%2c58013%5d%2c%22fid%22%3a%5b8900%5d%2c%22hrid%22%3a%5b7206%5d%2c%22ms%22%3a0%2c%22md%22%3a11153%2c%22sc%22%3a%7b%220%22%3a0%2c%221%22%3a0%2c%222%22%3a0%2c%223%22%3a0%2c%224%22%3a0%7d%2c%22es%22%3a0%2c%22at%22%3a0%2c%22he%22%3a0%2c%22da%22%3a0%2c%22ba%22%3a0%2c%22bu%22%3a0%2c%22job%22%3a%7b%220%22%3a4%2c%221%22%3a1%2c%222%22%3a0%2c%223%22%3a3%2c%224%22%3a0%7d%2c%22weapon%22%3a%7b%220%22%3a4%2c%221%22%3a0%2c%222%22%3a0%2c%223%22%3a1%2c%224%22%3a0%2c%225%22%3a3%2c%228%22%3a0%2c%229%22%3a0%2c%2210%22%3a0%7d%2c%22box%22%3a1%2c%22um%22%3a%7b%221%22%3a0%2c%222%22%3a0%2c%223%22%3a0%7d%2c%22fj%22%3a0%2c%22fw%22%3a0%2c%22fo%22%3a0%2c%22mlv%22%3a80%2c%22mbl%22%3a214%2c%22udj%22%3a0%2c%22sdmg%22%3a70827%2c%22tp%22%3a0%2c%22gma%22%3a8%2c%22gmr%22%3a2%2c%22gmp%22%3a0%2c%22stp%22%3a0%2c%22auto%22%3a1%2c%22uh%22%3a%7b%2210%22%3a3%2c%2215%22%3a1%2c%223%22%3a2%2c%228%22%3a1%2c%221%22%3a1%7d%2c%22cc%22%3a1%2c%22bf_atk%22%3a0%2c%22bf_hp%22%3a0%2c%22bf_spd%22%3a0%7d&nature=cnt%3d16b6444a389%26d%3d1%26mission%3d%257b%2522cid%2522%253a%255b6280%252c6251%252c11216%252c6222%252c8159%252c4202%255d%252c%2522sid%2522%253a%255b6055%252c2043%252c11217%252c248%252c8158%252c58013%255d%252c%2522fid%2522%253a%255b8900%255d%252c%2522hrid%2522%253a%255b7206%255d%252c%2522ms%2522%253a0%252c%2522md%2522%253a11153%252c%2522sc%2522%253a%257b%25220%2522%253a0%252c%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a0%252c%25224%2522%253a0%257d%252c%2522es%2522%253a0%252c%2522at%2522%253a0%252c%2522he%2522%253a0%252c%2522da%2522%253a0%252c%2522ba%2522%253a0%252c%2522bu%2522%253a0%252c%2522job%2522%253a%257b%25220%2522%253a4%252c%25221%2522%253a1%252c%25222%2522%253a0%252c%25223%2522%253a3%252c%25224%2522%253a0%257d%252c%2522weapon%2522%253a%257b%25220%2522%253a4%252c%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a1%252c%25224%2522%253a0%252c%25225%2522%253a3%252c%25228%2522%253a0%252c%25229%2522%253a0%252c%252210%2522%253a0%257d%252c%2522box%2522%253a1%252c%2522um%2522%253a%257b%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a0%257d%252c%2522fj%2522%253a0%252c%2522fw%2522%253a0%252c%2522fo%2522%253a0%252c%2522mlv%2522%253a80%252c%2522mbl%2522%253a214%252c%2522udj%2522%253a0%252c%2522sdmg%2522%253a70827%252c%2522tp%2522%253a0%252c%2522gma%2522%253a8%252c%2522gmr%2522%253a2%252c%2522gmp%2522%253a0%252c%2522stp%2522%253a0%252c%2522auto%2522%253a1%252c%2522uh%2522%253a%257b%252210%2522%253a3%252c%252215%2522%253a1%252c%25223%2522%253a2%252c%25228%2522%253a1%252c%25221%2522%253a1%257d%252c%2522cc%2522%253a1%252c%2522bf_atk%2522%253a0%252c%2522bf_hp%2522%253a0%252c%2522bf_spd%2522%253a0%257d%26res%3d1%26s%3d0%26time%3d1.45%26twid%3d15%26wc%3d3%26wvt%3d%255b%257b%2522wave_num%2522%253a1%252c%2522time%2522%253a1529%257d%252c%257b%2522wave_num%2522%253a2%252c%2522time%2522%253a3083%257d%255d", param, metadata.Sid)
	res = int(resp["res"].(float64))

	return resp, res
}

func AddTicket(metadata *clients.Metadata, twid int, item_type int, item_id int) (resp map[string]interface{}, res int) {
	api := "tower/add_ticket"
	param := map[string]interface{}{
		"twid":    twid,
		"type":    item_type,
		"item_id": item_id,
	}
	return general.GeneralAction(api, metadata.Sid, param)
}

func GetCurrentTowerInfo(metadata *clients.Metadata) (TowerInfo, error) {
	towerInfo := TowerInfo{}
	for _, data := range metadata.AllDataS.Body {
		if data.Type == 39 {
			towerInfoData, err := json.Marshal(data)
			if err != nil {
				return towerInfo, err
			}
			json.Unmarshal(towerInfoData, &towerInfo)
			return towerInfo, nil
		}
	}
	return towerInfo, errors.New("data type 39(tower info) not found")
}
