package dal

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reservation/constant"
	"reservation/github.com/reservation/api"
	"reservation/model"
	"reservation/utils"
	"strings"
)

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(101.43.39.188:3306)/reservation?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}

func GetSystemInfo() (*api.GetSystemInfoResp, error) {
	return &api.GetSystemInfoResp{
		ErrorCode: constant.SUCCESS_ERROR_CODE,
		Success:   true,
	}, nil
}

func QueryActivityInfo(req *api.QueryActivityInfoReq) (*api.QueryActivityInfoResp, error) {
	db := Init()
	var ac model.ActivityDO
	db.Model(&model.ActivityDO{}).Where("activity_id = ? ", req.Id).Find(&ac).Limit(1)
	fmt.Println("ac:", ac)

	activity := api.Activity{
		Id:          ac.Id,
		Name:        ac.Name,
		Label:       strings.Split(ac.Label, ","),
		Price:       ac.Price,
		Time:        utils.Time2String(ac.Time),
		Location:    ac.Location,
		SmallImg:    ac.SmallImg,
		BigImg:      ac.BigImg,
		Description: ac.Description,
		CreatedAt:   utils.Time2Milli(ac.CreatedAt),
		UpdatedAt:   utils.Time2Milli(ac.UpdatedAt),
	}
	//activities := make([]*api.Activity, 0)
	//activities = append(activities, &activity)
	return &api.QueryActivityInfoResp{
		Data:      &activity,
		Success:   true,
		ErrorCode: constant.SUCCESS_ERROR_CODE,
	}, nil
}

func JoinActivity(req *api.JoinActivityReq) (*api.JoinActivityResp, error) {
	db := Init()
	var ac model.ActivityDO
	e := db.Model(&model.ActivityDO{}).Where("activity_id = ? ", req.Id).Find(&ac).Limit(1)
	fmt.Println("ac:", ac)
	if e.Error != nil {
		return &api.JoinActivityResp{
			Success:   false,
			ErrorCode: constant.RECORD_NOT_FOUND,
			ErrorMsg:  "failed to join activity, " + e.Error.Error(),
		}, nil
	}
	return &api.JoinActivityResp{
		Price:     ac.Price,
		Success:   true,
		ErrorCode: constant.SUCCESS_ERROR_CODE,
	}, nil
}

func QueryActivityList() (*api.QueryActivityListResp, error) {
	db := Init()
	var activities []model.ActivityDO

	db.Model(&model.ActivityDO{}).Find(&activities)
	fmt.Println("activities:", activities)
	activitySlice := make([]*api.Activity, 0)
	for _, ac := range activities {

		activitySlice = append(activitySlice, &api.Activity{
			Id:          ac.Id,
			Name:        ac.Name,
			Label:       strings.Split(ac.Label, ","),
			Price:       ac.Price,
			Time:        utils.Time2String(ac.Time),
			Location:    ac.Location,
			SmallImg:    ac.SmallImg,
			BigImg:      ac.BigImg,
			Description: ac.Description,
			CreatedAt:   utils.Time2Milli(ac.CreatedAt),
			UpdatedAt:   utils.Time2Milli(ac.UpdatedAt),
		})
	}
	return &api.QueryActivityListResp{
		Data: &api.QueryActivityListResp_Data{
			ActivityList: activitySlice,
		},
		Success:   true,
		ErrorCode: constant.SUCCESS_ERROR_CODE,
	}, nil
}

//
//func QueryActivityListByDay(req *api.QueryActivityListByDayReq) (*api.QueryActivityListByDayResp, error) {
//	db := Init()
//	var activities []model.ActivityDO
//
//	db.Model(&model.ActivityDO{}).Where("time < ? ", req.Day).Where("time > ? ", req.Day).Find(&activities)
//	fmt.Println("activities:", activities)
//	activitySlice := make([]*api.Activity, 0)
//	for _, ac := range activities {
//		activitySlice = append(activitySlice, &api.Activity{
//			Name:  ac.Name,
//			Price: ac.Price,
//		})
//	}
//	return &api.QueryActivityListByDayResp{
//		Data: &api.QueryActivityListByDayResp_Data{
//			ActivityList: activitySlice,
//		},
//		Success:   true,
//		ErrorCode: constant.SUCCESS_ERROR_CODE,
//	}, nil
//}

// QueryOrderList 查询订单
func QueryOrderList(req *api.QueryOrderListReq) (*api.QueryOrderListResp, error) {
	db := Init()
	var orders []api.Order

	db.Model(&api.Order{}).Find(&orders)
	fmt.Println("orders:", orders)
	orderSlice := make([]*api.Order, 0)
	for _, or := range orders {
		orderSlice = append(orderSlice, &or)
	}
	return &api.QueryOrderListResp{
		Data: &api.QueryOrderListResp_Data{
			OrderList: orderSlice,
		},
		ErrorCode: constant.SUCCESS_ERROR_CODE,
		Success:   true,
	}, nil
}

// QuerySwipers 查询订单
func QuerySwipers(req *api.QuerySwipersReq) (*api.QuerySwipersResp, error) {
	db := Init()
	var swipers []model.SwiperDO
	db.Model(&model.SwiperDO{}).Find(&swipers)
	fmt.Println("swipers:", swipers)
	sw := make([]*api.Swiper, 0)
	for _, ac := range swipers {
		sw = append(sw, &api.Swiper{
			Url: ac.URL,
			Img: ac.Img,
		})
	}
	return &api.QuerySwipersResp{
		Data: &api.QuerySwipersResp_Data{
			SwiperList: sw,
		},
		Success:   true,
		ErrorCode: constant.SUCCESS_ERROR_CODE,
	}, nil
}
