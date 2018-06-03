package tower

import (
	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/general"
)

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
	return general.GeneralAction(api, metadata.Sid, param)
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
