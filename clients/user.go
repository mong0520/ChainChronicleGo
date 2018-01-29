package clients

import (
    "github.com/robfig/config"
    "gopkg.in/mgo.v2"
)

type Metadata struct {
    Config *config.Config
    Sid string `json:"first"`
    Uid string
    Token string
    Flow []string `json:"flow"`
    AllData map[string]interface{}  `json:"all_data"`
    DB *mgo.Session
}


