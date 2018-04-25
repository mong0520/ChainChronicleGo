package main

import (
	"log"
	"strings"

	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/item"
	"github.com/mong0520/ChainChronicleGo/clients/quest"
	"github.com/mong0520/ChainChronicleGo/clients/session"
	"github.com/mong0520/ChainChronicleGo/clients/user"
	"github.com/mong0520/ChainChronicleGo/models"
	"github.com/mong0520/ChainChronicleGo/utils"

	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/icza/dyno"
	"github.com/jessevdk/go-flags"
	"github.com/mong0520/ChainChronicleGo/clients/card"
	"github.com/mong0520/ChainChronicleGo/clients/explorer"
	"github.com/mong0520/ChainChronicleGo/clients/gacha"
	"github.com/mong0520/ChainChronicleGo/clients/general"
	"github.com/mong0520/ChainChronicleGo/clients/present"
	"github.com/mong0520/ChainChronicleGo/clients/raid"
	"github.com/mong0520/ChainChronicleGo/clients/teacher"
	"github.com/mong0520/ChainChronicleGo/clients/tutorial"
	"github.com/mong0520/ChainChronicleGo/clients/weapon"
	"github.com/mong0520/ChainChronicleGo/controllers"
	"github.com/robfig/config"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"path"
	"reflect"
	"strconv"
	"time"
)

