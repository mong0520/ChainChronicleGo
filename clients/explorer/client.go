package explorer

import (
	"fmt"

	"github.com/mong0520/ChainChronicleGo/clients/general"
)

type Pickup struct {
	Home       int    `json:"home"`
	Jobtype    int    `json:"jobtype"`
	LocationID int    `json:"location_id"`
	Text       string `json:"text"`
	Weapontype int    `json:"weapontype"`
}

func GetExplorerList(sid string) (resp map[string]interface{}, res int) {
	action := "list"
	api := fmt.Sprintf("%s/%s", "explorer", action)
	param := map[string]interface{}{}
	return general.GeneralAction(api, sid, param)
}

func GetExplorerResult(sid string, eid int) (resp map[string]interface{}, res int) {
	action := "result"
	api := fmt.Sprintf("%s/%s", "explorer", action)
	param := map[string]interface{}{
		"explorer_idx": eid,
	}
	return general.GeneralAction(api, sid, param)
}

func StartExplorer(sid string, parameter map[string]int) (resp map[string]interface{}, res int) {
	action := "entry"
	api := fmt.Sprintf("%s/%s", "explorer", action)
	param := map[string]interface{}{
		"explorer_idx": parameter["explorer_idx"],
		"location_id":  parameter["location_id"],
		"card_idx":     parameter["card_idx"],
		"pickup":       parameter["pickup"],
		"interval":     parameter["interval"],
		"helper1":      588707,
		"helper2":      588707,
	}
	return general.GeneralAction(api, sid, param)
}

func FinishExplorer(sid string, eid int) (resp map[string]interface{}, res int) {
	action := "finish"
	api := fmt.Sprintf("%s/%s", "explorer", action)
	param := map[string]interface{}{
		"explorer_idx": eid,
	}
	return general.GeneralAction(api, sid, param)
}

func CancelExplorer(sid string, eid int) (resp map[string]interface{}, res int) {
	action := "cancel"
	api := fmt.Sprintf("%s/%s", "explorer", action)
	param := map[string]interface{}{
		"explorer_idx": eid,
	}
	return general.GeneralAction(api, sid, param)
}
