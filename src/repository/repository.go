package repository

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/denizcamalan/key-value-store/configuration"
	"github.com/denizcamalan/key-value-store/model"
)

var db = configuration.NewDatabase()

func SettoRedis(workflow model.Workflow) error{

	bytevalue, err := json.Marshal(workflow)
	if err !=nil {
		return errors.New("marshall error")
	}

	err2 := db.Set(workflow.ID, bytevalue, 0).Err()
	if err2 != nil { return err }
	
	return nil
}

func GetRedis(id string) (string,error){
	val, err := db.Get(id).Result(); 
	if err != nil { return "",err }

	return val,nil
}

func CheckRedis(id string) bool{
	err := db.Get(id).Err()
    if err == nil {
        return false
    }else {
		return true	
	}
}

func DeleteKey(key string) error{

	err := db.Del(key).Err()
	
	if err != nil {
		return err
	}
	return nil
}


func ListfromRedis() ([]model.Workflow, error){
	
	var workflows []model.Workflow
	var workflow  model.Workflow

	ids:= getAllKeys()
	
	for _,id := range ids{
		if id != ""{
			if strData, err := GetRedis(id); err != nil { 
				return nil,errors.New("getfromRedis no value") 
			}else {
				json.Unmarshal([]byte(strData),&workflow)
				workflows = append(workflows, workflow)	
			}
		}else {
			log.Println("nil id")
			continue
		}
	}
	return workflows,nil
}

func UpdateRedis(id string, company model.Company) error{
	var workflow model.Workflow
	workflow.Company = company

	bytevalue, err := json.Marshal(workflow.Company)
	if err !=nil {
		return errors.New("marshall error")
	}

	if err := db.GetSet(id, bytevalue).Err(); err != nil {
		return err
	}
	return nil
}

func DeleteAll() error{

	ids:= getAllKeys()

	for _,id := range ids{
		if err := db.Del(id).Err(); err != nil {
			return err
		}
	}
	return nil
}

// get all created keys
func getAllKeys() (keys []string){
	iter := db.Scan(0, "*", 0).Iterator()
	for iter.Next() {
		keys = append(keys, iter.Val())
	}
	return
}
