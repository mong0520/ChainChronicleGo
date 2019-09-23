package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/tower"
	"github.com/mong0520/ChainChronicleGo/clients/user"
	"github.com/mong0520/ChainChronicleGo/clients/web"
	"github.com/mong0520/ChainChronicleGo/models"
	"github.com/mong0520/ChainChronicleGo/utils"
	"golang.org/x/net/proxy"
)

var API = "session/login"

type MyError struct {
	Msg string
}

func Error() {

}

func Login(uid string, token string, useProxy bool) (sid string, err error) {
	//logger := utils.GetLogger()
	requestUrl := GetEndpoint()
	postBody := GetPostBody(uid, token)
	resp, _ := post(requestUrl, postBody, useProxy)

	//fmt.Printf("Response = %v\n", utils.Map2JsonString(resp))
	if _, ok := resp["login"]; ok {
		sid = resp["login"].(map[string]interface{})["sid"].(string)
		return sid, nil
	} else {
		return "", errors.New(fmt.Sprintf("%s: %+v\n", "Login failed", resp))
	}

}

func GetEndpoint() (endpoint string) {
	endpoint = fmt.Sprintf("%s/%s", clients.HOST, API)
	return endpoint
}

func GetPostBody(uid string, token string) (body map[string]interface{}) {
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

func post(requestUrl string, body map[string]interface{}, useProxy bool) (respMap map[string]interface{}, err error) {

	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	// setup a http client
	httpTransport := &http.Transport{}
	//httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial

	//nowShort := strconv.Itoa(int(time.Now().Unix()))
	nowLong := strconv.Itoa(int(time.Now().UnixNano()))
	// Build up post data
	data := url.Values{}
	postBodyJson, err := json.Marshal(body)
	if err != nil {
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
	switch useProxy {
	case true:
		fmt.Println("Use socks proxy")
		c = &http.Client{Transport: httpTransport}
	case false:
		c = &http.Client{}
	default:
		c = &http.Client{}
	}
	//c := &http.Client{Transport: httpTransport}
	//c := &http.Client{}
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
	if err != nil {
		fmt.Printf("Decode response error, %v\n", err)
		return nil, err
	}

	return respMap, nil
}

func GetSummaryStatus(sid string) string {
	alldata, _ := user.GetAllData(sid)
	var result strings.Builder
	targets := []string{"comment", "uid", "heroName", "open_id", "lv", "cardMax", "accept_disciple", "name",
		"friendCnt", "only_friend_disciple", "staminaMax", "zuLastRefilledScheduleId", "uzu_key"}
	itemMapping := map[int]string{
		7:  "轉蛋卷",
		10: "金幣",
		11: "聖靈幣",
		13: "戒指",
		15: "賭場幣",
		20: "轉蛋幣",
		39: "幸運球",
	}
	specialData := alldata["body"].([]interface{})[8].(map[string]interface{})["data"]
	stoneCount := alldata["body"].([]interface{})[12].(map[string]interface{})["data"]
	msg := fmt.Sprintf("精靈石 = %.0f\n", stoneCount.(float64))
	result.WriteString(msg)
	for _, item := range specialData.([]interface{}) {
		itemId := item.(map[string]interface{})["item_id"]
		cnt := item.(map[string]interface{})["cnt"]
		//logger.Info(itemId, reflect.TypeOf(itemId))
		//logger.Info(cnt, reflect.TypeOf(cnt))
		//fmt.Println(itemId)
		if val, ok := itemMapping[int(itemId.(float64))]; ok {
			switch reflect.TypeOf(cnt).Kind() {
			case reflect.String:
				msg = fmt.Sprintf("%s = %s\n", val, cnt.(string))
				result.WriteString(msg)
			case reflect.Float64:
				msg = fmt.Sprintf("%s = %.0f\n", val, cnt.(float64))
				result.WriteString(msg)
			}
		}
	}

	userData := alldata["body"].([]interface{})[4].(map[string]interface{})["data"]
	//logger.Info(utils.Map2JsonString(metadata.AllData))
	for k, v := range userData.(map[string]interface{}) {
		if utils.InArray(k, targets) {
			switch v.(type) {
			case float64, float32:
				msg = fmt.Sprintf("%s = %.0f\n", k, v)
				result.WriteString(msg)
			default:
				msg = fmt.Sprintf("%s = %v\n", k, v)
				result.WriteString(msg)
			}
		}
	}

	// real augly but temporay works for stateless purpose
	metadata := &clients.Metadata{}
	metadata.AllData = alldata
	metadata.AllDataS = &models.AllData{}
	json.Unmarshal([]byte(utils.Map2JsonString(metadata.AllData)), metadata.AllDataS)

	towerInfo, _ := tower.GetCurrentTowerInfo(metadata)
	msg = fmt.Sprintf("年代塔之記 ID = %d\n", towerInfo.Data.TowerID)
	result.WriteString(msg)

	gachaType := []string{"event", "story", "legend"}
	for _, t := range gachaType {
		for page := 1; page <= 3; page++ {
			gachasInfo, _ := web.GetGachaInfo(t, sid, page)
			msg = fmt.Sprintf("轉蛋類型 = %s, Page = %d\n", t, page)
			result.WriteString(msg)
			for _, gachaInfo := range gachasInfo {
				msg = fmt.Sprintf("-> 轉蛋資訊: %+v\n", gachaInfo)
				result.WriteString(msg)
			}
		}
	}

	return result.String()
}
