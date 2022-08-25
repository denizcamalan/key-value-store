package repository

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/denizcamalan/key-value-store/configuration"
	"github.com/denizcamalan/key-value-store/model"
)

var db = configuration.NewDatabase()

func SettoRedis(workflow model.Workflow) error{

	bytevalue, err := json.Marshal(workflow)
	if err !=nil {
		return errors.New("marshall error")
	}

	err2 := db.Set(workflow.ID, bytevalue, time.Duration(time.Now().Second())).Err()
	if err2 != nil { return err }
	
	return nil
}

func GetfromRedis(id string) (string,error){
	val, err := db.Get(id).Result(); 
	if err != nil { return "",err }

	return val,nil
}

func CheckIfExist(id string) bool{
	err := db.Get(id).Err()
    if err == nil {
        return false
    }else {
		return true	
	}
}

func SetByIdRedis(id string, company model.Company) error{

	bytevalue, err := json.Marshal(company)
	if err !=nil {
		return errors.New("marshall error")
	}

	if err := db.Set(id, bytevalue, time.Duration(time.Now().Second())).Err(); err != nil {
		return err
	}
	return nil
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

	ids:= GetAllKeys()
	
	for _,id := range ids{
		if id != ""{
			if strData, err := GetfromRedis(id); err != nil { 
				return nil,errors.New("getfromRedis no value") 
			}else {
				json.Unmarshal([]byte(strData),&workflow)
				workflows = append(workflows, workflow)	
			}
		}else {
			log.Println("nil id")
		}
	}
	return workflows,nil
}

func UpdateRedis(id string, company model.Company) error{

	bytevalue, err := json.Marshal(company)
	if err !=nil {
		return errors.New("marshall error")
	}

	if err := db.GetSet(id, bytevalue).Err(); err != nil {
		return err
	}
	return nil
}

func DeleteAll() error{

	ids:= GetAllKeys()

	for _,id := range ids{
		if err := db.Del(id).Err(); err != nil {
			return err
		}
	}
	return nil
}

// get all created keys
func GetAllKeys() (keys []string){
	iter := db.Scan(0, "*", 0).Iterator()
	for iter.Next() {
		keys = append(keys, iter.Val())
	}
	return
}
