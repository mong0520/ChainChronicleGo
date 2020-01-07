package controllers

import (
	"fmt"
	"testing"
	"time"

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

func TestMyTest1(t *testing.T) {
	connStr := "mongodb://admin:gundam0079@cluster0-shard-00-00-krtk5.mongodb.net:27017,cluster0-shard-00-01-krtk5.mongodb.net:27017,cluster0-shard-00-02-krtk5.mongodb.net:27017/test?replicaSet=Cluster0-shard-0"
	dialInfo, err := mgo.ParseURL(connStr)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", dialInfo)
	dialInfo.Timeout = time.Second * 3
	dialInfo.Direct = true

	_, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		t.Fatal(err)
	}

}
func TestUpdateDB(t *testing.T) {
	connStr := "mongodb://admin:gundam0079@cluster0-shard-00-00-krtk5.mongodb.net:27017,cluster0-shard-00-01-krtk5.mongodb.net:27017,cluster0-shard-00-02-krtk5.mongodb.net:27017/test?replicaSet=Cluster0-shard-0"
	dialInfo, err := mgo.ParseURL(connStr)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", dialInfo)
	db, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		t.Fatal(err)
	}
	err = UpdateDB(db)
	if err != nil {
		t.Error(err)
	}
}