type Options struct {
	ConfigPath string `short:"c" long:"config" description:"Config path" required:"true"`
	Action     string `short:"a" long:"action" description:"Action to run" required:"false"`
	Repeat     int    `short:"r" long:"repeat" description:"Repeat action for r times" required:"false"`
	Timeout    int    `short:"t" long:"timeout" description:"Timeout in seconds between repeat" required:"false"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)
var metadata = &clients.Metadata{}
var logger *log.Logger
var LogROOT = "logs"
var actionMapping = map[string]interface{}{
	"QUEST":          doQuest,
	"STATUS":         doStatus,
	"PASSWORD":       doPassword,
	"TAKEOVER":       doTakeOver,
	"BUY":            doBuy,
	"GACHA":          doGacha,
	"TUTORIAL":       doTutorial,
	"DRAMA":          doDrama,
	"DEBUG":          doDebug,
	"RESET_DISCIPLE": doResetDisciple,
	"CHARS":          doShowChars,
	"ALLDATA":        doShowAllData,
	"COMPOSE":        doCompose,
	"TEACHER":        doTeacher,
	"DISCIPLE":       doDisciple,
	"UPDATE":         doUpdateDB,
	"EXPLORER":       doExplorer,
}

func doAction(sectionName string) {
	for action, actionFunction := range actionMapping {
		//logger.Println(action, actionFunction)
		if strings.HasPrefix(sectionName, action) {
			logger.Printf("### Current Flow = [%s] ###", sectionName)
			actionFunction.(func(*clients.Metadata, string))(metadata, sectionName)
		}
	}
}

func initLogFile() (logFile *os.File, err error) {
	logFileName := path.Base(options.ConfigPath)
	logFilePath := path.Join("logs", logFileName)
	if _, err := os.Stat(LogROOT); os.IsNotExist(err) {
		os.Mkdir(LogROOT, 0755)
	}
	return os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func start() {
	logFile, err := initLogFile()
	defer logFile.Close()

	if err != nil {
		logger.Fatalln("open file error !")
	}
	logger = utils.GetLogger(logFile)

	config, err := config.ReadDefault(options.ConfigPath)
	if err != nil {
		logger.Fatalln("Unable to read config, ", err)
		return
	}

	metadata.Config = config

	if db, err := mgo.Dial("localhost:27017"); err != nil {
		logger.Fatalln("Unable to connect DB", err)
	} else {
		metadata.DB = db
	}

	//utils.DumpConfig(metadata.Config)
	uid, _ := metadata.Config.String("GENERAL", "Uid")
	metadata.Uid = uid

	//logger.Println(uid)
	token, _ := metadata.Config.String("GENERAL", "Token")
	metadata.Token = token
	if options.Action == "" {
		flowString, _ := metadata.Config.String("GENERAL", "Flow")
		flowString = strings.Replace(flowString, " ", "", -1)
		metadata.Flow = strings.Split(flowString, ",")
	} else {
		flowString := options.Action
		flowString = strings.Replace(flowString, " ", "", -1)
		metadata.Flow = strings.Split(flowString, ",")
	}

	sid, err := session.Login(uid, token, false)
	if err != nil {
		logger.Printf("%s\n", err)
		return
	}
	alldata, _ := user.GetAllData(sid)
	metadata.AllData = alldata
	metadata.Sid = sid
	//metadata.AllDataS = &models.AllData{}
	//
	//err = utils.Map2Struct(alldata, metadata.AllDataS)
	//if err!= nil {
	//	log.Println(err)
	//	os.Exit(-1)
	//}else{
	//	log.Println(metadata.AllDataS)
	//}
	//dumpUser(metadata)
	flowLoop, _ := metadata.Config.Int("GENERAL", "FlowLoop")
	sleepDuration, _ := metadata.Config.Int("GENERAL", "FlowLoopDelay")

	if options.Repeat > 0 {
		flowLoop = options.Repeat
	}
	if options.Timeout >= 0 {
		sleepDuration = options.Timeout
	}

	for i := 1; i <= flowLoop; i++ {
		logger.Printf("Start #%d Flow\n", i)
		for _, flow := range metadata.Flow {
			logger.Printf("Current action = [%s]\n", flow)
			doAction(strings.ToUpper(flow))
		}
		if sleepDuration > 0 {
			logger.Println("Waiting", sleepDuration, "seconds")
			time.Sleep(time.Second * time.Duration(sleepDuration))
		}
	}
}

func doDrama(metadata *clients.Metadata, section string) {
	questInfo := quest.NewQuest()
	//questList, _ := dyno.GetSlice(metadata.AllData, "body", 29, "data")
	//logger.Println(questList)
	logger.Printf("開始通過主線任務...\n")
	maxRetryCount := 10
	currentRetryCount := 0
	flag := 0
	lastQid := 331043
	counter := 0
	dramaLevel := 1
	gradudateThreshold := 38
	hcid := 9420
	lvThreshold := 50

	for {
		//logger.Println(v, reflect.TypeOf(v))
		//logger.Println(dyno.Get(metadata.AllData, "body", 29, "data", flag, "id"))
		if qType, err := dyno.GetFloat64(metadata.AllData, "body", 29, "data", flag, "type"); err != nil {
			logger.Println(qType, err)
		} else {
			questInfo.Type = int(qType)
		}
		if qId, err := dyno.GetFloat64(metadata.AllData, "body", 29, "data", flag, "id"); err != nil {
			logger.Println(qId, err)
		} else {
			questInfo.QuestId = int(qId)
		}

		counter += 1
		if counter >= gradudateThreshold {
			// check if current LV >= 50
			break
		}
		questInfo.Fid = 1965350
		questInfo.Lv = dramaLevel
		questInfo.Hcid = hcid
		questInfo.Pt = 0
		questInfo.Version = 3
		errSet := mapset.NewSet()
		errSet.Clear()
		//logger.Printf("%+v\n", questInfo)
		resp, err := questInfo.StartQeust(metadata)
		errSet.Add(err)
		switch err {
		case 0:
			_, err = questInfo.EndQeust(metadata)
			logger.Printf("%d/%d 完成關卡\n", counter, gradudateThreshold)
			errSet.Add(err)
		case 103:
			logger.Printf("體力不足\n")
			if _, err := user.RecoveryAp(1, 1, metadata.Sid); err != 0 {
				logger.Println("無法恢復體力")
				panic(err)
			}
		default:
			logger.Println("未知的錯誤")
			errSet.Add(err)
			logger.Println(utils.Map2JsonString(resp))
			if resp, err := questInfo.GetTreasure(metadata); err != 0 {
				logger.Println(resp)
			}
		}
		if errSet.Contains(0) == false {
			logger.Printf("Unknown error in drama: %s", errSet)
			currentRetryCount++
			if currentRetryCount >= maxRetryCount {
				uid, _ := metadata.Config.String("GENERAL", "Uid")
				logger.Printf("UID %s is is uable to complete Drama", uid)
			} else {
				logger.Println("Retry...")
				continue
			}
		} else {
			currentRetryCount = 0
			if questInfo.QuestId == lastQid {
				continue
			} else {
				if flag >= 4 {
					dramaLevel = 2
				}
				flag++
			}
		}
	}
	alldata, _ := user.GetAllData(metadata.Sid)
	metadata.AllData = alldata
	if currentLv, err := dyno.GetFloat64(metadata.AllData, "body", 4, "data", "lv"); err != nil {
		logger.Println(err)
	} else {
		logger.Printf("Current LV %0.f", currentLv)
		if int(currentLv) >= lvThreshold {
			teacher.IS_GRADUATED = true
		}
	}
}

func doTutorial(metadata *clients.Metadata, section string) {
	tutorialInfo := []map[string]int{
		{"tid": 0, "qid": -1},
		{"tid": 1, "qid": -1},
		{"tid": 2, "qid": -1},
		{"tid": 3, "qid": 210001},
		{"tid": 4, "qid": 210001},
		{"tid": 5, "qid": -1},
		{"tid": 6, "qid": 210002},
		{"tid": 7, "qid": -1},
		{"tid": 8, "qid": 210101},
		{"tid": 9, "qid": -1},
		{"tid": 10, "qid": 210101},
		{"tid": 11, "qid": -1},
		{"tid": 12, "qid": -1},
		{"tid": 13, "qid": 210102},
		{"tid": 14, "qid": -1},
		{"tid": 15, "qid": 210102},
		{"tid": 16, "qid": -1},
		{"tid": 17, "qid": 215000},
		{"tid": 18, "qid": 215000},
		{"tid": 19, "qid": -1},
		{"tid": 20, "qid": -1},
	}
	newUid := fmt.Sprintf("ANDO%s", uuid.Must(uuid.NewV4()).String())
	logger.Printf("New UUID = %s", newUid)
	// set tor proxy
	sid, err := session.Login(newUid, "", false)
	if err != nil {
		log.Printf("無法建立新帳號，嘗試使用 TOR...\n")
		if sid, err = session.Login(newUid, "", true); err != nil {
			log.Printf("建立帳號失敗 %s\n", err)
			os.Exit(1)
		}
	}
	metadata.Uid = newUid
	//logger.Println(uid)
	token, _ := metadata.Config.String("GENERAL", "Token")
	metadata.Token = token
	metadata.Sid = sid
	resp, _ := user.GetAllData(sid)
	openId, _ := dyno.Get(resp, "body", 4, "data", "uid")
	logger.Printf("新帳號創立成功, UID = %s, OpenID = %.0f\n", newUid, openId)
	//
	for _, t := range tutorialInfo {
		if t["qid"] != -1 {
			param := map[string]interface{}{
				"pt":  0,
				"tid": t["tid"],
			}
			tutorial.Tutorial(sid, true, param)
			questInfo := quest.NewQuest()
			questInfo.QuestId = t["qid"]
			questInfo.Fid = 1965350
			questInfo.Pt = 0
			if resp, err := questInfo.EndQeust(metadata); err != 0 {
				logger.Println(utils.Map2JsonString(resp), err)
				break
			}
		} else {
			if t["tid"] == 1 {
				param := map[string]interface{}{
					"name": "Allen",
					"hero": "Allen",
					"tid":  t["tid"],
				}
				if resp, err := tutorial.Tutorial(sid, false, param); err != 0 {
					logger.Println(utils.Map2JsonString(resp), err)
					break
				}
			} else {
				param := map[string]interface{}{
					"tid": t["tid"],
				}
				if resp, err := tutorial.Tutorial(sid, false, param); err != 0 {
					logger.Println(utils.Map2JsonString(resp), err)
					break
				}
			}
		}
	}
	logger.Printf("新帳號完成新手教學, UID = %s, OpenID = %.0f\n", newUid, openId)
	getPresents(metadata, nil)
}

func doGacha(metadata *clients.Metadata, section string) {
	gachaInfo := gacha.NewGachaInfo()
	utils.ParseConfig2Struct(metadata.Config, section, gachaInfo)
	logger.Println("開始轉蛋")
	if resp, ret := gachaInfo.Gacha(metadata); ret == 0 {
		gachaResult := processGachaResult(resp)
		for _, card := range gachaResult["char"].([]models.GachaResultChara) {
			myCard := models.Charainfo{}
			query := bson.M{"cid": card.ID}
			controllers.GeneralQuery(metadata.DB, "charainfo", query, &myCard)
			//logger.Printf("得到 %d星卡: %s-%s", myCard.Rarity, myCard.Title, myCard.Name)
			if gachaInfo.AutoSell && myCard.Rarity <= gachaInfo.AutoSellRarityThreshold {
				logger.Println("賣出卡片...")
				doSellItem(metadata, card.Idx, "")
			}
		}

	} else {
		logger.Println(utils.Map2JsonString(resp))
	}
}

func doSellItem(metadata *clients.Metadata, cid int, section string) {
	if ret, err := card.Sell(metadata, cid); err != 0 {
		logger.Println("\t-> 賣出卡片失敗", utils.Map2JsonString(ret))
	} else {
		logger.Println("\t-> 賣出卡片成功")
	}
}

func doExplorer(metadata *clients.Metadata, section string) {
	getPresents(metadata, []string{"card"})
	setCharaData()
	metadata.ExplorerExcludeCids = []int{2007}
	explorerList, _ := explorer.GetExplorerList(metadata.Sid)
	pickup, _ := dyno.GetSlice(explorerList, "pickup")
	//log.Printf("%s\n", utils.Map2JsonString(explorerList))
	//log.Printf("%v\n", pickup)
	pickupList := []explorer.Pickup{}

	for _, p := range pickup {
		pickupItem := &explorer.Pickup{}
		utils.Map2Struct(p.(map[string]interface{}), pickupItem)
		pickupList = append(pickupList, *pickupItem)
	}
	//for _, pickupIteam := range pickupList{
	//	log.Printf("%+v\n", pickupIteam)
	//}
	useStone, _ := metadata.Config.Bool(section, "StoneFinish")
	explorerAreaStr, _ := metadata.Config.String(section, "area")
	explorerAreas := strings.Split(explorerAreaStr, ",")

	for i, e := range explorerAreas {
		area, _ := strconv.Atoi(e)
		resp, err := explorer.GetExplorerResult(metadata.Sid, i+1)
		switch err {
		case 0, 2308:
			logger.Println("可以開始探索")

		case 2302:
			logger.Println("探索尚未結束")
			logger.Printf("Use stone to finish? %t\n", useStone)
			if useStone {
				explorer.FinishExplorer(metadata.Sid, i+1)
				//logger.Println(utils.Map2JsonString(resp))

				resp, _ = explorer.GetExplorerResult(metadata.Sid, i+1)
				//logger.Println(utils.Map2JsonString(resp))

				rewards, _ := dyno.GetSlice(resp, "explorer_reward")
				for _, reward := range rewards {
					//logger.Println(reward)
					rewardItem := reward.(map[string]interface{})["item_id"].(float64)
					rewardType := reward.(map[string]interface{})["item_type"].(string)
					//logger.Println("Reward Type ", rewardType)
					if rewardType == "card" {
						tmpCardInfo := &models.Charainfo{} // for mongodb result
						cid := int(rewardItem)
						//logger.Println("cid =", cid)
						query := bson.M{"cid": cid}
						if err := controllers.GeneralQuery(metadata.DB, "charainfo", query, &tmpCardInfo); err != nil {
							logger.Println("得到角色", rewardItem, err)
						} else {
							logger.Println("得到", tmpCardInfo.Rarity, "星角色", tmpCardInfo.Name)
							if tmpCardInfo.Cid == 6248 {
								os.Exit(0)
							}
						}
					} else {
						//logger.Println("得到ID", rewardItem)
					}
				}
			}
		case 1:
			logger.Println("已被登出")
		default:
			logger.Println("未知的結果")
			logger.Println(resp)
		}
		for _, pickupItem := range pickupList {
			if pickupItem.LocationID == area {
				//logger.Printf("Start to find best card to explorer area %d\n", pickupItem.LocationID)
				result := findBestCardToExplorer(&pickupItem)
				param := map[string]int{
					"explorer_idx": i + 1,
					"location_id":  area,
					"card_idx":     result["idx"],
					"pickup":       1,
					"interval":     2,
				}
				resp, err := explorer.StartExplorer(metadata.Sid, param)
				switch err {
				case 0:
					break
				case 2311:
					param["pickup"] = 0
					explorer.StartExplorer(metadata.Sid, param)
				default:
					logger.Printf("%s\n", utils.Map2JsonString(resp))
				}
			}
		}
	}
}

func setCharaData() {
	chars, _ := dyno.GetSlice(metadata.AllData, "body", 6, "data")
	charaInfo := []models.Charainfo{} // for mongodb result
	charaData := []models.CharaData{} // for alldata structure
	for _, c := range chars {
		tmpCardInfo := &models.Charainfo{} // for mongodb result
		tmpCharData := &models.CharaData{} // for alldata structure
		utils.Map2Struct(c.(map[string]interface{}), tmpCharData)
		if tmpCharData.Type != 0 {
			continue
		}
		query := bson.M{"cid": tmpCharData.ID}
		if err := controllers.GeneralQuery(metadata.DB, "charainfo", query, &tmpCardInfo); err != nil {
			logger.Println(tmpCharData.ID, err)
		} else {
			charaInfo = append(charaInfo, *tmpCardInfo)
			charaData = append(charaData, *tmpCharData)
		}
	}
	if metadata.CharInfo == nil {
		metadata.CharInfo = charaInfo
	}
	if metadata.CharData == nil {
		metadata.CharData = charaData
	}
}

func findBestCardToExplorer(pickupItem *explorer.Pickup) (result map[string]int) {
	result = map[string]int{
		"cid": 0,
		"idx": 0,
	}
	for idx, charInfo := range metadata.CharInfo {
		charData := metadata.CharData[idx]
		if charInfo.Rarity >= 5 {
			// 不使用五星卡探索
			continue
		} else if ((pickupItem.Home == charInfo.Home) || (pickupItem.Jobtype == charInfo.Jobtype)) && !utils.InArray(charInfo.Cid, metadata.ExplorerExcludeCids) {
			// 適合的
			logger.Printf("Pick %s to explorer, cid = %d, idx = %d, rank = %d\n", charInfo.Name, charInfo.Cid, charData.Idx, charInfo.Rarity)
			metadata.ExplorerExcludeCids = append(metadata.ExplorerExcludeCids, charInfo.Cid)
			result["cid"] = charInfo.Cid
			result["idx"] = charData.Idx
			break
		} else {
			// 找不到適合的
			result["cid"] = charInfo.Cid
			result["idx"] = charData.Idx
		}
	}
	return result
}

func processGachaResult(resp map[string]interface{}) (gachaResult map[string]interface{}) {
	gachaData, _ := dyno.GetSlice(resp, "body")
	//logger.Println(utils.Map2JsonString(resp))
	gachaResult = map[string]interface{}{
		"char":   []models.GachaResultChara{},
		"item":   []models.GachaResultItem{},
		"weapon": []models.GachaResultWeapon{},
	}

	gachaResult["char"] = []models.GachaResultChara{}
	gachaResult["item"] = []models.GachaResultItem{}
	gachaResult["weapon"] = []models.GachaResultWeapon{}

	charList := make([]models.GachaResultChara, 0)
	itemList := make([]models.GachaResultItem, 0)
	weaponList := make([]models.GachaResultWeapon, 0)

	for i, data := range gachaData {
		if i == 0 {
			continue
		}
		dataType, _ := dyno.GetFloat64(data, "type")
		switch dataType {
		case 15:
			logger.Println(i, "Type 15", data)
		case 1:
			//logger.Println(i, "得到角色")
			list := data.(map[string]interface{})["data"].([]interface{})
			for _, item := range list {
				tmpItem := &models.GachaResultChara{}
				tmpDBItem := &models.Charainfo{}
				if err := utils.Map2Struct(item.(map[string]interface{}), tmpItem); err != nil {
					logger.Println("Unable to convert to struct", err)
				} else {
					query := bson.M{"cid": tmpItem.ID}
					if err := controllers.GeneralQuery(metadata.DB, "charainfo", query, tmpDBItem); err != nil {
						logger.Println(i, "得到", tmpItem.ID)
					} else {
						logger.Printf("得到 %d星卡: %s-%s", tmpDBItem.Rarity, tmpDBItem.Title, tmpDBItem.Name)
					}
					charList = append(charList, *tmpItem)
				}
			}
		case 2:
			//logger.Println(i, "得到成長卡/冶鍊卡", data)
			list := data.(map[string]interface{})["data"].([]interface{})
			for _, item := range list {
				tmpItem := &models.GachaResultItem{}
				tmpDBItem := &models.Chararein{}
				if err := utils.Map2Struct(item.(map[string]interface{}), tmpItem); err != nil {
					logger.Println("Unable to convert to struct", err)
				} else {
					query := bson.M{"id": tmpItem.ItemID}
					controllers.GeneralQuery(metadata.DB, "chararein", query, tmpDBItem)
					logger.Println(i, "得到", tmpDBItem.Name)
					itemList = append(itemList, *tmpItem)
				}
			}
		case 14:
			continue
		case 105:
			list := data.(map[string]interface{})["data"].([]interface{})
			for _, item := range list {
				tmpItem := &models.GachaResultWeapon{}
				tmpDBItem := &models.Weapon{}
				if err := utils.Map2Struct(item.(map[string]interface{}), tmpItem); err != nil {
					logger.Println("Unable to convert to struct", err)
				} else {
					query := bson.M{"id": tmpItem.ItemID}
					if err := controllers.GeneralQuery(metadata.DB, "evolve", query, tmpDBItem); err != nil {
						logger.Println(i, "得到", tmpItem.ItemID)
					} else {
						logger.Println(i, "得到", tmpDBItem.Name)
					}
					weaponList = append(weaponList, *tmpItem)
				}
			}
		default:
			logger.Println(dataType)
		}
	}
	gachaResult["char"] = charList
	gachaResult["item"] = itemList

	return gachaResult
}

func doDebug(metadata *clients.Metadata, section string) {
	api := "data/weaponlist"
	param := map[string]interface{}{}
	ret, _ := general.GeneralAction(api, metadata.Sid, param)
	logger.Println(utils.Map2JsonString(ret))
}

func doUpdateDB(metadata *clients.Metadata, section string) {
	controllers.UpdateDB(metadata)
}

func getPresents(metadata *clients.Metadata, excludeTypes []string) {
	if presents, res := present.GetPresnetList(metadata.Sid); res == 0 {
		for _, p := range presents.Data {
			if utils.InArray(p.Data.Type, excludeTypes) {
				continue
			}
			if _, err := present.ReceievePresent(p.Idx, metadata.Sid); err == 0 {
				logger.Printf("-> 接收禮物 {%+v}\n", p)
			} else {
				logger.Printf("-> 接收禮物失敗 {%s}, %s\n", p.Text, err)
			}
		}
	} else {
		logger.Println(res)
	}
}

func doBuy(metadata *clients.Metadata, section string) {
	count, _ := metadata.Config.Int(section, "Count")
	itemType, _ := metadata.Config.String(section, "Type")

	for i := 0; i <= count; i++ {
		logger.Printf("#%d 購買道具, %s", i+1, itemType)
		if resp, res := item.BuyItemByType(itemType, metadata.Sid); res == 0 {
			logger.Println("\t-> 完成")
		} else {
			logger.Println("\t-> 失敗")
			logger.Println(resp, res)
		}
	}
}

func doShowChars(metadata *clients.Metadata, section string) {
	autoCompose, _ := metadata.Config.Bool("GENERAL", "AutoCompose")

	chars, _ := dyno.GetSlice(metadata.AllData, "body", 6, "data")
	threshold := 5
	//logger.Println(chars)
	for _, c := range chars {
		cardInfo := &models.Charainfo{} // for mongodb result
		charData := &models.CharaData{} // for alldata structure
		utils.Map2Struct(c.(map[string]interface{}), charData)
		if charData.Type != 0 {
			continue // non-char
		}
		query := bson.M{"cid": charData.ID}
		if err := controllers.GeneralQuery(metadata.DB, "charainfo", query, &cardInfo); err != nil {
			logger.Println(charData.ID)
		} else {
			if cardInfo.Rarity >= threshold {
				logger.Printf("%d, %s-%s, 目前等級: %d, 界限突破:%d",
					cardInfo.Cid, cardInfo.Title, cardInfo.Name, charData.Lv, charData.LimitBreak)
				if autoCompose == false {
					continue
				}
				for charData.Lv < charData.Maxlv {
					if ret, err := card.Compose(metadata, charData.Idx, 0); err == 0 {
						//log.Println(utils.Map2JsonString(res), err)
						lv, _ := dyno.GetFloat64(ret, "base_card", "lv")
						maxLv, _ := dyno.GetFloat64(ret, "base_card", "maxlv")
						logger.Printf("目前進度 %.0f/%0.f\n", lv, maxLv)
						charData.Lv = int(lv)
						charData.Maxlv = int(maxLv)
						//os.Exit(0)
					} else {
						logger.Printf("Unable to compose: %s\n", utils.Map2JsonString(ret))
						return
					}
				}
			}
		}
	}
}

func doPassword(metadata *clients.Metadata, section string) {
	tempPassword := "aaa123"

	resp, _ := user.GetAccount(metadata.Sid)
	account := resp["account"].(string)
	//logger.Printf("%s\n", utils.Map2JsonString(resp))

	resp, _ = user.SetPassword(tempPassword, metadata.Sid)
	//logger.Println(utils.Map2JsonString(resp))

	logger.Printf("Account: [%s] has set password: [%s]", account, tempPassword)
}

func doTakeOver(metadata *clients.Metadata, section string) {
	tempPassword := "aaa123"
	account, _ := metadata.Config.String("GENERAL", "Account")
	uuid, _ := metadata.Config.String("GENERAL", "Uid")
	if ret, err := user.Takeover(uuid, account, tempPassword); err != 0 {
		logger.Println("Unable to takeover account", utils.Map2JsonString(ret))
	} else {
		logger.Println("帳號轉移完成")
	}

}

func doStatus(metadata *clients.Metadata, section string) {
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
	specialData := metadata.AllData["body"].([]interface{})[8].(map[string]interface{})["data"]
	stoneCount := metadata.AllData["body"].([]interface{})[12].(map[string]interface{})["data"]
	logger.Printf("精靈石 = %.0f\n", stoneCount.(float64))
	for _, item := range specialData.([]interface{}) {
		itemId := item.(map[string]interface{})["item_id"]
		cnt := item.(map[string]interface{})["cnt"]
		//logger.Println(itemId, reflect.TypeOf(itemId))
		//logger.Println(cnt, reflect.TypeOf(cnt))
		//fmt.Println(itemId)
		if val, ok := itemMapping[int(itemId.(float64))]; ok {
			switch reflect.TypeOf(cnt).Kind() {
			case reflect.String:
				logger.Printf("%s = %s\n", val, cnt.(string))
			case reflect.Float64:
				logger.Printf("%s = %.0f\n", val, cnt.(float64))
			}
		}
	}

	userData := metadata.AllData["body"].([]interface{})[4].(map[string]interface{})["data"]
	//logger.Println(utils.Map2JsonString(metadata.AllData))
	for k, v := range userData.(map[string]interface{}) {
		if utils.InArray(k, targets) {
			switch v.(type) {
			case float64, float32:
				logger.Printf("%s = %.0f\n", k, v)
			default:
				logger.Printf("%s = %v\n", k, v)
			}
		}
	}
}

func doShowAllData(metadata *clients.Metadata, section string) {
	fmt.Println(utils.Map2JsonString(metadata.AllData))
}

func recoverAp(metadata *clients.Metadata) (ret int) {
	resp, res := user.RecoveryAp(1, 1, metadata.Sid)
	ret = 0
	switch res {
	case 0:
		logger.Println("恢復體力完成")
	case 703:
		logger.Println("恢復體力失敗，體力果實不足，購買體力果實")
		if _, err := item.BuyItemByType(item.AP_FRUIT, metadata.Sid); err != 0 {
			logger.Println("購買體力果實失敗")
			ret = 1
		}
	default:
		logger.Println("未知的錯誤")
		logger.Println(utils.Map2JsonString(resp))
		ret = 2
	}
	return ret
}

func doQuest(metadata *clients.Metadata, section string) {
	//logger.Println("enter doQuest")
	conf := metadata.Config
	questInfo := quest.NewQuest()
	count, _ := conf.Int(section, "Count")
	infinite := false
	if count == -1 {
		infinite = true
	}

	// Read config to Struct
	utils.ParseConfig2Struct(conf, section, questInfo)
	current := 0
	for {
		current++
		if current > count && infinite == false {
			break
		}
		logger.Printf("#%d 開始關卡:[%d]", current, questInfo.QuestId)
		resp, res := questInfo.StartQeust(metadata)
		switch res {
		case 0:
			//do nothing
		case 103:
			logger.Println("AP 不足，使用體力果")
			if ret := recoverAp(metadata); ret != 0 {
				logger.Println("回復AP失敗")
				break
			}
			current -= 1
			continue
		default:
			logger.Println("未知的錯誤")
			logger.Println(resp)
			break
		}
		resp, res = questInfo.EndQeust(metadata)
		//fmt.Println(utils.Map2JsonString(resp))
		switch res {
		case 0:
			logger.Println("關卡完成")
			//Check if need to sell cards
		case 1:
			logger.Println("關卡失敗，已被登出")
		default:
			logger.Println("未知的錯誤")
			logger.Println(resp)
		}

		if questInfo.AutoRaid {
			//time.Sleep(time.Second)
			//logger.Println("Checking 魔神戰")
			raidQuest(metadata, questInfo.AutoRaidRecover, section)
		}
	}
}

func raidQuest(metadata *clients.Metadata, recovery bool, section string) {
	//ret, _ := raid.RaidList(metadata.Sid)
	if bossInfo := raid.GetRaidBossInfo(metadata.Sid); bossInfo != nil {
		//logger.Printf("%+v", bossInfo)
		logger.Printf("魔神來襲! BossId = %d, bossLv = %d\n", bossInfo.BossID, bossInfo.BossParam.Lv)
		param := map[string]interface{}{}
		ret, err := bossInfo.StartQuest(metadata, param)

		switch err {
		case 0:
			if ret, err := bossInfo.EndQuest(metadata, param); err != 0 {
				logger.Println("無法結束魔神戰")
				logger.Println(utils.Map2JsonString(ret))
			} else {
				bossInfo.GetBonus(metadata, param)
			}
		case 104:
			logger.Println("魔神體力不足")
			if recovery {
				if ret, err := user.RecoveryBp(0, 2, metadata.Sid); err != 0 {
					logger.Println("\t ->回復體力失敗", ret)
				} else {
					logger.Println("\t ->回復體力成功")
				}
			}

		case 603:
		case 608:
			logger.Println("魔神已結束")
			bossInfo.EndQuest(metadata, param)
			bossInfo.GetBonus(metadata, param)
		default:
			logger.Println("未知的魔神戰錯誤", utils.Map2JsonString(ret))
		}

	} else {
		logger.Println("No Boss info found")
	}
}

func doDisciple(metadata *clients.Metadata, section string) {
	teacherId, _ := metadata.Config.Int(section, "TeacherId")
	if isGraduate, err := metadata.Config.Bool(section, "Graduate"); err != nil {
		// do nothing, use va
	} else {
		// use config value
		teacher.IS_GRADUATED = isGraduate
	}
	logger.Println("Teacher ID", teacherId, "Is Graduate?", teacher.IS_GRADUATED)

	if teacher.IS_GRADUATED {
		// thanks teacher
		for _, lv := range []int{5, 10, 15, 20, 25, 30, 35, 40, 45} {
			if ret, err := teacher.ThanksAchievement(metadata, lv); err != 0 {
				logger.Printf("UID %s 無法 給與 Rank %d 獎勵, res = %s\n", metadata.Uid, lv, utils.Map2JsonString(ret))
			} else {
				logger.Printf("UID %s 給與 Rank %d 獎勵\n", metadata.Uid, lv)
			}
		}
		if ret, err := teacher.ThanksGgraduate(metadata); err != 0 {
			logger.Printf("UID %s 無法畢業, res = %s, trying...\n", metadata.Uid, utils.Map2JsonString(ret))
            time.Sleep(3 * time.Second)
            teacher.ThanksGgraduate(metadata)
		} else {
			logger.Printf("UID %s 畢業\n", metadata.Uid)
            teacher.IS_GRADUATED = false
		}
	} else {
		logger.Printf("UID %s 選擇 %d 為師父", metadata.Uid, teacherId)
		if ret, err := teacher.ApplyTeacher(metadata, teacherId); err != 0 {
			logger.Printf("UID %s 選擇 %d 為師父 失敗! %d", metadata.Uid, teacherId, err)
			logger.Println(utils.Map2JsonString(ret))
			os.Exit(-1)
		}
	}
}

func doTeacher(metadata *clients.Metadata, section string) {
	if res, err := teacher.EnableTeacher(metadata); err != 0 {
		logger.Println("Unable to enable teacher", utils.Map2JsonString(res))
	} else {
		logger.Println("Enable teacher success")
	}

}

func doResetDisciple(metadata *clients.Metadata, section string) {
	param := map[string]interface{}{}
	disciples := teacher.ListDisciple(metadata, param)
	for _, d := range disciples {
		fmt.Println("Trying to reset disciple", d.UID, d.Resetable, d.Name)
		if resp, err := teacher.ResetDisciple(metadata, d.UID); err != 0 {
			logger.Println("Unable to reset Disciple", d.UID, utils.Map2JsonString(resp))
		} else {
			logger.Println("Reset Disciple success")
		}
	}
}

func doCompose(metadata *clients.Metadata, section string) {
	//mockList := []int{96064,96064,96064,96064,96064}
	//weapon.Compose(metadata, mockList, 14)
	//os.Exit(0)
	weaponListRank3 := make([]int, 0)
	weaponListRank4 := make([]int, 0)
	weaponBaseRank5Idx := 0
	count, _ := metadata.Config.Int(section, "Count")
	eid := -1
	if tmpEid, err := metadata.Config.Int(section, "Eid"); err != nil {
		eid = -1
		logger.Println("EID:", eid)
	} else {
		eid = tmpEid
	}

	baseWeaponId, _ := metadata.Config.Int(section, "BaseWeaponID")
	weaponData := map[string]interface{}{
		"kind":    "item",
		"type":    "weapon_ev",
		"id":      baseWeaponId,
		"val":     1,
		"price":   10,
		"buy_cnt": 1,
	}
	//
	targetWeaponKeyword, _ := metadata.Config.String(section, "TargetsKeyWords")
	targetWeaponKeywordList := strings.Split(strings.Replace(targetWeaponKeyword, " ", "", -1), ",")
	foundTarget := false

	for i := 0; i < count; i++ {
		buyCount := 5
		if weaponBaseRank5Idx != 0 {
			buyCount = 4
		}

		//買 五把或四把 三星武器
		for j := 0; j < buyCount; j++ {
			//logger.Println("購買武器", weaponData["id"])
			if ret, err := item.BuyItemGeneral(metadata, weaponData); err != 0 {
				logger.Println("Unable to buy item", utils.Map2JsonString(ret))
				os.Exit(0)
			} else {
				baseWeaponIdx, _ := dyno.GetFloat64(ret, "body", 1, "data", 0, "item_id")
				//logger.Println(baseWeaponIdx)
				weaponListRank3 = append(weaponListRank3, int(baseWeaponIdx))
			}
		}
		// 五把三星器武器，鍊成四星武器
		if len(weaponListRank3) == 5 {
			//logger.Println("開始鍊金，三星*5")
			if ret, err := weapon.Compose(metadata, weaponListRank3, eid); err != 0 {
				logger.Println("Compose error", utils.Map2JsonString(ret), err)
				os.Exit(0)
			} else {
				weaponListRank3 = nil // clear slice
				body, _ := dyno.GetSlice(ret, "body")
				lastIndex := len(body) - 1
				itemId, _ := dyno.GetFloat64(ret, "body", lastIndex-1, "data", 0, "item_id")
				//logger.Println("得到武器：", itemId)
				weaponListRank4 = append(weaponListRank4, int(itemId))
			}
		}
		// 有一張基底五星武器，且有四張三星武器
		if weaponBaseRank5Idx != 0 && len(weaponListRank3) == 4 {
			// self.logger.info(u'開始鍊金 -  5星*1 + 3星*4')
			weaponListRank3 = append(weaponListRank3, weaponBaseRank5Idx)
			ret, _ := weapon.Compose(metadata, weaponListRank3, eid)
			weaponListRank3 = nil // clear slice
			body, _ := dyno.GetSlice(ret, "body")
			lastIndex := len(body) - 1
			itemId, _ := dyno.GetFloat64(ret, "body", lastIndex-1, "data", 0, "item_id")
			myWeapon := models.Evolve{}
			query := bson.M{"id": int(itemId)}
			controllers.GeneralQuery(metadata.DB, "evolve", query, &myWeapon)
			for _, targetKeyWord := range targetWeaponKeywordList {
				if strings.Index(myWeapon.Name, targetKeyWord) != -1 {
					foundTarget = true
					break
				}
			}

			if foundTarget {
				logger.Println("!! 得到神器：", myWeapon.Name)
				weaponBaseRank5Idx = 0
				break // break main for loop
			} else {
				logger.Println("得到武器：", myWeapon.Name)
				weaponBaseRank5Idx = int(itemId)
			}
		} else if len(weaponListRank4) == 5 {
			// 鍊出做為基底的五星武器
			//logger.Println("開始鍊金，四星*5")
			if ret, err := weapon.Compose(metadata, weaponListRank4, eid); err != 0 {
				logger.Println("Compose error", utils.Map2JsonString(ret), err)
				os.Exit(0)
			} else {
				weaponListRank4 = nil // clear slice
				body, _ := dyno.GetSlice(ret, "body")
				lastIndex := len(body) - 1
				itemId, _ := dyno.GetFloat64(ret, "body", lastIndex-1, "data", 0, "item_id")
				//logger.Println("得到武器：", itemId)
				weaponBaseRank5Idx = int(itemId)
			}
		}
	}
}

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	start()
}

func dumpUser(u *clients.Metadata) {
	logger.Printf("%+v\n", *u)
}
