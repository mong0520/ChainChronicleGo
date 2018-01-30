package models

//for mongodb
type Charainfo struct {
    ID struct {
        Oid string `json:"$oid"`
    } `json:"_id"`
    Iparam1      int     `json:"iparam1"`
    Title        string  `json:"title"`
    DenjuFrom    int     `json:"denju_from"`
    Defbonus     int     `json:"defbonus"`
    Growth       int     `json:"growth"`
    Skillparam8  int     `json:"skillparam8"`
    Pattern7     int     `json:"pattern7"`
    Pattern8     int     `json:"pattern8"`
    Revision     int     `json:"revision"`
    Movspeed     float64 `json:"movspeed"`
    LimitBreak   int     `json:"limit_break"`
    Pattern6     int     `json:"pattern6"`
    Skilltext    string  `json:"skilltext"`
    Profile      string  `json:"profile"`
    Iniap        int     `json:"iniap"`
    Skillflag10  int     `json:"skillflag1_0"`
    Weakflag     int     `json:"weakflag"`
    Skillparam2  int     `json:"skillparam2"`
    Jobtype      int     `json:"jobtype"`
    Str4         string  `json:"str4"`
    Skillparam5  int     `json:"skillparam5"`
    Movetype     int     `json:"movetype"`
    Str5         string  `json:"str5"`
    Bgm          int     `json:"bgm"`
    DispQuest    int     `json:"disp_quest"`
    Guard        int     `json:"guard"`
    L1Voice      int     `json:"l1_voice"`
    DVoice       int     `json:"d_voice"`
    Battletype   int     `json:"battletype"`
    Atkrange     float64 `json:"atkrange"`
    Iparam0      int     `json:"iparam0"`
    Rarity       int     `json:"rarity"`
    Spmotname    string  `json:"spmotname"`
    SkillCost    int     `json:"skill_cost"`
    Skillflag01  int     `json:"skillflag0_1"`
    Home         int     `json:"home"`
    Aparam       int     `json:"Aparam"`
    Skillparam3  float64 `json:"skillparam3"`
    CRInumber    int     `json:"CRInumber"`
    Sup1         int     `json:"sup1"`
    Skillparam1  int     `json:"skillparam1"`
    Skillparam9  int     `json:"skillparam9"`
    Weaponid     int     `json:"weaponid"`
    Str1         string  `json:"str1"`
    Isv2         int     `json:"isv2"`
    Skillname    string  `json:"skillname"`
    Str0         string  `json:"str0"`
    Weaponid2    int     `json:"weaponid2"`
    Atknumber    int     `json:"atknumber"`
    DenjuTo      int     `json:"denju_to"`
    Dsptext      string  `json:"dsptext"`
    Skillparam6  int     `json:"skillparam6"`
    Str2         string  `json:"str2"`
    Defphysics   int     `json:"defphysics"`
    Illustrator  string  `json:"illustrator"`
    Name         string  `json:"name"`
    HpOffset     int     `json:"hp_offset"`
    Effectid     int     `json:"effectid"`
    Modelscale   float64 `json:"modelscale"`
    Chargemotion string  `json:"chargemotion"`
    Skillid      []int   `json:"skillid"`
    Str3         string  `json:"str3"`
    Shotspeed    int     `json:"shotspeed"`
    Place        int     `json:"place"`
    BuddyID      int     `json:"buddyId"`
    Maxlv        int     `json:"maxlv"`
    Critical     int     `json:"critical"`
    FirstQuest   int     `json:"first_quest"`
    Ring         int     `json:"ring"`
    Pattern1     int     `json:"pattern1"`
    MeetChara    int     `json:"meet_chara"`
    Skillparam0  int     `json:"skillparam0"`
    Attackflag   int     `json:"attackflag"`
    VVoice       int     `json:"v_voice"`
    Pattern3     int     `json:"pattern3"`
    L2Voice      int     `json:"l2_voice"`
    Cid          int     `json:"cid"`
    Charge       int     `json:"charge"`
    Skillflag11  int     `json:"skillflag1_1"`
    Gender       int     `json:"gender"`
    Skillflag00  int     `json:"skillflag0_0"`
    Bodyrange    float64 `json:"bodyrange"`
    Cost         int     `json:"cost"`
    CharaType    int     `json:"chara_type"`
    VMotionid    int     `json:"v_motionid"`
    Skillparam7  int     `json:"skillparam7"`
    Pattern4     int     `json:"pattern4"`
    Atkspeed     int     `json:"atkspeed"`
    SAbility     int     `json:"s_ability"`
    Pattern2     int     `json:"pattern2"`
    Defmagic     int     `json:"defmagic"`
    Sptext       string  `json:"sptext"`
    Pattern5     int     `json:"pattern5"`
    CRIrevision  int     `json:"CRIrevision"`
    Motionid     int     `json:"motionid"`
    VoiceArtist  string  `json:"voice_artist"`
    SupCost      int     `json:"sup_cost"`
    Sup2         int     `json:"sup2"`
    ExpType      int     `json:"exp_type"`
    Inihp        int     `json:"inihp"`
    MasterFlag   int     `json:"master_flag"`
    Pattern0     int     `json:"pattern0"`
    Skillparam4  int     `json:"skillparam4"`
}

