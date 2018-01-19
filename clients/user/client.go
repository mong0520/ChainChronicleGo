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

//def set_password(password, sid):
//poster = utils.poster.Poster
//url = "{0}/user/set_password".format(utils.global_config.get_hostname())
//data = {'pass': password}
//headers = {'Cookie': 'sid={0}'.format(sid)}
//cookies = {'sid': sid}
//# self.poster.set_header(headers)
//# self.poster.set_cookies(cookies)
//ret = poster.post_data(url, headers, cookies, **data)
//return ret
