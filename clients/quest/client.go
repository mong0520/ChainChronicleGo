package quest

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/general"
	"github.com/mong0520/ChainChronicleGo/utils"
)

var API_ENTRY = "quest/entry"
var API_ENTRY_V3 = "quest/v3_entry"

var API_RESULT = "quest/result"
var API_RESULT_V3 = "quest/v3_result"

var ApiMapping = map[int]string{
	2: API_ENTRY,
	3: API_ENTRY_V3,
}

var ApiResultMapping = map[int]string{
	2: API_RESULT,
	3: API_RESULT_V3,
}

type quest struct {
	AutoSell        bool
	AutoBuy         bool
	AutoRaid        bool
	AutoRaidRecover bool
	Count           int
	Type            int
	QuestId         int
	QuestIds        string
	Fid             int
	Pt              int
	Htype           int //unknown
	Lv              int
	Hcid            int
	Version         int
	Res             int
	Bt              int
	Wc              int
	Wn              int
	Time            string
	D               int
	S               int
	Cc              int
	Oc              int
}

type Wvt struct {
	WaveNum int `json:"wave_num"`
	Time    int `json:"time"`
}

type Mission struct {
	Cid []int `json:"cid"`
	Sid []int `json:"sid"`
	Fid []int `json:"fid"`
	Ms  int   `json:"ms"`
	Md  int   `json:"md"`
	Sc  struct {
		Num0 int `json:"0"`
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
		Num4 int `json:"4"`
	} `json:"sc"`
	Es  int `json:"es"`
	At  int `json:"at"`
	He  int `json:"he"`
	Da  int `json:"da"`
	Ba  int `json:"ba"`
	Bu  int `json:"bu"`
	Job struct {
		Num0 int `json:"0"`
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
		Num4 int `json:"4"`
	} `json:"job"`
	Weapon struct {
		Num0  int `json:"0"`
		Num1  int `json:"1"`
		Num2  int `json:"2"`
		Num3  int `json:"3"`
		Num4  int `json:"4"`
		Num5  int `json:"5"`
		Num8  int `json:"8"`
		Num9  int `json:"9"`
		Num10 int `json:"10"`
	} `json:"weapon"`
	Box int `json:"box"`
	Um  struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
	} `json:"um"`
	Fj   int `json:"fj"`
	Fw   int `json:"fw"`
	Fo   int `json:"fo"`
	Mlv  int `json:"mlv"`
	Mbl  int `json:"mbl"`
	Udj  int `json:"udj"`
	Sdmg int `json:"sdmg"`
	Tp   int `json:"tp"`
	Gma  int `json:"gma"`
	Gmr  int `json:"gmr"`
	Gmp  int `json:"gmp"`
	Stp  int `json:"stp"`
	Uh   struct {
		Num3  int `json:"3"`
		Num5  int `json:"5"`
		Num6  int `json:"6"`
		Num9  int `json:"9"`
		Num14 int `json:"14"`
		Num20 int `json:"20"`
	} `json:"uh"`
	Cc    int `json:"cc"`
	BfAtk int `json:"bf_atk"`
	BfHp  int `json:"bf_hp"`
	BfSpd int `json:"bf_spd"`
}

// Beloved
type BL struct {
	SrcCid   int  `json:"src_cid"`
	Mana     int  `json:"mana"`
	UseSkill bool `json:"use_skill"`
}

func (q *quest) StartQeust(u *clients.Metadata) (resp map[string]interface{}, res int) {
	requestUrl := GetEndpoint(q.Version)
	//fmt.Printf("%+v", q)
	postBody := q.getPostBody()
	resp, _ = utils.PostV2(requestUrl, "", postBody, u.Sid)
	res = int(resp["res"].(float64))
	return resp, res
}

