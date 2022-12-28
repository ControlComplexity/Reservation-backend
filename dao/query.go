package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reservation/github.com/reservation/api"
)

func Init() *gorm.DB {
	fmt.Println("init gorm db...")
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/reservation?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}

func GetSystemInfo() (*api.GetSystemInfoResp, error) {
	return &api.GetSystemInfoResp{
		ErrorCode: "0",
	}, nil
}

func QueryActivityList() (*api.QueryActivityListResp, error) {
	db := Init()
	var activities []api.Activity
	fmt.Println("db2:", db)
	db.Model(&api.Activity{}).Find(&activities)
	fmt.Println("activities:", activities)
	activitySlice := make([]*api.Activity, 0)
	for _, ac := range activities {
		activitySlice = append(activitySlice, &ac)
	}
	return &api.QueryActivityListResp{
		Data: &api.QueryActivityListResp_Data{
			ActivityList: activitySlice,
		},
		ErrorCode: "0",
	}, nil
}

func QueryActivityListByDay(req *api.QueryActivityListByDayReq) (*api.QueryActivityListByDayResp, error) {
	db := Init()
	var activities []api.Activity
	fmt.Println("db2:", db)
	db.Model(&api.Activity{}).Where("time < ? ", req.Day).Where("time > ? ", req.Day).Find(&activities)
	fmt.Println("activities:", activities)
	activitySlice := make([]*api.Activity, 0)
	for _, ac := range activities {
		activitySlice = append(activitySlice, &ac)
	}
	return &api.QueryActivityListByDayResp{
		Data: &api.QueryActivityListByDayResp_Data{
			ActivityList: activitySlice,
		},
		ErrorCode: "0",
	}, nil
}
