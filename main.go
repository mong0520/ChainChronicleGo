package main

import (
	"log"

	"github.com/mong0520/ChainChronicleGo/clients"

	//"github.com/mong0520/ChainChronicleGo/clients/quest"
	"github.com/mong0520/ChainChronicleGo/clients/session"
	"github.com/mong0520/ChainChronicleGo/utils"
	"github.com/robfig/config"
    "strings"
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

	//for i := 1; i < 10; i++ {
	//	questInfo := &quest.Quest{}
	//	questInfo.Type = 4
	//	questInfo.Fid = 300467
	//	questInfo.Qid = 220134
	//	questInfo.Hcid = 0
	//	questInfo.Htype = 0
	//	questInfo.Lv = 0
	//	questInfo.Pt = 0
	//	questInfo.Version = 2
	//	questInfo.StartQeust(user)
    //
	//	questInfo.Res = 1
	//	questInfo.Bt = 10
	//	questInfo.Wc = 13
	//	questInfo.Wn = 1
	//	questInfo.Cc = 1
	//	questInfo.Time = "10.0"
	//	questInfo.D = 1
	//	questInfo.S = 1
	//	questInfo.EndQeust(user)
	//}

}

func doQuest(user *clients.User, section string){
    logger.Println("enter doQuest")
    conf := user.Config

    // Parse config
    questRaw, _ := conf.String(section, "QuestId")

    questInfo := &quest.Quest{}


}

func main() {
	start()
}

func dumpUser(u *clients.User) {
	logger.Printf("%+v\n", *u)
}