func (q *quest) EndQeustV2(u *clients.Metadata) (resp map[string]interface{}, res int) {
	requestUrl := GetResultEndpoint(q.Version)
	// nowHex := fmt.Sprintf("%x", int(time.Now().UnixNano()))
	wvt := []Wvt{
		{WaveNum: 1, Time: 1333}, {WaveNum: 2, Time: 1333}, {WaveNum: 3, Time: 1333},
	}
	// fmt.Println(wvt)

	mission := Mission{
		Cid:   []int{1213, 1216, 224, 5275, 1254, 9216},
		Sid:   []int{0, 0, 0, 0, 0, 0},
		Fid:   []int{9222},
		Ms:    0,
		Md:    18934,
		Es:    0,
		At:    5,
		He:    1,
		Da:    0,
		Ba:    0,
		Bu:    1,
		Box:   1,
		Fj:    2,
		Fw:    4,
		Fo:    0,
		Mlv:   80,
		Mbl:   152,
		Udj:   0,
		Sdmg:  140000,
		Tp:    8,
		Gma:   13,
		Gmr:   0,
		Gmp:   0,
		Stp:   0,
		Cc:    1,
		BfAtk: 0,
		BfHp:  0,
		BfSpd: 0,
	}
	mission.Sc.Num0 = 2
	mission.Sc.Num1 = 1
	mission.Sc.Num2 = 1
	mission.Sc.Num3 = 1
	mission.Sc.Num4 = 1

	mission.Job.Num0 = 0
	mission.Job.Num1 = 1
	mission.Job.Num2 = 4
	mission.Job.Num3 = 2
	mission.Job.Num4 = 0

	mission.Weapon.Num0 = 0
	mission.Weapon.Num1 = 0
	mission.Weapon.Num2 = 0
	mission.Weapon.Num3 = 0
	mission.Weapon.Num4 = 4
	mission.Weapon.Num5 = 2
	mission.Weapon.Num8 = 0
	mission.Weapon.Num9 = 0
	mission.Weapon.Num10 = 1

	mission.Um.Num1 = 5
	mission.Um.Num2 = 1
	mission.Um.Num3 = 0

	mission.Uh.Num3 = 1
	mission.Uh.Num5 = 1
	mission.Uh.Num6 = 1
	mission.Uh.Num9 = 2
	mission.Uh.Num14 = 1
	mission.Uh.Num20 = 1
	// fmt.Println(mission)

	bl := []BL{
		{SrcCid: 1017, Mana: 6, UseSkill: true},
		{SrcCid: 5001, Mana: 6, UseSkill: true},
		{SrcCid: 1002, Mana: 6, UseSkill: true},
		{SrcCid: 7012, Mana: 6, UseSkill: true},
		{SrcCid: 1254, Mana: 6, UseSkill: true},
		{SrcCid: 9012, Mana: 6, UseSkill: true},
	}
	// fmt.Println(bl)

	blf := []BL{
		{SrcCid: 9222, Mana: 5, UseSkill: true},
	}
	// fmt.Println(blf)

	wvtValue := utils.Struct2JsonString(wvt)
	// fmt.Println(wvtValue)

	missionValue := utils.Struct2JsonString(mission)
	// fmt.Println(missionValue)

	blValue := utils.Struct2JsonString(bl)
	// fmt.Println(blValue)

	blfValue := utils.Struct2JsonString(blf)
	// fmt.Println(blfValue)

	queryStringPart1 := fmt.Sprintf("wvt=%s&mission=%s&bl=%s&blf=%s",
		url.QueryEscape(wvtValue),
		url.QueryEscape(missionValue),
		url.QueryEscape(blValue),
		url.QueryEscape(blfValue))

	body := q.getEndPostBody()
	tempNow := int(time.Now().UnixNano())
	// nowShort := strconv.Itoa(int(time.Now().Unix()))
	nowHex := fmt.Sprintf("%x", tempNow)
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
			fmt.Println("Unhandled type")
		}
	}
	data.Add("cnt", nowHex)
	queryStringPart2 := fmt.Sprintf("nature=%s&%s", url.QueryEscape(data.Encode()), queryStringPart1)

	// fmt.Println("===============")
	// fmt.Println(queryStringPart1)
	finalPostData := fmt.Sprintf("%s&%s", queryStringPart1, url.QueryEscape(queryStringPart2))
	// fmt.Println(finalPostData)

	resp, _ = utils.PostV2(requestUrl, finalPostData, q.getEndPostBody(), u.Sid)
	res = int(resp["res"].(float64))
	return resp, res
}

