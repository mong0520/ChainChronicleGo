package clients

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mong0520/ChainChronicleGo/models"
	"github.com/robfig/config"
	"gopkg.in/mgo.v2"
)

type Metadata struct {
	Config              *config.Config
	Sid                 string `json:"first"`
	Uid                 string
	Token               string
	Flow                []string `json:"flow"`
	AllData             map[string]interface{}
	AllDataS            *models.AllData
	DB                  *mgo.Session
	RedisConn           redis.Conn
	CharInfo            []models.Charainfo
	CharData            []models.CharaData
	ExplorerExcludeCids []int
}
