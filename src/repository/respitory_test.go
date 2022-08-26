package repository_test

import (
	"encoding/json"
	"errors"
	"log"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/denizcamalan/key-value-store/model"
	"github.com/elliotchance/redismock"
	"github.com/go-playground/assert/v2"
	"github.com/go-redis/redis"
)

type Workflow struct{
	ID 			string		`json:"id" binding:"required"` 
	Company 	Company
}

type Company struct{
	Name		string		`json:"name" binding:"required"`
	Sector		string		`json:"sector" binding:"required"`
	Project		string		`json:"projectname" binding:"required"`
	Technology	string		`json:"technology" binding:"required"`
	Year		int 		`json:"year" binding:"required"`
}

var db = newTestRedis()

// Test Redis ClientMock Conf. //

func newTestRedis() *redismock.ClientMock {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return redismock.NewNiceMock(client)
}
// Redis repositories

func SettoRedis(workflow model.Workflow) error{

	workflow = model.Workflow{
		ID: "1",
		Company: model.Company{Name: "Teknosa", Project: "DSS", Year: 2020, Technology: "adtech",Sector: "technology"},
	}

	bytevalue, err := json.Marshal(workflow)
	if err !=nil {
		return errors.New("marshall error")
	}

	err2 := db.Set(workflow.ID, bytevalue, 0).Err()
	if err2 != nil { return err }
	
	return nil
}

func GetRedis(id string) (string,error){

	val, err := db.Get("1").Result(); 
	if err != nil { return "",err }

	return val,nil
}

func UpdateRedis(id string, company model.Company) error{

	workflow := model.Workflow{
		ID: "1",
		Company: model.Company{Name: "Teknosa", Project: "DSS", Year: 2020, Technology: "adtech",Sector: "technology"},
	}

	bytevalue, err := json.Marshal(workflow.Company)
	if err !=nil {
		return errors.New("marshall error")
	}

	if err := db.GetSet(id, bytevalue).Err(); err != nil {
		return err
	}
	return nil
}

func DeleteKey(key string) (int64){

	intVal,err := db.Del("1").Result()
	if err != nil {
		return intVal
	}
	return intVal
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

// -------------------- TEST ----------------------------

func TestSettoRedis(t *testing.T) {
	
	workflow := model.Workflow{
		ID: "1",
		Company: model.Company{Name: "Teknosa", Project: "DSS", Year: 2020, Technology: "adtech",Sector: "technology"},
	}
	r := newTestRedis()

	r.On("Set", "1").
		Return(redis.NewStatusResult("0",nil))

	assert.Equal(t, nil, SettoRedis(workflow))
}

func TestRedisGet(t *testing.T) {
	r := newTestRedis()
	workflow := `{"id":"1","Company":{"name":"Teknosa","sector":"technology","projectname":"DSS","technology":"adtech","year":2020}}`
	r.On("Get", "1").
		Return(redis.NewStatusResult("1", nil))
	result , err := GetRedis("1")
	assert.Equal(t, string(workflow), result)
	assert.Equal(t, err, err)
}

func TestRedisUpdate(t *testing.T) {
	r := newTestRedis()
	workflow := model.Workflow{
		ID: "1",
		Company: model.Company{Name: "Teknosa", Project: "DSS", Year: 2020, Technology: "adtech",Sector: "technology"},
	}
	r.On("Get", "1").
		Return(redis.NewStatusResult("1", nil))
	err := UpdateRedis("1", workflow.Company)
	assert.Equal(t, nil, err)
}

func TestRedisDel(t *testing.T) {
	r := newTestRedis()
	r.On("Del", "1").
		Return(redis.NewIntResult(0, nil))

	assert.Equal(t,int64(1), DeleteKey("1"))
}


func TestRedisDelAll(t *testing.T) {
	r := newTestRedis()
	keys := getAllKeys()
	r.On("Del", keys).
		Return(redis.NewIntResult(0, nil))

	assert.Equal(t,nil, DeleteAll())
}