func (q *quest) EndQeust(u *clients.Metadata) (resp map[string]interface{}, res int) {
	requestUrl := GetResultEndpoint(q.Version)
	postBody := q.getEndPostBody()
	nowHex := fmt.Sprintf("%x", int(time.Now().UnixNano()))
	payloadString := "ch=&eh=&ec=&mission=%7b%22cid%22%3a%5b6201%2c6003%5d%2c%22sid%22%3a%5b0%2c0%5d%2c%22fid%22%3a%5b3001%5d%2c%22ms%22%3a0%2c%22md%22%3a15072%2c%22sc%22%3a%7b%220%22%3a2%2c%221%22%3a0%2c%222%22%3a0%2c%223%22%3a0%2c%224%22%3a0%7d%2c%22es%22%3a0%2c%22at%22%3a1%2c%22he%22%3a0%2c%22da%22%3a0%2c%22ba%22%3a0%2c%22bu%22%3a1%2c%22job%22%3a%7b%220%22%3a3%2c%221%22%3a0%2c%222%22%3a0%2c%223%22%3a0%2c%224%22%3a0%7d%2c%22weapon%22%3a%7b%220%22%3a2%2c%221%22%3a1%2c%222%22%3a0%2c%223%22%3a0%2c%224%22%3a0%2c%225%22%3a0%2c%228%22%3a0%2c%229%22%3a0%2c%2210%22%3a0%7d%2c%22box%22%3a1%2c%22um%22%3a%7b%221%22%3a1%2c%222%22%3a1%2c%223%22%3a0%7d%2c%22fj%22%3a0%2c%22fw%22%3a0%2c%22fo%22%3a0%2c%22mlv%22%3a80%2c%22mbl%22%3a150%2c%22udj%22%3a0%2c%22sdmg%22%3a35133%2c%22tp%22%3a12%2c%22gma%22%3a5%2c%22gmr%22%3a5%2c%22gmp%22%3a0%2c%22stp%22%3a0%2c%22uh%22%3a%7b%2210%22%3a2%2c%227%22%3a1%7d%2c%22cc%22%3a1%2c%22bf_atk%22%3a0%2c%22bf_hp%22%3a0%2c%22bf_spd%22%3a0%7d&bl=%5b%7b%22src_cid%22%3a59009%2c%22mana%22%3a2%2c%22use_skill%22%3atrue%7d%2c%7b%22src_cid%22%3a6003%2c%22mana%22%3a1%2c%22use_skill%22%3atrue%7d%2c%7b%7d%2c%7b%7d%5d&blf=%5b%7b%22src_cid%22%3a3001%2c%22mana%22%3a0%2c%22use_skill%22%3afalse%7d%5d&nature=bl%3d%255b%257b%2522src_cid%2522%253a59009%252c%2522mana%2522%253a2%252c%2522use_skill%2522%253atrue%257d%252c%257b%2522src_cid%2522%253a6003%252c%2522mana%2522%253a1%252c%2522use_skill%2522%253atrue%257d%252c%257b%257d%252c%257b%257d%255d%26blf%3d%255b%257b%2522src_cid%2522%253a3001%252c%2522mana%2522%253a0%252c%2522use_skill%2522%253afalse%257d%255d%26bt%3d5459%26cc%3d1%26ch%3d%26cnt%3dTBD%26d%3d1%26ec%3d%26eh%3d%26mission%3d%257b%2522cid%2522%253a%255b6201%252c6003%255d%252c%2522sid%2522%253a%255b0%252c0%255d%252c%2522fid%2522%253a%255b3001%255d%252c%2522ms%2522%253a0%252c%2522md%2522%253a15072%252c%2522sc%2522%253a%257b%25220%2522%253a2%252c%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a0%252c%25224%2522%253a0%257d%252c%2522es%2522%253a0%252c%2522at%2522%253a1%252c%2522he%2522%253a0%252c%2522da%2522%253a0%252c%2522ba%2522%253a0%252c%2522bu%2522%253a1%252c%2522job%2522%253a%257b%25220%2522%253a3%252c%25221%2522%253a0%252c%25222%2522%253a0%252c%25223%2522%253a0%252c%25224%2522%253a0%257d%252c%2522weapon%2522%253a%257b%25220%2522%253a2%252c%25221%2522%253a1%252c%25222%2522%253a0%252c%25223%2522%253a0%252c%25224%2522%253a0%252c%25225%2522%253a0%252c%25228%2522%253a0%252c%25229%2522%253a0%252c%252210%2522%253a0%257d%252c%2522box%2522%253a1%252c%2522um%2522%253a%257b%25221%2522%253a1%252c%25222%2522%253a1%252c%25223%2522%253a0%257d%252c%2522fj%2522%253a0%252c%2522fw%2522%253a0%252c%2522fo%2522%253a0%252c%2522mlv%2522%253a80%252c%2522mbl%2522%253a150%252c%2522udj%2522%253a0%252c%2522sdmg%2522%253a35133%252c%2522tp%2522%253a12%252c%2522gma%2522%253a5%252c%2522gmr%2522%253a5%252c%2522gmp%2522%253a0%252c%2522stp%2522%253a0%252c%2522uh%2522%253a%257b%252210%2522%253a2%252c%25227%2522%253a1%257d%252c%2522cc%2522%253a1%252c%2522bf_atk%2522%253a0%252c%2522bf_hp%2522%253a0%252c%2522bf_spd%2522%253a0%257d%26qid%3d220103%26res%3d1%26s%3d0%26time%3d2.63%26wc%3d4%26wn%3d4"
	payloadString = strings.Replace(payloadString, "TBD", nowHex, 1)
	resp, _ = utils.PostV2(requestUrl, payloadString, postBody, u.Sid)
	res = int(resp["res"].(float64))
	return resp, res
}

