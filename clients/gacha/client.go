package gacha

import (
    "github.com/mong0520/ChainChronicleGo/clients"
    "github.com/mong0520/ChainChronicleGo/clients/general"
)

type gacha struct {
    Type int
    Count int
    Area int
    Place int
    AutoSaleThreshold int
    Verbose bool
}

var initValue = -1

func (g *gacha) Gacha(u *clients.Metadata) (resp map[string]interface{}, res int) {
    api := "gacha"
    param := map[string]interface{}{
        "t": g.Type,
        "c": g.Count,
    }
    if g.Area != initValue{
        param["area"] = g.Area
    }
    if g.Place != initValue{
        param["place"] = g.Place
    }
    return general.GeneralAction(api, u.Sid, param)
}

func NewGachaInfo()(g *gacha){
    g = &gacha{
        Type: initValue,
        Count: initValue,
        Area: initValue,
        Place: initValue,
        AutoSaleThreshold: 4,
        Verbose: true,
    }
    return g
}