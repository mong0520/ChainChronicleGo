package card

import (
    "github.com/mong0520/ChainChronicleGo/clients"
    "github.com/mong0520/ChainChronicleGo/clients/general"
    "encoding/json"
    "fmt"
)

func Sell(u *clients.Metadata, cid int) (resp map[string]interface{}, res int) {
    _ = fmt.Sprintf("")
    api := "card/sell"
    param := map[string]interface{}{
        "c": cid,
    }
    return general.GeneralAction(api, u.Sid, param)
}

func Compose(metadata *clients.Metadata, baseId int, expup_id int) (respMap map[string]interface{}, res int) {
    api := "card/compose"
    im := map[string]int{
        "90904": 10,  // 四星成長卡
    }

    imString, _ := json.Marshal(im)
    param := map[string]interface{}{
        "ba": baseId,
        "im": string(imString),
    }
    return general.GeneralAction(api, metadata.Sid, param)

}
