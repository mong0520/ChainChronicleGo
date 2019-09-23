package models

//for mongodb

var mapping = map[string]interface{}{
	"reinforce":   &Reinforce{},
	"evolve":      &Evolve{},
	"weaponlist":  &Weapon{},
	"charainfo":   &Charainfo{},
	"chararein":   &Chararein{},
	"questdigest": &QuestDigest{},
	"skilllist":   &Skilllist{},
}

type Charainfo struct {
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

type QuestDigest struct {
	AppendSkill struct {
	} `json:"append_skill"`
	AreaID      int   `json:"area_id"`
	ChapterList []int `json:"chapter_list"`
	Difficulty  int   `json:"difficulty"`
	Flag        struct {
	} `json:"flag"`
	Kind           int           `json:"kind"`
	KindPrm        int           `json:"kind_prm"`
	Name           string        `json:"name"`
	PlaceID        int           `json:"place_id"`
	QuestID        int           `json:"quest_id"`
	Stamina        int           `json:"stamina"`
	BgList         []int         `json:"bg_list"`
	DifficultyList []interface{} `json:"difficulty_list"`
	ChapterCnt     int           `json:"chapter_cnt"`
}

type Skilllist struct {
	Ability   int           `json:"ability"`
	Flag      int           `json:"flag"`
	Flag00    int           `json:"flag0_0"`
	Flag01    int           `json:"flag0_1"`
	Flag10    int           `json:"flag1_0"`
	Flag11    int           `json:"flag1_1"`
	Flavor    string        `json:"flavor"`
	IParam0   int           `json:"iParam0"`
	IParam1   int           `json:"iParam1"`
	IconType  int           `json:"icon_type"`
	Name      string        `json:"name"`
	Param0    int           `json:"param0"`
	Param1    int           `json:"param1"`
	Param2    int           `json:"param2"`
	Param3    int           `json:"param3"`
	Param4    int           `json:"param4"`
	Param5    int           `json:"param5"`
	Param6    int           `json:"param6"`
	Param7    int           `json:"param7"`
	Param8    int           `json:"param8"`
	Param9    int           `json:"param9"`
	Skillid   int           `json:"skillid"`
	Str0      string        `json:"str0"`
	Str1      string        `json:"str1"`
	Str2      string        `json:"str2"`
	Str3      string        `json:"str3"`
	Str4      string        `json:"str4"`
	Str5      string        `json:"str5"`
	Text      string        `json:"text"`
	Timestamp string        `json:"timestamp"`
	Sub       []interface{} `json:"sub"`
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

type Weapon struct {
	AttackMax   int    `json:"attackMax"`
	CriticalMax int    `json:"criticalMax"`
	EquipType   int    `json:"equip_type"`
	GuardMax    int    `json:"guardMax"`
	ID          int    `json:"id"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	Rank        int    `json:"rank"`
	Skill       int    `json:"skill"`
	Type        int    `json:"type"`
	TypeAtk     int    `json:"type_atk"`
	TypeCri     int    `json:"type_cri"`
	TypeGrd     int    `json:"type_grd"`
}

type Reinforce struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Profile     string  `json:"profile"`
	RankLimit   int     `json:"rank_limit"`
	Rarity      int     `json:"rarity"`
	Ring        int     `json:"ring"`
	SuccessRate float64 `json:"success_rate"`
	Type        int     `json:"type"`
}

type Chararein struct {
	Exp     int    `json:"exp"`
	ID      int    `json:"id"`
	Jobtype int    `json:"jobtype"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
	Rarity  int    `json:"rarity"`
	Ring    int    `json:"ring"`
}

type GachaResultItem struct {
	Cnt    int `json:"cnt"`
	ItemID int `json:"item_id"`
	UID    int `json:"uid"`
}

type GachaResultWeapon struct {
	Cnt       int `json:"cnt"`
	ItemID    int `json:"item_id"`
	UID       int `json:"uid"`
	Timestamp int `json:"timestamp"`
}

type GachaResult interface {
	dummy() int
}

type AllData struct {
	Body []struct {
		Data interface{} `json:"data"`
		Proc int         `json:"proc"`
		Type int         `json:"type"`
	} `json:"body"`
	FinishedQuest []struct {
		QuestID         int         `json:"quest_id"`
		Timestamp       int         `json:"timestamp"`
		Type            int         `json:"type"`
		Chapter         int         `json:"chapter,omitempty"`
		ContinueCount   int         `json:"continue_count,omitempty"`
		EpisodeID       interface{} `json:"episode_id,omitempty"`
		EpisodeIdx      interface{} `json:"episode_idx,omitempty"`
		EpisodeQuestIdx interface{} `json:"episode_quest_idx,omitempty"`
		EventCnt        int         `json:"event_cnt,omitempty"`
		Finished        bool        `json:"finished,omitempty"`
		TimeRecord      int         `json:"time_record,omitempty"`
		UID             int         `json:"uid,omitempty"`
		V3              bool        `json:"v3,omitempty"`
		Wave            int         `json:"wave,omitempty"`
		WinCnt          int         `json:"win_cnt,omitempty"`
	} `json:"finished_quest"`
	MVer struct {
		Battleinfo                int `json:"battleinfo"`
		Booklist                  int `json:"booklist"`
		Bossinfo                  int `json:"bossinfo"`
		Casinoinfo                int `json:"casinoinfo"`
		Charainfo                 int `json:"charainfo"`
		Chropowerinfo             int `json:"chropowerinfo"`
		Coverinfo                 int `json:"coverinfo"`
		Denju                     int `json:"denju"`
		Enemyinfo                 int `json:"enemyinfo"`
		Episodeinfo               int `json:"episodeinfo"`
		EventPortal               int `json:"eventPortal"`
		Eventareainfo             int `json:"eventareainfo"`
		Eventbadge                int `json:"eventbadge"`
		Eventplace                int `json:"eventplace"`
		Explorerinfo              int `json:"explorerinfo"`
		Explorerlocation          int `json:"explorerlocation"`
		Gachainfo                 int `json:"gachainfo"`
		Heroicskilllist           int `json:"heroicskilllist"`
		Homelist                  int `json:"homelist"`
		Itemlist                  int `json:"itemlist"`
		Medalinfo                 int `json:"medalinfo"`
		Missionlist               int `json:"missionlist"`
		Motionlist                int `json:"motionlist"`
		Openseainfo               int `json:"openseainfo"`
		Questdigest               int `json:"questdigest"`
		Questroute                int `json:"questroute"`
		Questtree                 int `json:"questtree"`
		Reload                    int `json:"reload"`
		Repeatmissionlist         int `json:"repeatmissionlist"`
		Servertime                int `json:"servertime"`
		Skilllist                 int `json:"skilllist"`
		Sortieinfo                int `json:"sortieinfo"`
		SoundExData               int `json:"soundExData"`
		Spdrama                   int `json:"spdrama"`
		Speedruninfo              int `json:"speedruninfo"`
		Subjugationexpeditioninfo int `json:"subjugationexpeditioninfo"`
		Supporterskill            int `json:"supporterskill"`
		Teacherdisciple           int `json:"teacherdisciple"`
		Towerinfo                 int `json:"towerinfo"`
		Uzuinfo                   int `json:"uzuinfo"`
		Various                   int `json:"various"`
		Vignette                  int `json:"vignette"`
		Weaponcompose             int `json:"weaponcompose"`
		Weaponcomposeevent        int `json:"weaponcomposeevent"`
		Weaponlist                int `json:"weaponlist"`
		Worldinfo                 int `json:"worldinfo"`
	} `json:"mVer"`
	Res int `json:"res"`
}

type QuestResponse struct {
	Body []struct {
		Data []map[string]interface{} `json:"data"`
		Proc int                      `json:"proc"`
		Type int                      `json:"type"`
	} `json:"body"`
	DiffBeloved []struct {
		SrcCid int `json:"src_cid"`
		Val    int `json:"val"`
	} `json:"diff_beloved"`
	Earns struct {
		Bonus struct {
			D bool `json:"d"`
			S bool `json:"s"`
			T bool `json:"t"`
		} `json:"bonus"`
		BonusGold       int  `json:"bonus_gold"`
		CharaExp        int  `json:"chara_exp"`
		Exp             int  `json:"exp"`
		Gold            int  `json:"gold"`
		IsComebackExpUp bool `json:"is_comeback_exp_up"`
		IsNewcomerExpUp bool `json:"is_newcomer_exp_up"`
		Point           int  `json:"point"`
		Treasure        []struct {
			EnemyIdx    int    `json:"enemyIdx"`
			ID          int    `json:"id"`
			Treasure    int    `json:"treasure"`
			TreasureIdx int    `json:"treasureIdx"`
			Type        string `json:"type"`
			Val         int    `json:"val"`
		} `json:"treasure"`
	} `json:"earns"`
	MVer struct {
		Battleinfo                int `json:"battleinfo"`
		Booklist                  int `json:"booklist"`
		Bossinfo                  int `json:"bossinfo"`
		Casinoinfo                int `json:"casinoinfo"`
		Charainfo                 int `json:"charainfo"`
		Chropowerinfo             int `json:"chropowerinfo"`
		Coverinfo                 int `json:"coverinfo"`
		Denju                     int `json:"denju"`
		Enemyinfo                 int `json:"enemyinfo"`
		Episodeinfo               int `json:"episodeinfo"`
		Eventareainfo             int `json:"eventareainfo"`
		Eventbadge                int `json:"eventbadge"`
		Eventplace                int `json:"eventplace"`
		Explorerinfo              int `json:"explorerinfo"`
		Explorerlocation          int `json:"explorerlocation"`
		Gachainfo                 int `json:"gachainfo"`
		Heroicskilllist           int `json:"heroicskilllist"`
		Homelist                  int `json:"homelist"`
		Itemlist                  int `json:"itemlist"`
		Missionlist               int `json:"missionlist"`
		Motionlist                int `json:"motionlist"`
		Openseainfo               int `json:"openseainfo"`
		Questdigest               int `json:"questdigest"`
		Questroute                int `json:"questroute"`
		Questtree                 int `json:"questtree"`
		Reload                    int `json:"reload"`
		Repeatmissionlist         int `json:"repeatmissionlist"`
		Servertime                int `json:"servertime"`
		Skilllist                 int `json:"skilllist"`
		Sortieinfo                int `json:"sortieinfo"`
		SoundExData               int `json:"soundExData"`
		Spdrama                   int `json:"spdrama"`
		Speedruninfo              int `json:"speedruninfo"`
		Subjugationexpeditioninfo int `json:"subjugationexpeditioninfo"`
		Supporterskill            int `json:"supporterskill"`
		Teacherdisciple           int `json:"teacherdisciple"`
		Towerinfo                 int `json:"towerinfo"`
		Uzuinfo                   int `json:"uzuinfo"`
		Various                   int `json:"various"`
		Vignette                  int `json:"vignette"`
		Weaponcompose             int `json:"weaponcompose"`
		Weaponcomposeevent        int `json:"weaponcomposeevent"`
		Weaponlist                int `json:"weaponlist"`
		Worldinfo                 int `json:"worldinfo"`
	} `json:"mVer"`
	Res int `json:"res"`
}

type EndRaidResponse struct {
	Body []interface{} `json:"body"`
	MVer struct {
		Battleinfo                int `json:"battleinfo"`
		Booklist                  int `json:"booklist"`
		Bossinfo                  int `json:"bossinfo"`
		Casinoinfo                int `json:"casinoinfo"`
		Charainfo                 int `json:"charainfo"`
		Chropowerinfo             int `json:"chropowerinfo"`
		Coverinfo                 int `json:"coverinfo"`
		Denju                     int `json:"denju"`
		Enemyinfo                 int `json:"enemyinfo"`
		Episodeinfo               int `json:"episodeinfo"`
		EventPortal               int `json:"eventPortal"`
		Eventareainfo             int `json:"eventareainfo"`
		Eventbadge                int `json:"eventbadge"`
		Eventplace                int `json:"eventplace"`
		Explorerinfo              int `json:"explorerinfo"`
		Explorerlocation          int `json:"explorerlocation"`
		Gachainfo                 int `json:"gachainfo"`
		Heroicskilllist           int `json:"heroicskilllist"`
		Homelist                  int `json:"homelist"`
		Itemlist                  int `json:"itemlist"`
		Medalinfo                 int `json:"medalinfo"`
		Missionlist               int `json:"missionlist"`
		Motionlist                int `json:"motionlist"`
		Openseainfo               int `json:"openseainfo"`
		Questdigest               int `json:"questdigest"`
		Questroute                int `json:"questroute"`
		Questtree                 int `json:"questtree"`
		Reload                    int `json:"reload"`
		Repeatmissionlist         int `json:"repeatmissionlist"`
		Servertime                int `json:"servertime"`
		Skilllist                 int `json:"skilllist"`
		Sortieinfo                int `json:"sortieinfo"`
		SoundExData               int `json:"soundExData"`
		Spdrama                   int `json:"spdrama"`
		Speedruninfo              int `json:"speedruninfo"`
		Subjugationexpeditioninfo int `json:"subjugationexpeditioninfo"`
		Supporterskill            int `json:"supporterskill"`
		Teacherdisciple           int `json:"teacherdisciple"`
		Towerinfo                 int `json:"towerinfo"`
		Uzuinfo                   int `json:"uzuinfo"`
		Various                   int `json:"various"`
		Vignette                  int `json:"vignette"`
		Weaponcompose             int `json:"weaponcompose"`
		Weaponcomposeevent        int `json:"weaponcomposeevent"`
		Weaponlist                int `json:"weaponlist"`
		Worldinfo                 int `json:"worldinfo"`
	} `json:"mVer"`
	Res   int `json:"res"`
	Earns struct {
		Bonus struct {
			D bool `json:"d"`
			S bool `json:"s"`
			T bool `json:"t"`
		} `json:"bonus"`
		BonusGold       int  `json:"bonus_gold"`
		CharaExp        int  `json:"chara_exp"`
		Exp             int  `json:"exp"`
		Gold            int  `json:"gold"`
		IsComebackExpUp bool `json:"is_comeback_exp_up"`
		IsNewcomerExpUp bool `json:"is_newcomer_exp_up"`
		Point           int  `json:"point"`
		Treasure        []struct {
			EnemyIdx    int    `json:"enemyIdx"`
			ID          int    `json:"id"`
			Treasure    int    `json:"treasure"`
			TreasureIdx int    `json:"treasureIdx"`
			Type        string `json:"type"`
			Val         int    `json:"val"`
		} `json:"treasure"`
	} `json:"earns"`
}

type EndRaidCardResponse struct {
	Data []struct {
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
	} `json:"data"`
	Proc int `json:"proc"`
	Type int `json:"type"`
}

func (g GachaResultChara) dummy() int {
	return 0
}

func (g GachaResultItem) dummy() int {
	return 0
}

func GetStruct(s string) (m interface{}) {
	//mapping := map[string]interface{}{
	//    "reinforce": &Reinforce{},
	//    "evolve": &Evolve{},
	//    "weaponlist": &Weapon{},
	//    "charainfo": &Charainfo{},
	//    "chararein": &Chararein{},
	//    "questdigest": &QuestDigest{},
	//    "skilllist": &Skilllist{},
	//}
	if _, ok := mapping[s]; ok {
		return mapping[s]
	} else {
		return nil
	}
}

type Book struct {
	Body []struct {
		Data struct {
			Cover  interface{} `json:"cover"`
			Expup  interface{} `json:"expup"`
			Global struct {
				Purchase    int `json:"purchase"`
				RaidItem    int `json:"raid_item"`
				Review      int `json:"review"`
				TmpInterval int `json:"tmp_interval"`
			} `json:"global"`
			PartyOpt struct {
				CostBorder1 int `json:"cost_border1"`
				CostBorder2 int `json:"cost_border2"`
				Type        int `json:"type"`
			} `json:"party_opt"`
			Raid          interface{}   `json:"raid"`
			Subjugation   []interface{} `json:"subjugation"`
			Totalwar      interface{}   `json:"totalwar"`
			Weaponcompose struct {
				WeaponComposeDisabled bool `json:"weapon_compose_disabled"`
			} `json:"weaponcompose"`
		} `json:"data"`
		Proc int `json:"proc"`
		Type int `json:"type"`
	} `json:"body"`
	CardList []struct {
		Cid       int   `json:"cid"`
		Skillid   []int `json:"skillid"`
		Quest     []int `json:"quest,omitempty"`
		Expansion []struct {
			ID    int `json:"id"`
			LvIdx int `json:"lv_idx"`
		} `json:"expansion,omitempty"`
	} `json:"card_list"`
	CoverList []struct {
		CoverID int    `json:"cover_id"`
		Name    string `json:"name"`
		Point   int    `json:"point"`
		Rank    int    `json:"rank"`
		Text    string `json:"text"`
	} `json:"cover_list"`
	DramaList     []int `json:"drama_list"`
	FinishedQuest []struct {
		QuestID int `json:"quest_id"`
		Type    int `json:"type"`
	} `json:"finished_quest"`
	MVer struct {
		Battleinfo                int `json:"battleinfo"`
		Booklist                  int `json:"booklist"`
		Bossinfo                  int `json:"bossinfo"`
		Casinoinfo                int `json:"casinoinfo"`
		Charainfo                 int `json:"charainfo"`
		Chropowerinfo             int `json:"chropowerinfo"`
		Coverinfo                 int `json:"coverinfo"`
		Denju                     int `json:"denju"`
		Enemyinfo                 int `json:"enemyinfo"`
		Episodeinfo               int `json:"episodeinfo"`
		EventPortal               int `json:"eventPortal"`
		Eventareainfo             int `json:"eventareainfo"`
		Eventbadge                int `json:"eventbadge"`
		Eventplace                int `json:"eventplace"`
		Explorerinfo              int `json:"explorerinfo"`
		Explorerlocation          int `json:"explorerlocation"`
		Gachainfo                 int `json:"gachainfo"`
		Heroicskilllist           int `json:"heroicskilllist"`
		Homelist                  int `json:"homelist"`
		Itemlist                  int `json:"itemlist"`
		Medalinfo                 int `json:"medalinfo"`
		Missionlist               int `json:"missionlist"`
		Motionlist                int `json:"motionlist"`
		Openseainfo               int `json:"openseainfo"`
		Questdigest               int `json:"questdigest"`
		Questroute                int `json:"questroute"`
		Questtree                 int `json:"questtree"`
		Reload                    int `json:"reload"`
		Repeatmissionlist         int `json:"repeatmissionlist"`
		Servertime                int `json:"servertime"`
		Skilllist                 int `json:"skilllist"`
		Sortieinfo                int `json:"sortieinfo"`
		SoundExData               int `json:"soundExData"`
		Spdrama                   int `json:"spdrama"`
		Speedruninfo              int `json:"speedruninfo"`
		Spquestinfo               int `json:"spquestinfo"`
		Subjugationexpeditioninfo int `json:"subjugationexpeditioninfo"`
		Supporterskill            int `json:"supporterskill"`
		Teacherdisciple           int `json:"teacherdisciple"`
		Towerinfo                 int `json:"towerinfo"`
		Uzuinfo                   int `json:"uzuinfo"`
		Various                   int `json:"various"`
		Vignette                  int `json:"vignette"`
		Weaponcompose             int `json:"weaponcompose"`
		Weaponcomposeevent        int `json:"weaponcomposeevent"`
		Weaponlist                int `json:"weaponlist"`
		Worldinfo                 int `json:"worldinfo"`
	} `json:"mVer"`
	Res int `json:"res"`
}
