package user


import (
    "github.com/mong0520/ChainChronicleGo/utils"
    "github.com/mong0520/ChainChronicleGo/clients"
    "fmt"
)



func GetAllData(sid string) (resp map[string]interface{}, res int) {
    api := "user/all_data"
    requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
    postBody := map[string]interface{}{}
    resp, _ = utils.PostV2(requestUrl, "", postBody, sid)
    res = int(resp["res"].(float64))
    return resp, res
}