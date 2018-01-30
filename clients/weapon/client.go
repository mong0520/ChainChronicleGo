package weapon

import (
    "github.com/mong0520/ChainChronicleGo/clients"
    "fmt"
    "bytes"
    "time"
    "strconv"
    "net/url"
    "net/http"
    "strings"
    "io/ioutil"
    "github.com/mong0520/ChainChronicleGo/utils"
)

func Compose(metadata *clients.Metadata, wiaList []int, eid int) (respMap map[string]interface{}, res int) {
    api := "weapon/compose"
    cookie := &http.Cookie{
        Name: "sid",
        Value: metadata.Sid,
    }

    // Prepare Post Data and query string
    tempNow := int(time.Now().UnixNano())
    nowShort := strconv.Itoa(int(time.Now().Unix()))
    nowHex := fmt.Sprintf("%x", tempNow)

    var buffer bytes.Buffer
    var queryString string
    for _, wia := range wiaList{
        buffer.WriteString(fmt.Sprintf("wia=%d&", wia))
    }
    buffer.WriteString(fmt.Sprintf("timestamp=%s&", nowShort))
    buffer.WriteString(fmt.Sprintf("cnt=%s", nowHex))

    if eid != -1{
        buffer.WriteString(fmt.Sprintf("&eid=%d", eid))
    }
    queryString = "?" + buffer.String()  // ?wia=.....&wia=xx&eid=14
    requestUrl := fmt.Sprintf("%s/%s%s", clients.HOST, api, queryString)
    postData := fmt.Sprintf("nature=%s", url.QueryEscape(queryString))
    //fmt.Println(requestUrl)
    //fmt.Println(postData)

    // Prepare HTTP Request
    req, err := http.NewRequest("POST", requestUrl, strings.NewReader(postData))
    if err != nil {
        fmt.Printf("http.NewRequest() error: %v\n", err)
        return nil, -1
    }
    req.AddCookie(cookie)

    c := &http.Client{}
    resp, err := c.Do(req)
    if err != nil {
        fmt.Printf("http.Do() error: %v\n", err)
        return nil, -1
    }
    defer resp.Body.Close()

    // Read response body
    ret, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("ioutil.ReadAll() error: %v\n", err)
        return nil, -1
    }

    //fmt.Printf("read resp.Body successfully:\n%v\n", ret)
    // Unzip response body
    respMap, err = utils.DecodeResponse(ret)
    if err != nil{
        fmt.Println("Decode response error, %v", err)
        return nil, -1
    }
    //fmt.Println(utils.Map2JsonString(respMap))
    return respMap, int(respMap["res"].(float64))
}
