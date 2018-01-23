package clients

import "github.com/robfig/config"

type Metadata struct {
    Config *config.Config
    Sid string `json:"first"`
    Flow []string `json:"flow"`
    AllData map[string]interface{}  `json:"all_data"`
}


