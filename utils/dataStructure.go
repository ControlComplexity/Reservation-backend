package utils

import (
	"reservation/github.com/reservation/api"
	"reservation/model"
	"strings"
)

func CvtDataStructure(input interface{}) interface{} {
	switch input.(type) {
	case model.ActivityDO:
		ac := input.(model.ActivityDO)
		return api.Activity{
			Id:          ac.Id,
			Name:        ac.Name,
			Label:       strings.Split(ac.Label, ","),
			Price:       ac.Price,
			Time:        Time2String(ac.Time),
			Location:    ac.Location,
			SmallImg:    ac.SmallImg,
			BigImg:      ac.BigImg,
			Description: ac.Description,
			CreatedAt:   Time2Milli(ac.CreatedAt),
			UpdatedAt:   Time2Milli(ac.UpdatedAt),
		}
	case model.UserDO:
		user := input.(model.UserDO)
		return api.User{
			Id:              user.Id,
			HeadImage:       user.HeadImage,
			NickName:        user.NickName,
			BirthDate:       user.BirthDate,
			Gender:          api.Gender(api.Gender_value[user.Gender]),
			Height:          user.Height,
			Weight:          user.Weight,
			Hometown:        user.Hometown,
			Location:        user.Location,
			EmotionalStatus: api.EmotionalStatus(api.EmotionalStatus_value[user.EmotionalStatus]),
			Education:       api.Education(api.Education_value[user.Education]),
			Occupation:      user.Occupation,
			University:      user.University,
			WechatNumber:    user.WechatNumber,
			PhoneNumber:     user.PhoneNumber,
		}
	default:
	}
	return nil
}
