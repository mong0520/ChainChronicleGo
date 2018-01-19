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
	"reflect"
)

var metadata = &clients.Metadata{}
var logger *log.Logger
var actionMapping = map[string]interface{}{
	"QUEST": doQuest,
}

func init() {
	logger = utils.GetLogger()
	config, err := config.ReadDefault("conf/mong.conf")
	if err != nil {
		logger.Fatalln("Unable to read config, ", err)
		return
	}
	metadata.Config = config
}

func doAction(sectionName string) {
	for action, actionFunction := range actionMapping {
		logger.Println(action, actionFunction)
		if strings.HasPrefix(sectionName, action) {
			logger.Printf("### Current Flow = [%s] ###", sectionName)
			actionFunction.(func(*clients.Metadata, string))(metadata, sectionName)
		}
	}

}

func start() {
	utils.DumpConfig(metadata.Config)
	uid, _ := metadata.Config.String("GENERAL", "Uid")
	logger.Println(uid)
	token, _ := metadata.Config.String("GENERAL", "Token")
	flowString, _ := metadata.Config.String("GENERAL", "Flow")
	metadata.Flow = strings.Split(flowString, ",")
	sid := session.Login(uid, token)
	metadata.Sid = sid
	dumpUser(metadata)
	doShowStatus(metadata)

	//resp, _ := user.GetAllData(metadata.Sid)
	//respStr, _ := json.Marshal(resp)
	//logger.Println(string(respStr))

    //
	//for idx, flow := range metadata.Flow {
	//	logger.Printf("[%d] Processing [%s]\n", idx, flow)
	//	doAction(flow)
	//}
}

func doShowStatus(metadata *clients.Metadata) {
	targets := []string{"comment", "uid", "heroName", "open_id", "lv", "cardMax", "accept_disciple", "name", "friendCnt",
	"only_friend_disciple", "staminaMax", "zuLastRefilledScheduleId", "uzu_key"}
	resp, _ := user.GetAllData(metadata.Sid)
	userData := resp["body"].([]interface{})[4].(map[string]interface {})["data"]
	//logger.Println(userData, reflect.TypeOf(userData))
	for k, v := range userData.(map[string]interface{}) {
		if utils.InArray(k, targets) {
			logger.Printf("%s = %v, Type = %s\n", k, v, reflect.TypeOf(v))
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
