package general

import (
    "github.com/mong0520/ChainChronicleGo/utils"
    "fmt"
    "github.com/mong0520/ChainChronicleGo/clients"
)

func GeneralAction(api string, sid string, parateter map[string]interface{}) (resp map[string]interface{}, res int) {
    requestUrl := fmt.Sprintf("%s/%s", clients.HOST, api)
    resp, _ = utils.PostV2(requestUrl, "", parateter, sid)
    res = int(resp["res"].(float64))
    return resp, res
}