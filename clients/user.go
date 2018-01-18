package clients

import "github.com/robfig/config"

type User struct {
    Config *config.Config
    Sid string `json:"first"`
    Flow []string `json:"flow"`
}
