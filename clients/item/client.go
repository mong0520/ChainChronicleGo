package item

import (
	"fmt"
	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/utils"
)

//json string to map https://play.golang.org/p/m-pUtgwFsQs
var api = "token"

var AP_FRUIT = "ap_fruit"
var WEAPON_SWORD = "itm_weapon"
var WEAPON_BOW = "itm_weapon_bow"
var WEAPON_MAGIC = "itm_weapon_magic"
var CHAR = "char"
var WEAPON = "weapon"

var iteaMapping = map[string]interface{}{
	AP_FRUIT:     map[string]interface{}{"kind": "item", "id": 1, "type": "item", "price": 10, "val": 1},
	WEAPON_SWORD: map[string]interface{}{"kind": "item", "id": 96019, "type": "weapon_ev", "price": 10, "val": 1},
	WEAPON_BOW:   map[string]interface{}{"kind": "item", "id": 96064, "type": "weapon_ev", "price": 10, "val": 1},
	WEAPON_MAGIC: map[string]interface{}{"kind": "item", "val": 1, "id": 96126, "type": "weapon_ev", "price": 10},
	CHAR:         map[string]interface{}{"kind": "item", "type": "chara_rf", "price": 30, "val": 1, "id": 90904},
	WEAPON:       map[string]interface{}{"kind": "item", "price": 30, "val": 1, "id": 93902, "type": "weapon_rf"}}

func BuyItemByType(itemType string, sid string) (resp map[string]interface{}, res int) {
	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
	if itemData, ok := iteaMapping[itemType] ; ok{
        postBody := itemData.(map[string]interface{})
        resp, _ = utils.PostV2(requestUrl, "", postBody, sid)
        res = int(resp["res"].(float64))
        return resp, res
    }else{
        return nil, -1
    }
}
