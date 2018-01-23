package session


import (
    "github.com/mong0520/ChainChronicleGo/utils"
    "github.com/mong0520/ChainChronicleGo/clients"
    "strconv"
    "time"
    "fmt"
)

var API = "session/login"

func Login(uid string, token string) (sid string) {
    //logger := utils.GetLogger()
    requestUrl := GetEndpoint()
    postBody := GetPostBody(uid, token)
    resp, _ := utils.Post(requestUrl, postBody)
    //fmt.Printf("Response = %v\n", utils.Map2JsonString(resp))
    sid = resp["login"].(map[string]interface{})["sid"].(string)
    return sid
}


func GetEndpoint()(endpoint string){
    endpoint = fmt.Sprintf("%s/%s", clients.HOST, API)
    return endpoint
}

func GetPostBody(uid string, token string)(body map[string]interface{}){
    nowShort := strconv.Itoa(int(time.Now().Unix()))
    body = map[string]interface{}{
        "APP": map[string]interface{}{
            "Version":  "3.2.2",
            "Revision": "2014",
            "time":     nowShort,
            "Lang":     "Chinese",
        },
        "DEV": map[string]interface{}{
            "UserUniqueID": uid,
            "Token":        token,
            "OS":           "1",
        },
    }
    return body
}