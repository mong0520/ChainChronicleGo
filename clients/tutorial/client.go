package tutorial

import (
    "github.com/mong0520/ChainChronicleGo/clients/general"
)


func Tutorial(sid string, entry bool, param map[string]interface{}) (resp map[string]interface{}, res int) {
    api := ""
    if entry{
        api = "tutorial/entry"
    }else{
        api = "tutorial"
    }
    return general.GeneralAction(api, sid, param)
}