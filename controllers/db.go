package controllers

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "fmt"
    "github.com/mong0520/ChainChronicleGo/models"
)

func QueryCharaByField(session *mgo.Session, field string, value string, card *models.Charainfo){
    query(session, "charainfo", "title", value, card)
}

func GeneralQuery(session *mgo.Session, collection string, query interface{}, result interface{})(err error){
    if err := session.DB("cc").C(collection).Find(query).One(result) ; err != nil{
        fmt.Println(err)
        return err
    }else{
        return nil
    }
}


func query(session *mgo.Session, collection string, field string, value string, result *models.Charainfo){
    if err := session.DB("cc").C(collection).Find(bson.M{field: value}).One(result) ; err != nil{
        fmt.Println(err)
    }else{
        fmt.Println(result.Name)
    }
}