func (q *quest) GetTreasure(u *clients.Metadata) (resp map[string]interface{}, res int) {
	api := "quest/treasure"
	sid := u.Sid
	param := map[string]interface{}{
		"type": q.Type,
		"qid":  q.QuestId,
	}
	return general.GeneralAction(api, sid, param)
}

func GetEndpoint(version int) (endpoint string) {
	endpoint = fmt.Sprintf("%s/%s", clients.HOST, ApiMapping[version])
	return endpoint
}

func GetResultEndpoint(version int) (endpoint string) {
	endpoint = fmt.Sprintf("%s/%s", clients.HOST, ApiResultMapping[version])
	return endpoint
}

func (q *quest) getEndPostBody() (body map[string]interface{}) {
	body = map[string]interface{}{
		"qid":  q.QuestId,
		"fid":  q.Fid,
		"res":  q.Res,
		"bt":   q.Bt,
		"time": q.Time,
		"d":    q.D,
		"s":    q.S,
		"cc":   q.Cc,
		"wc":   q.Wc,
		"wn":   q.Wn,
		// "mtr":  8,
		// "mtn":  8,
	}
	if q.Version == 3 {
		body["lv"] = q.Lv
	}
	return body
}

func (q *quest) getPostBody() (body map[string]interface{}) {
	body = map[string]interface{}{
		"type":  q.Type,
		"qid":   q.QuestId,
		"fid":   q.Fid,
		"pt":    q.Pt,
		"htype": q.Htype,
		"oc":    q.Oc,
	}
	if q.Version == 3 {
		body["lv"] = q.Lv
		body["hcid"] = q.Hcid
	}
	return body
}

func NewQuest() (q *quest) {
	q = &quest{
		Hcid:    0,
		Htype:   0,
		Lv:      0,
		Version: 2,
		Fid:     229741,
		Time:    "10.0",
		Res:     1,
		Bt:      10,
		Wc:      4,
		Wn:      1,
		Cc:      1,
		D:       1,
		S:       1,
		Oc:      1,
	}
	return q
}
