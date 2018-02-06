package clients

import (
    "github.com/robfig/config"
    "gopkg.in/mgo.v2"
    "github.com/mong0520/ChainChronicleGo/models"
)

type Metadata struct {
    Config *config.Config
    Sid string `json:"first"`
    Uid string
    Token string
    Flow []string `json:"flow"`
    AllData map[string]interface{}
    AllDataS *models.AllData
    DB *mgo.Session
}


