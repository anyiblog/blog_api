package service

import (
	"anyiblog/model"
	"anyiblog/serializer"
)

func GetUserOrderListService(userId string, page, limit, queryType int) serializer.Response {
	res,hasData := model.QueryOrderBriefInfoByUserID(userId, page, limit, queryType)
	if hasData {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "获取成功",
			Data: res,
		}
	}else{
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "暂无更多数据",
		}
	}
}

func GetUserOrderDetailService(userId, OrderInfoID string) serializer.Response {
	res := model.QueryOrderDetailInfoByUserID(userId, OrderInfoID)
	return serializer.Response{
		Code: serializer.SystemOk,
		Msg:  "获取成功",
		Data: res,
	}
}
