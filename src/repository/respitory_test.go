package repository_test

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
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


func SettoRedis(workflow string) error{

	var db = newTestRedis()

	err2 := db.Set("1", "starts", time.Duration(time.Now().Second())).Err()
	if err2 != nil { return	err2 }
	
	return nil
}

func GetfromRedis(id string) string{

	var db = newTestRedis()

	val, err := db.Get("key1").Result(); 
	if err != nil { return "" }

	return val
}

func DeleteKey(key string) (int64){

	var db = newTestRedis()

	intVal,err := db.Del("key1").Result()
	if err != nil {
		return intVal
	}
	return intVal
}
func TestSettoRedis(t *testing.T) {

	strs := "value1"
	var value error = nil

	r := newTestRedis()
	r.On("Set", "1",strs,time.Duration(time.Now().Second())).
		Return(redis.NewCmdResult(value, nil))

	assert.Equal(t,	value, SettoRedis(strs))
}

func TestRedisMGet(t *testing.T) {
	r := newTestRedis()
	strs := "value1"
	var value string
	r.On("Get", "key1").
		Return(redis.NewCmdResult(value, nil))

	assert.Equal(t, value, GetfromRedis(strs))
}

func TestRedisDel(t *testing.T) {
	r := newTestRedis()
	r.On("Del", "key1").
		Return(redis.NewIntResult(0, nil))

	assert.Equal(t,int64(0), DeleteKey("key1"))
}



// func TestSet(t *testing.T) {
// 	Init()
// 	exp := time.Duration(0)

// 	mock := redismock.NewNiceMock(client)
// 	mock.On("Set", key, val, exp).Return(redis.NewStatusResult("", nil))

// 	p1 :=model.Person{
// 		UID:      "123456789",
// 		Name:     "Juan",
// 		LastName: "Orjuela",
// 	}

// 	personService = services.New(mock, &logger)
// 	err := personService.Set(p1)
// 	assert.NoError(t, err)
// }

// func TestGet(t *testing.T) {
// 	Init()
// 	mock := redismock.NewNiceMock(client)
// 	mock.On("Get", key).Return(redis.NewStringResult(string(val), nil))

// 	personService = services.New(mock, &logger)
// 	p2, err := personService.Get(key)
// 	assert.NoError(t, err)

// 	assert.Equal(t, "123456789", p2.UID)
// 	assert.Equal(t, "Juan", p2.Name)
// 	assert.Equal(t, "Orjuela", p2.LastName)
// }