package present

import (
    "fmt"
    "github.com/mong0520/ChainChronicleGo/clients/general"
    "github.com/goinggo/mapstructure"
)

type Presents struct {
    Data []struct {
        Data struct {
            ID   int    `json:"id"`
            Type string `json:"type"`
            Val  int    `json:"val"`
        } `json:"data"`
        Idx    int    `json:"idx"`
        Reason int    `json:"reason"`
        Text   string `json:"text"`
    } `json:"data"`
    Proc int `json:"proc"`
    Type int `json:"type"`
}

var api = "present"
func GetPresnetList(sid string) (presents *Presents, res int) {
    res = 0
    action := "list"
    api := fmt.Sprintf("%s/%s", api, action)
    param := map[string]interface{}{}
    resp, _ := general.GeneralAction(api, sid, param)
    presentList := resp["body"].([]interface{})[0].(map[string]interface{})
    //fmt.Println(presentList)
    presents = &Presents{}
    if err := mapstructure.Decode(presentList, presents) ; err != nil {
        fmt.Println(err)
        return nil, -1
    }
    //fmt.Printf("%+v", presents)
    //fmt.Println(utils.Map2JsonString(presentList))
    //fmt.Println(resp)
    return presents, res
}


func (p *Presents) ReceievePresent(presentType int){
    for _, present := range p.Data{
        fmt.Printf("%+v\n", present)
    }
}