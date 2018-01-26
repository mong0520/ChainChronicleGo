package main

import (
	"log"
	"strings"

	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/item"
	"github.com/mong0520/ChainChronicleGo/clients/quest"
	"github.com/mong0520/ChainChronicleGo/clients/session"
	"github.com/mong0520/ChainChronicleGo/clients/user"
	"github.com/mong0520/ChainChronicleGo/utils"

	"fmt"
	"github.com/icza/dyno"
	"github.com/jessevdk/go-flags"
	"github.com/mong0520/ChainChronicleGo/clients/gacha"
	"github.com/mong0520/ChainChronicleGo/clients/tutorial"
	"github.com/robfig/config"
	"github.com/satori/go.uuid"
	"os"
	"time"
    "github.com/mong0520/ChainChronicleGo/clients/present"
	"github.com/deckarep/golang-set"
	"github.com/mong0520/ChainChronicleGo/clients/raid"
)

type Options struct {
	ConfigPath string `short:"c" long:"config" description:"Config path" required:"true"`
	Action     string `short:"a" long:"action" description:"Action to run" required:"false"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)
var metadata = &clients.Metadata{}
var logger *log.Logger
var actionMapping = map[string]interface{}{
	"QUEST":    doQuest,
	"STATUS":   doStatus,
	"PASSWORD": doPassword,
	"BUY":      doBuy,
	"GACHA":    doGacha,
	"TUTORIAL": doTutorial,
	"DRAMA":    doDrama,
	"DEBUG":    doDebug,
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

func start() {
	logger = utils.GetLogger()
	config, err := config.ReadDefault(options.ConfigPath)
	if err != nil {
		logger.Fatalln("Unable to read config, ", err)
		return
	}

	metadata.Config = config
	//utils.DumpConfig(metadata.Config)
	uid, _ := metadata.Config.String("GENERAL", "Uid")
	logger.Println(uid)
	token, _ := metadata.Config.String("GENERAL", "Token")
	if options.Action == "" {
		flowString, _ := metadata.Config.String("GENERAL", "Flow")
		metadata.Flow = strings.Split(flowString, ",")
	} else {
		flowString := options.Action
		metadata.Flow = strings.Split(flowString, ",")
	}

	sid := session.Login(uid, token)
	alldata, _ := user.GetAllData(sid)
	metadata.AllData = alldata
	metadata.Sid = sid
	//dumpUser(metadata)
	for _, flow := range metadata.Flow {
		logger.Printf("Processing [%s]\n", flow)
		doAction(flow)
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

    for{
        //logger.Println(v, reflect.TypeOf(v))
        //logger.Println(dyno.Get(metadata.AllData, "body", 29, "data", flag, "id"))
        if qType, err := dyno.GetFloat64(metadata.AllData, "body", 29, "data", flag, "type") ; err != nil{
            logger.Println(qType, err)
        }else{
            questInfo.Type = int(qType)
        }
        if qId, err := dyno.GetFloat64(metadata.AllData, "body", 29, "data", flag, "id") ; err != nil{
            logger.Println(qId, err)
        }else{
            questInfo.QuestId = int(qId)
        }

        counter += 1
        if counter >= gradudateThreshold{
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
        switch err{
        case 0:
            _, err = questInfo.EndQeust(metadata)
            logger.Printf("%d/%d 完成關卡\n", counter, gradudateThreshold)
			errSet.Add(err)
        case 103:
            logger.Printf("體力不足\n")
            if _, err := user.RecoveryAp(1,1, metadata.Sid) ; err != 0 {
                logger.Println("無法恢復體力")
                panic(err)
            }
        default:
            logger.Println("未知的錯誤")
			errSet.Add(err)
            logger.Println(utils.Map2JsonString(resp))
            if resp, err := questInfo.GetTreasure(metadata) ; err != 0{
                logger.Println(resp)
            }
        }
        if errSet.Contains(0) == false {
        	logger.Printf("Unknown error in drama: %s", errSet)
			currentRetryCount ++
        	if currentRetryCount >= maxRetryCount {
        		uid, _ := metadata.Config.String("GENERAL", "Uid")
				logger.Printf("UID %s is is uable to complete Drama", uid)
			} else{
				logger.Println("Retry...")
				continue
			}
		} else{
			currentRetryCount = 0
			if questInfo.QuestId == lastQid{
				continue
			}else{
				if flag >= 4{
					dramaLevel = 2
				}
				flag ++
			}
		}
    }
	alldata, _ := user.GetAllData(metadata.Sid)
	metadata.AllData = alldata
	if currentLv, err := dyno.GetFloat64(metadata.AllData, "body", 4, "data", "lv") ; err != nil{
		logger.Println(err)
	}else{
		logger.Printf("Current LV %0.f", currentLv)
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
	sid := session.Login(newUid, "")
	logger.Println(sid)
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
			if resp, err := questInfo.EndQeust(metadata) ; err != 0 {
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
				if resp, err := tutorial.Tutorial(sid, false, param) ; err != 0 {
                    logger.Println(utils.Map2JsonString(resp), err)
                    break
                }
			} else {
				param := map[string]interface{}{
					"tid": t["tid"],
				}
				if resp, err := tutorial.Tutorial(sid, false, param) ; err != 0 {
                    logger.Println(utils.Map2JsonString(resp), err)
                    break
                }
			}
		}
	}
	logger.Printf("新帳號完成新手教學, UID = %s, OpenID = %.0f\n", newUid, openId)
    getPresents(metadata)
}

func doGacha(metadata *clients.Metadata, section string) {
	gachaInfo := gacha.NewGachaInfo()
	utils.ParseConfig2Struct(metadata.Config, section, gachaInfo)
	if resp, ret := gachaInfo.Gacha(metadata); ret == 0 {
		logger.Println(utils.Map2JsonString(resp))
	} else {
		logger.Println(utils.Map2JsonString(resp))
	}

}

func doDebug(metadata *clients.Metadata, section string) {
    //if presents, res := present.GetPresnetList(metadata.Sid) ; res == 0 {
    //    presents.ReceievePresent(0)
    //}
    return
}

func getPresents(metadata *clients.Metadata) {
    if presents, res := present.GetPresnetList(metadata.Sid) ; res == 0 {
        for _, p := range presents.Data{
        	if _, err := present.ReceievePresent(p.Idx, metadata.Sid) ; err == 0 {
        		logger.Printf("-> 接收禮物 {%+v}\n", p)
			}else{
				logger.Printf("-> 接收禮物失敗 {%s}, %s\n", p.Text, err)
			}
		}
    }else{
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

func doPassword(metadata *clients.Metadata, section string) {
	tempPassword := "aaa123"

	resp, _ := user.GetAccount(metadata.Sid)
	account := resp["account"].(string)
	logger.Printf("%s\n", utils.Map2JsonString(resp))

	resp, _ = user.SetPassword(tempPassword, metadata.Sid)
	logger.Println(utils.Map2JsonString(resp))

	logger.Printf("Account: [%s] has set password: [%s]", account, tempPassword)
}

func doStatus(metadata *clients.Metadata, section string) {
	targets := []string{"comment", "uid", "heroName", "open_id", "lv", "cardMax", "accept_disciple", "name",
		"friendCnt", "only_friend_disciple", "staminaMax", "zuLastRefilledScheduleId", "uzu_key"}
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
	logger.Println("enter doQuest")
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
			time.Sleep(1)
			logger.Println("Checking 魔神戰")
			raidQuest(metadata, questInfo.AutoRaidRecover, section)
		}
	}
}

func raidQuest(metadata *clients.Metadata, recovery bool, section string){
	//ret, _ := raid.RaidList(metadata.Sid)
	if bossInfo := raid.GetRaidBossInfo(metadata.Sid) ; bossInfo != nil {
		//logger.Printf("%+v", bossInfo)
		logger.Printf("魔神來襲! BossId = %d, bossLv = %d\n", bossInfo.BossID, bossInfo.BossParam.Lv)
		param := map[string]interface{}{}
		ret, err := bossInfo.StartQuest(metadata, param)

		switch err{
        case 0:
            if ret, err := bossInfo.EndQuest(metadata, param) ; err != 0{
                logger.Println("無法結束魔神戰")
                logger.Println(utils.Map2JsonString(ret))
            }else{
                bossInfo.GetBonus(metadata, param)
            }
        case 104:
            logger.Println("魔神體力不足")
            if recovery {
                if ret, err := user.RecoveryBp(0, 2, metadata.Sid) ; err != 0 {
                    logger.Println("Recover Raid point failed", ret)
                }else{
                    logger.Println("Recover Raid point success")
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


	}else{
		logger.Println("No Boss info found")
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
