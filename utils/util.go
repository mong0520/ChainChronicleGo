package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/oleiade/reflections"
	"github.com/robfig/config"
)

func PostV2(requestUrl string, rawPayload string, body map[string]interface{}, sid string) (respMap map[string]interface{}, err error) {
	//fmt.Println(requestUrl)
	//fmt.Println(body)
	logger := GetLogger(nil)
	tempNow := int(time.Now().UnixNano())
	nowShort := strconv.Itoa(int(time.Now().Unix()))
	nowHex := fmt.Sprintf("%x", tempNow)

	// Build up post data
	data := url.Values{}
	for field := range body {
		vType := reflect.TypeOf(body[field].(interface{}))
		switch vType.String() {
		case "int":
			value := strconv.Itoa(body[field].(int))
			data.Add(field, value)
		case "string":
			value := body[field].(string)
			data.Add(field, value)
		default:
			logger.Println("Unhandled type")
		}

	}
	data.Add("cnt", nowHex)
	postData := fmt.Sprintf("nature=%s", url.QueryEscape(data.Encode()))
	if rawPayload != "" {
		postData = rawPayload
	}
	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(postData))

	// Build up query string
	queries := req.URL.Query()
	queries.Set("timestamp", nowShort)
	queries.Set("cnt", nowHex)
	for field := range body {
		vType := reflect.TypeOf(body[field].(interface{}))
		switch vType.String() {
		case "int":
			value := strconv.Itoa(body[field].(int))
			queries.Set(field, value)
		case "string":
			value := body[field].(string)
			queries.Set(field, value)
		default:
			logger.Println("Unhandled type")
		}
	}
	req.URL.RawQuery = queries.Encode()

	if err != nil {
		logger.Printf("http.NewRequest() error: %v\n", err)
		return nil, err
	}
	c := &http.Client{}
	cookie := &http.Cookie{
		Name:  "sid",
		Value: sid,
	}
	req.AddCookie(cookie)
	// logger.Printf("[POST] %s\n", req.URL.String())
	// logger.Printf("Post Data = %s\n", postData)
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

	// Unzip response body
	respMap, err = DecodeResponse(ret)
	if err != nil {
		fmt.Printf("Decode response error, %s\n", err)
		fmt.Println("Raw response = ", string(ret))
		return nil, err
	}
	//fmt.Println(Map2JsonString(respMap))
	//logger.Printf("res = %d\n", int(respMap["res"].(float64)))
	return respMap, nil
}

func DecodeResponse(raw []byte) (result map[string]interface{}, err error) {
	rdata := bytes.NewReader(raw)
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return nil, err
	}
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(s, &result)
	return result, err
}

func DecodeResponseV2(raw []byte) (b []byte, err error) {
	rdata := bytes.NewReader(raw)
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return nil, err
	}
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return s, err
}

func Map2JsonString(m map[string]interface{}) (ret string) {
	retBytes, _ := json.Marshal(m)
	ret = string(retBytes)
	return ret
}

func Map2Struct(m map[string]interface{}, v interface{}) (ret error) {
	if d, err := json.Marshal(m); err != nil {
		return err
	} else {
		if err := json.Unmarshal(d, v); err != nil {
			return err
		}
	}
	return nil
}

func GetLogger(f *os.File) (logger *log.Logger) {
	if f != nil {
		logger = log.New(io.MultiWriter(os.Stdout, f), "", log.LstdFlags|log.Lshortfile)
	} else {
		logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	}

	return logger
}

func DumpConfig(c *config.Config) {
	fmt.Println("=== Dump configs ===")
	for _, section := range c.Sections() {
		fmt.Printf("[%s]\n", section)
		//    //sectiondata, _ := file.GetSection(section)
		//    //for _, item := range sectiondata.Keys(){
		//    //   fmt.Printf("%s = %s\n", item.Name(), item.String())
		//    //}
		//
	}
}

func ParseConfig2Struct(conf *config.Config, section string, data interface{}) {
	fields, _ := conf.SectionOptions(section)
	for _, field := range fields {
		strValue, _ := conf.String(section, field)
		//fmt.Println(field, strValue)
		if intValue, err := strconv.Atoi(strValue); err == nil {
			reflections.SetField(data, field, intValue)
			continue
		}
		if boolValue, err := strconv.ParseBool(strValue); err == nil {
			//fmt.Println("set bool value ", field, strValue)
			reflections.SetField(data, field, boolValue)
			continue
		}
		reflections.SetField(data, field, strValue)
	}
}

func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false
	//index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				//index = i
				exists = true
				return
			}
		}
	}
	return
}