// for all data API
type CharaData struct {
    Atk            int           `json:"atk"`
    CurrentWeapon  int           `json:"currentWeapon"`
    DispExp        int           `json:"disp_exp"`
    Exp            int           `json:"exp"`
    Flag           int           `json:"flag"`
    Hp             int           `json:"hp"`
    ID             int           `json:"id"`
    Idx            int           `json:"idx"`
    LimitBreak     int           `json:"limit_break"`
    Lv             int           `json:"lv"`
    MasterFlag     interface{}   `json:"masterFlag"`
    Maxlv          int           `json:"maxlv"`
    NextExp        int           `json:"next_exp"`
    SellPrice      int           `json:"sellPrice"`
    Skillid        []interface{} `json:"skillid"`
    Type           int           `json:"type"`
    WeaponAttack   int           `json:"weaponAttack"`
    WeaponCritical int           `json:"weaponCritical"`
    WeaponGuard    int           `json:"weaponGuard"`
    WeaponReserve  []struct {
        WeaponAttack   int `json:"weaponAttack"`
        WeaponCritical int `json:"weaponCritical"`
        WeaponGuard    int `json:"weaponGuard"`
        Weaponid       int `json:"weaponid"`
    } `json:"weaponReserve"`
    Weaponid int `json:"weaponid"`
}

type GachaResultChara struct {
    Atk            int   `json:"atk"`
    CurrentWeapon  int   `json:"currentWeapon"`
    DispExp        int   `json:"disp_exp"`
    Exp            int   `json:"exp"`
    Flag           int   `json:"flag"`
    Hp             int   `json:"hp"`
    ID             int   `json:"id"`
    Idx            int   `json:"idx"`
    IsNew          bool  `json:"is_new"`
    LimitBreak     int   `json:"limit_break"`
    Lv             int   `json:"lv"`
    Maxlv          int   `json:"maxlv"`
    NextExp        int   `json:"next_exp"`
    SellPrice      int   `json:"sellPrice"`
    Skillid        []int `json:"skillid"`
    Type           int   `json:"type"`
    WeaponAttack   int   `json:"weaponAttack"`
    WeaponCritical int   `json:"weaponCritical"`
    WeaponGuard    int   `json:"weaponGuard"`
    WeaponReserve  []struct {
        WeaponAttack   int `json:"weaponAttack"`
        WeaponCritical int `json:"weaponCritical"`
        WeaponGuard    int `json:"weaponGuard"`
        Weaponid       int `json:"weaponid"`
    } `json:"weaponReserve"`
    Weaponid int `json:"weaponid"`
}

type Evolve struct {
    EvolDst   int    `json:"evol_dst"`
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Profile   string `json:"profile"`
    RankLimit int    `json:"rank_limit"`
    Rarity    int    `json:"rarity"`
    Usable0   int    `json:"usable0"`
    Usable1   int    `json:"usable1"`
    Usable2   int    `json:"usable2"`
    Ring      int    `json:"ring"`
    Material  int    `json:"material"`
}

type GachaResultItem struct {
    Cnt    int `json:"cnt"`
    ItemID int `json:"item_id"`
    UID    int `json:"uid"`
}

type GachaResult interface {
    dummy() int
}


func (g GachaResultChara) dummy() int {
    return 0
}

func (g GachaResultItem) dummy() int {
    return 0
}