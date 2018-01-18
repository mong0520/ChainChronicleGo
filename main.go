package main

import (
	"log"

	"github.com/mong0520/ChainChronicleGo/clients"

	"github.com/mong0520/ChainChronicleGo/clients/quest"
	"github.com/mong0520/ChainChronicleGo/clients/session"
	"github.com/mong0520/ChainChronicleGo/utils"
	"github.com/robfig/config"
)

var user = &clients.User{}
var logger *log.Logger

func init() {
	logger = utils.GetLogger()
	config, err := config.ReadDefault("conf/mong.conf")
	if err != nil {
		logger.Fatalln("Unable to read config, ", err)
		return
	}
	user.Config = config
}

func start() {
	dumpUser(user)
	utils.DumpConfig(user.Config)
	uid, _ := user.Config.String("GENERAL", "Uid")
	logger.Println(uid)
	token, _ := user.Config.String("GENERAL", "Token")
	//flowString, _ := conf.String("GENERAL", "Flow")
	sid := session.Login(uid, token)
	user.Sid = sid
	dumpUser(user)

	for i := 1; i < 10; i++ {
		questInfo := &quest.Quest{}
		questInfo.Type = 4
		questInfo.Fid = 300467
		questInfo.Qid = 220134
		questInfo.Hcid = 0
		questInfo.Htype = 0
		questInfo.Lv = 0
		questInfo.Pt = 0
		questInfo.Version = 2
		questInfo.StartQeust(user)

		questInfo.Res = 1
		questInfo.Bt = 10
		questInfo.Wc = 13
		questInfo.Wn = 1
		questInfo.Cc = 1
		questInfo.Time = "10.0"
		questInfo.D = 1
		questInfo.S = 1
		questInfo.EndQeust(user)
	}

}

func main() {
	start()
}

func dumpUser(u *clients.User) {
	logger.Printf("%+v\n", *u)
}
