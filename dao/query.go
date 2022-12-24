package dao

import (
	"reservation/github.com/reservation/api"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB
func init(){
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/reservation?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	defer db.Close()
}

func GetSystemInfo() (*api.GetSystemInfoResp, error) {
	return &api.GetSystemInfoResp{
		ErrorCode: "0",
	}, nil
}

func QueryActivityList() (*api.QueryActivityListResp, error) {
	return &api.QueryActivityListResp{
		ErrorCode: "0",
	}, nil
}

func QueryActivityListByDay() (*api.QueryActivityListByDayResp, error) {
	return &api.QueryActivityListByDayResp{
		ErrorCode: "0",
	}, nil
}
