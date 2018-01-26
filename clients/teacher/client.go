package teacher

import (
	"fmt"
	"github.com/icza/dyno"
	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/general"
	"github.com/mong0520/ChainChronicleGo/utils"
)

type Disciple struct {
	Beloved int `json:"beloved"`
	Card    struct {
		AddExp         int         `json:"add_exp"`
		Atk            int         `json:"atk"`
		CurrentWeapon  int         `json:"currentWeapon"`
		DispExp        int         `json:"disp_exp"`
		Exp            int         `json:"exp"`
		Flag           int         `json:"flag"`
		Hp             int         `json:"hp"`
		ID             int         `json:"id"`
		Idx            int         `json:"idx"`
		LimitBreak     int         `json:"limit_break"`
		Lv             int         `json:"lv"`
		MasterFlag     interface{} `json:"masterFlag"`
		Maxlv          int         `json:"maxlv"`
		NextExp        int         `json:"next_exp"`
		SellPrice      int         `json:"sellPrice"`
		Skillid        []int       `json:"skillid"`
		Type           int         `json:"type"`
		WeaponAttack   int         `json:"weaponAttack"`
		WeaponCritical int         `json:"weaponCritical"`
		WeaponGuard    int         `json:"weaponGuard"`
		WeaponReserve  []struct {
			WeaponAttack   int `json:"weaponAttack"`
			WeaponCritical int `json:"weaponCritical"`
			WeaponGuard    int `json:"weaponGuard"`
			Weaponid       int `json:"weaponid"`
		} `json:"weaponReserve"`
		Weaponid int `json:"weaponid"`
	} `json:"card"`
	Comment            interface{} `json:"comment"`
	EstablishedOnToday bool        `json:"established_on_today"`
	Lasttime           int         `json:"lasttime"`
	LoginTime          int         `json:"login_time"`
	Lv                 int         `json:"lv"`
	Name               string      `json:"name"`
	Resetable          bool        `json:"resetable"`
	Status             int         `json:"status"`
	UID                int         `json:"uid"`
}

func ListDisciple(metadata *clients.Metadata, param map[string]interface{}) (discipleList []Disciple) {
	api := "teacher/confirm_disciple"
	ret, _ := general.GeneralAction(api, metadata.Sid, param)
	//fmt.Println(utils.Map2JsonString(ret))

	if data, err := dyno.GetSlice(ret, "body", 0, "data"); err != nil {
		fmt.Println(err)
	} else {
		//fmt.Printf("%+v\n", data)
		for _, d := range data {
			item := &Disciple{}
			utils.Map2Struct(d.(map[string]interface{}), item)
			discipleList = append(discipleList, *item)
		}
	}
	//for _, d := range discipleList {
	//	fmt.Println(d.UID, d.Resetable, d.Name)
	//}

	return discipleList
}

func ResetDisciple(metadata *clients.Metadata, discipleId int) (resp map[string]interface{}, res int) {
    api := "teacher/reset_from_teacher"
    param := map[string]interface{}{
        "did": discipleId,
    }
    return general.GeneralAction(api, metadata.Sid, param)
}