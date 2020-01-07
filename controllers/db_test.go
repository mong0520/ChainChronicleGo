package controllers

import (
	"testing"

	"github.com/mong0520/ChainChronicleGo/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestGeneralQuery(t *testing.T) {
	query := bson.M{"name": "席爾瓦登場"}
	if db, err := mgo.Dial("localhost:27017"); err != nil {
		t.Error("Unable to connect DB", err)
	} else {
		quest := models.QuestDigest{}
		GeneralQuery(db, "questdigest", query, &quest)
		t.Log(quest.QuestID)
	}
}

func TestGetQuestByName(t *testing.T) {
	db, _ := mgo.Dial("localhost:27017")
	quest, err := GetQuestByName(db, "席爾瓦登場")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(quest.QuestID)
	}
}

func TestGetQuestsByName(t *testing.T) {
	db, _ := mgo.Dial("localhost:27017")
	// defaultSelector := &bson.M{
	// 	"questid":   1,
	// 	"questtype": 1,
	// 	"name":      1,
	// }
	quests, err := GetQuestsByName(db, "某個", nil)
	if err != nil {
		t.Error(err)
	} else {
		for _, q := range quests {
			t.Log(q)
		}
	}
}

func TestUpdateDB(t *testing.T) {
	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Fatal(err)
	}
	err = UpdateDB(db)
	if err != nil {
		t.Error(err)
	}
}
