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

	"github.com/jessevdk/go-flags"
	"github.com/robfig/config"
	"os"
	"time"
    "github.com/mong0520/ChainChronicleGo/clients/gacha"
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
	metadata.Sid = sid
	//dumpUser(metadata)
	for _, flow := range metadata.Flow {
		logger.Printf("Processing [%s]\n", flow)
		doAction(flow)
	}
}

func doGacha(metadata *clients.Metadata, section string) {
    gachaInfo := gacha.NewGachaInfo()
    utils.ParseConfig2Struct(metadata.Config, section, gachaInfo)
    if resp, ret := gachaInfo.Gacha(metadata) ; ret == 0{
        logger.Println(utils.Map2JsonString(resp))
    }else{
        logger.Println(utils.Map2JsonString(resp))
    }

}

func doDebug(metadata *clients.Metadata, section string) {
	//if resp, res := user.RecoveryAp(1, metadata.Sid) ; res == 0 {
	//    logger.Println("Success")
	//} else {
	//    logger.Printf("Failed:\n%s\n", utils.Map2JsonString(resp))
	//}
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

	resp, _ := user.GetAllData(metadata.Sid)
	userData := resp["body"].([]interface{})[4].(map[string]interface{})["data"]
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
	resp, res := user.RecoveryAp(1, metadata.Sid)
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
			// start raid
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
