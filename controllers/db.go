package controllers

import (
	"fmt"
	"reflect"

	"github.com/icza/dyno"
	"github.com/mong0520/ChainChronicleGo/clients"
	"github.com/mong0520/ChainChronicleGo/clients/general"
	"github.com/mong0520/ChainChronicleGo/models"
	"github.com/mong0520/ChainChronicleGo/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//func QueryCharaByField(session *mgo.Session, field string, value string, card *models.Charainfo){
//    query(session, "charainfo", "title", value, card)
//}

func GeneralQuery(session *mgo.Session, collection string, query interface{}, result interface{}) (err error) {
	if err := session.DB("cc").C(collection).Find(query).One(result); err != nil {
		return err
	} else {
		return nil
	}
}

func GeneralQueryAll(session *mgo.Session, collection string, query interface{}, result interface{}) (err error) {
	if err := session.DB("cc").C(collection).Find(query).All(result); err != nil {
		return err
	}

	return nil
}

func GetQuestByName(session *mgo.Session, name string) (*models.QuestDigest, error) {
	result := models.QuestDigest{}
	query := bson.M{"name": name}
	err := GeneralQuery(session, "questdigest", query, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetQuestsByName(session *mgo.Session, name string) ([]models.QuestDigest, error) {
	results := []models.QuestDigest{}
	// queryString := fmt.Sprintf(`/%s/`, name)
	// query := bson.M{"name": bson.RegEx{Pattern: queryString, Options: ""}}
	query := bson.M{
		"name": bson.RegEx{Pattern: name, Options: "i"},
	}
	err := session.DB("cc").C("questdigest").Find(query).All(&results)

	if err != nil {
		return nil, err
	}

	return results, nil
}

//
//
//func query(session *mgo.Session, collection string, field string, value string, result *models.Charainfo){
//    if err := session.DB("cc").C(collection).Find(bson.M{field: value}).One(result) ; err != nil{
//        fmt.Println(err)
//    }else{
//        fmt.Println(result.Name)
//    }
//}

func UpdateDB(metadata *clients.Metadata) {
	session := metadata.DB
	apiMapping := map[string][]string{
		"data/questdigest": {"questdigest"},
		"data/charainfo":   {"charainfo", "chararein"},
		"data/weaponlist":  {"evolve", "reinforce", "weaponlist"},
		"data/skilllist":   {"skilllist"},
	}
	for api, fields := range apiMapping {
		fmt.Println("### Updating DB from...", api)
		param := map[string]interface{}{}
		ret, _ := general.GeneralAction(api, metadata.Sid, param)
		//fmt.Println(utils.Map2JsonString(ret))
		//break
		//fields = []string{"charainfo", "chararein"}
		for _, field := range fields {
			fmt.Println("Updating collection", field)
			dataList, _ := dyno.GetSlice(ret, field)
			session.DB("cc").C(field).DropCollection()

			for _, data := range dataList {
				tmpEnt := models.GetStruct(field)
				switch reflect.TypeOf(data).Kind() {
				case reflect.Map:
					utils.Map2Struct(data.(map[string]interface{}), tmpEnt)
					session.DB("cc").C(field).Insert(&tmpEnt)
				case reflect.Slice:
					for _, item := range data.([]interface{}) {
						utils.Map2Struct(item.(map[string]interface{}), tmpEnt)
						session.DB("cc").C(field).Insert(&tmpEnt)
					}
				}

			}
		}
	}
}
