package user

import (
	"fmt"

	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/general"
	"github.com/mong0520/ChainChronicleGo/models"
	"github.com/mong0520/ChainChronicleGo/utils"
)

func GetUserBook(sid string) (*models.Book, int) {
	api := "user/book"
	book := &models.Book{}
	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
	postBody := map[string]interface{}{}
	resp, _ := utils.PostV2(requestUrl, "", postBody, sid)
	res := int(resp["res"].(float64))
	if res != 0 {
		return nil, res
	}
	// fmt.Println(utils.Map2JsonString(resp))
	if err := utils.Map2Struct(resp, book); err != nil {
		return nil, -1
	}

	return book, 0

}

func GetAllData(sid string) (resp map[string]interface{}, res int) {
	api := "user/all_data"
	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
	postBody := map[string]interface{}{}
	resp, _ = utils.PostV2(requestUrl, "", postBody, sid)
	res = int(resp["res"].(float64))
	return resp, res
}

func GetAccount(sid string) (resp map[string]interface{}, res int) {
	api := "user/get_account"
	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
	postBody := map[string]interface{}{}
	resp, _ = utils.PostV2(requestUrl, "", postBody, sid)
	res = int(resp["res"].(float64))
	return resp, res
}

func SetPassword(password string, sid string) (resp map[string]interface{}, res int) {
	api := "user/set_password"
	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
	postBody := map[string]interface{}{
		"pass": password,
	}
	resp, _ = utils.PostV2(requestUrl, "", postBody, sid)
	res = int(resp["res"].(float64))
	return resp, res
}

func Takeover(uuid string, account string, password string) (resp map[string]interface{}, res int) {
	api := "user/takeover"
	requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
	postBody := map[string]interface{}{
		"uuid":    uuid,
		"account": account,
		"pass":    password,
	}
	resp, _ = utils.PostV2(requestUrl, "", postBody, "")
	res = int(resp["res"].(float64))
	return resp, res
}

func RecoveryAp(itemType int, itemId int, sid string) (resp map[string]interface{}, res int) {
	api := "user/recover_ap"
	param := map[string]interface{}{
		"type":    itemType,
		"item_id": itemId,
	}
	return general.GeneralAction(api, sid, param)
}

func RecoveryBp(itemType int, itemId int, sid string) (resp map[string]interface{}, res int) {
	api := "user/recover_bp"
	param := map[string]interface{}{
		"type":    itemType,
		"item_id": itemId,
	}
	return general.GeneralAction(api, sid, param)
}
