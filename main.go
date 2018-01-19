package main

import (
	"log"
	"strings"

	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/session"
	"github.com/mong0520/ChainChronicleGo/utils"
	"github.com/mong0520/ChainChronicleGo/clients/quest"
	"github.com/mong0520/ChainChronicleGo/clients/user"

	"github.com/robfig/config"
    "flag"
)

var metadata = &clients.Metadata{}
var logger *log.Logger
var actionMapping = map[string]interface{}{
	"QUEST": doQuest,
	"STATUS": doStatus,
	"PASSWORD": doPassword,
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
    configPath := flag.String("config", "", "Config file path")
    action := flag.String("action", "", "Action to run")
    //flag.BoolVar(&showVersion, "version", false, "Print version information.")
    flag.Parse()

    logger = utils.GetLogger()
    config, err := config.ReadDefault(*configPath)
    if err != nil {
        logger.Fatalln("Unable to read config, ", err)
        return
    }
    metadata.Config = config
	utils.DumpConfig(metadata.Config)
	uid, _ := metadata.Config.String("GENERAL", "Uid")
	logger.Println(uid)
	token, _ := metadata.Config.String("GENERAL", "Token")
	if *action == ""{
        flowString, _ := metadata.Config.String("GENERAL", "Flow")
        metadata.Flow = strings.Split(flowString, ",")
    }else{
        flowString := *action
        metadata.Flow = strings.Split(flowString, ",")
    }

	sid := session.Login(uid, token)
	metadata.Sid = sid
	//dumpUser(metadata)
	for idx, flow := range metadata.Flow {
		logger.Printf("[%d] Processing [%s]\n", idx, flow)
		doAction(flow)
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
	userData := resp["body"].([]interface{})[4].(map[string]interface {})["data"]
	for k, v := range userData.(map[string]interface{}) {
		if utils.InArray(k, targets) {
			switch v.(type){
			case float64, float32:
				logger.Printf("%s = %.0f\n", k, v)
			default:
				logger.Printf("%s = %v\n", k, v)
			}
		}
	}
}

func doQuest(user *clients.Metadata, section string) {
	logger.Println("enter doQuest")
	conf := user.Config
	questInfo := quest.NewQuest()

	// Read config to Struct
	utils.ParseConfig2Struct(conf, section, questInfo)

	resp, res := questInfo.StartQeust(user)
	switch res {
	case 0:
		//do nothing
	case 103:
		logger.Println("AP 不足，使用體力果")
	default:
		logger.Println("未知的錯誤")
		logger.Println(resp)
	}

	resp, res = questInfo.EndQeust(user)
	switch res {
	case 0:
		logger.Println("關卡完成")
	case 1:
		logger.Println("關卡失敗，已被登出")
	default:
		logger.Println("未知的錯誤")
		logger.Println(resp)
	}
}

func main() {
	start()
}

func dumpUser(u *clients.Metadata) {
	logger.Printf("%+v\n", *u)
}
