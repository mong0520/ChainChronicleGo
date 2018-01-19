package main

import (
	"log"

	"github.com/mong0520/ChainChronicleGo/clients"

	//"github.com/mong0520/ChainChronicleGo/clients/quest"
	"github.com/mong0520/ChainChronicleGo/clients/session"
	"github.com/mong0520/ChainChronicleGo/utils"
	"github.com/robfig/config"

    "strings"
    //"github.com/mong0520/ChainChronicleGo/clients/quest"
	"github.com/mong0520/ChainChronicleGo/clients/quest"
)

var user = &clients.User{}
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
	user.Config = config
}

func doAction(sectionName string){
    for action, actionFunction := range actionMapping{
        logger.Println(action, actionFunction)
        if strings.HasPrefix(sectionName, action){
            logger.Printf("### Current Flow = [%s] ###", sectionName)
            actionFunction.(func(*clients.User, string))(user, sectionName)
        }
    }

}

func start() {
	utils.DumpConfig(user.Config)
	uid, _ := user.Config.String("GENERAL", "Uid")
	logger.Println(uid)
	token, _ := user.Config.String("GENERAL", "Token")
	flowString, _ := user.Config.String("GENERAL", "Flow")
	user.Flow = strings.Split(flowString, ",")
	sid := session.Login(uid, token)
	user.Sid = sid
	dumpUser(user)

	for idx, flow := range user.Flow{
	    logger.Printf("[%d] Processing [%s]\n", idx, flow)
	    doAction(flow)
    }
}

func doQuest(user *clients.User, section string){
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
	switch res{
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

func dumpUser(u *clients.User) {
	logger.Printf("%+v\n", *u)
}
