package web

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type GachaInfo struct {
	GachaTitle string
	GachaCount string
	GachaID    string
	ItemName   string
}

func GetGachaInfo(t string, sid string) ([]GachaInfo, error) {

	gachasInfo := []GachaInfo{}
	url := fmt.Sprintf("http://v3810.cc.mobimon.com.tw/web/gacha?type=%s&gacha_id=1", t)
	cookieString := fmt.Sprintf("Cookie: devicewidth=568; framewidth=568; sid=%s", sid)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cookie", cookieString)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "v3810.cc.mobimon.com.tw")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return gachasInfo, err
	}

	defer res.Body.Close()
	var reader io.ReadCloser

	reader, err = gzip.NewReader(res.Body)
	defer reader.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	// newStr := buf.String()
	// fmt.Println(newStr)

	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".btn_area").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		content, _ := s.Html()
		// fmt.Println(content)
		for {
			idx := strings.Index(content, "execGacha(")
			if idx == -1 {
				break
			}
			gachaInfo := GachaInfo{}
			gachaInfo.GachaTitle = t
			idxEnd := strings.Index(content[idx:len(content)], "}") + idx + 1
			gachaRawData := content[idx+len("execGacha(") : idxEnd]
			tokens := strings.Split(gachaRawData, ",")
			gachaInfo.GachaID = strings.Trim(tokens[0], " ")
			gachaInfo.GachaCount = strings.Trim(tokens[1], " ")
			if strings.Index(gachaRawData, "精靈石") != -1 {
				gachaInfo.ItemName = "精靈石"
			} else if strings.Index(gachaRawData, "轉蛋幣") != -1 {
				gachaInfo.ItemName = "轉蛋幣"
			} else {
				gachaInfo.ItemName = "Undefined"
			}

			gachasInfo = append(gachasInfo, gachaInfo)

			// next content
			content = content[idxEnd:len(content)]
		}

	})

	return gachasInfo, nil
}
