package card

import (
    "github.com/mong0520/ChainChronicleGo/clients"
    "github.com/mong0520/ChainChronicleGo/clients/general"
)

func Sell(u *clients.Metadata, cid int) (resp map[string]interface{}, res int) {
    api := "card/sell"
    param := map[string]interface{}{
        "c": cid,
    }
    return general.GeneralAction(api, u.Sid, param)
}
