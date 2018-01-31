package session


import (
    "github.com/mong0520/ChainChronicleGo/utils"
    "github.com/mong0520/ChainChronicleGo/clients"
    "strconv"
    "time"
    "fmt"
    "net/url"
    "encoding/json"
    "log"
    "net/http"
    "strings"
    "io/ioutil"
)

var API = "session/login"

func Login(uid string, token string) (sid string) {
    //logger := utils.GetLogger()
    requestUrl := GetEndpoint()
    postBody := GetPostBody(uid, token)
    resp, _ := post(requestUrl, postBody)

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

func post(requestUrl string, body map[string]interface{}) (respMap map[string]interface{}, err error){

    //nowShort := strconv.Itoa(int(time.Now().Unix()))
    nowLong := strconv.Itoa(int(time.Now().UnixNano()))
    // Build up post data
    data := url.Values{}
    postBodyJson, err := json.Marshal(body)
    if err != nil{
        log.Println("Parameter error")
        return
    }

    postBodyJsonString := string(postBodyJson)
    data.Set("param", postBodyJsonString)
    req, err := http.NewRequest("POST", requestUrl, strings.NewReader(data.Encode()))

    // Build up query string
    queries := req.URL.Query()
    queries.Set("cnt", nowLong)
    req.URL.RawQuery = queries.Encode()

    //fmt.Println("Post Host = ", req.URL.String())
    //fmt.Println("Post Data = ", postBodyJsonString)
    if err != nil {
        fmt.Printf("http.NewRequest() error: %v\n", err)
        return nil, err
    }
    c := &http.Client{}
    resp, err := c.Do(req)
    if err != nil {
        fmt.Printf("http.Do() error: %v\n", err)
        return nil, err
    }
    defer resp.Body.Close()

    // Read response body
    ret, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("ioutil.ReadAll() error: %v\n", err)
        return nil, err
    }

    //fmt.Printf("read resp.Body successfully:\n%v\n", ret)
    // Unzip response body
    respMap, err = utils.DecodeResponse(ret)
    if err != nil{
        fmt.Println("Decode response error, %v", err)
        return nil, err
    }

    return respMap, nil
}