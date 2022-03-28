package admin

import (
	"anyiblog/model"
	"anyiblog/serializer"
	"anyiblog/serializer/params/admin"
	"anyiblog/service"
	"anyiblog/util"
	"time"
)

// GetReserve 根据查询参数，获取预约记录
func GetReserve(param *admin.GetReserveParam) serializer.Response {
	queryParams := make(map[string]interface{})
	if len(param.ServiceProductID) > 0 {
		queryParams["service_product_id"] = param.ServiceProductID
	}
	if len(param.ToDoorService) > 0 {
		queryParams["to_door_service"] = param.ToDoorService
	}
	if len(param.Status) > 0 {
		queryParams["status"] = param.Status
	}
	if param.Limit == 0 {
		param.Limit = 10
	}
	if param.Page == 0 {
		param.Page = 1
	}
	petReserveRes, isOK := model.SelectReserveByQuery(queryParams, param.Limit, param.Page)
	if isOK {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "获取预约记录成功",
			Data: petReserveRes,
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "暂无更多记录",
		}
	}
}

// GetAllServiceProductService 获取所有服务项目
func GetAllServiceProductService() serializer.Response {
	res := model.GetAllServiceProduct()
	return serializer.Response{
		Code: serializer.SystemOk,
		Msg:  "获取成功",
		Data: res,
	}
}

// GetAllServiceItemInfo 获取所有服务子项目
func GetAllServiceItemInfo(serviceProductId string) serializer.Response {
	res := model.GetAllServiceItemInfo(serviceProductId)
	return serializer.Response{
		Code: serializer.SystemOk,
		Msg:  "获取成功",
		Data: res,
	}
}

// AddServiceProductService 添加服务类别
func AddServiceProductService(serviceProductName string) serializer.Response {
	currentRes, isOk := model.AddServiceProduct(serviceProductName)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "添加成功",
			Data: currentRes,
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "添加失败",
		}
	}
}

// AddServiceItemService 添加服务类别子项目
func AddServiceItemService(param *admin.AddServiceItemParam) serializer.Response {
	res, isOk := model.AddServiceItem(param)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "添加成功",
			Data: res,
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "添加失败",
		}
	}
}

// DeleteServiceItemService 添加服务类别子项目
func DeleteServiceItemService(ServiceItemServiceID string) serializer.Response {
	isOk := model.DeleteServiceItem(ServiceItemServiceID)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "删除成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "删除失败",
		}
	}
}

// AcceptReserve 接受预约订单
func AcceptReserve(param *admin.AcceptReserveParam) serializer.Response {
	data := make(map[string]interface{})
	data["status"] = 1 //接受订单
	isOk := model.UpdateReserve(data, param.ReserveID)
	if isOk {
		fields := make([]string, 0)
		fields = append(fields, "pet_id", "service_product_id", "reserve_time")
		queryParam := make(map[string]interface{})
		queryParam["reserve_id"] = param.ReserveID
		reserveRes := model.GetResByQuery(fields, queryParam).(map[string]interface{})
		openId := model.GetOpenIdByUserId(param.UserID)
		service.PushReserveStatus(openId, reserveRes["pet_id"].(string), reserveRes["service_product_id"].(string), util.TimeFormat(reserveRes["reserve_time"].(time.Time)), "预约成功", "预约成功，请您在预约时间内保持电话通畅")
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "接受预订成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "接受预订失败",
		}
	}
}

// CancelReserve 取消预约订单
func CancelReserve(reserveId string) serializer.Response {
	data := make(map[string]interface{})
	data["status"] = 3 //商家取消
	isOk := model.UpdateReserve(data, reserveId)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "取消成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "取消失败",
		}
	}
}

// endTime2TextTime 结束时间转文本时间，返回相差小时数
func endTime2TextTime(str string) int {
	nowTime := time.Now()
	nowHour := nowTime.Hour() + 1 //当前小时
	actionTime := util.StrTimeFormat("2006-01-02 15:04", str)
	actionHour := actionTime.Hour() + 1
	return actionHour - nowHour
}

// ActionReserve 开始服务预约订单，并下发微信通知
func ActionReserve(param *admin.ActionReserveParam) serializer.Response {
	data := make(map[string]interface{})
	data["action_user"] = param.ActionUser
	data["status"] = 4 // 订单正在服务中
	data["start_time"] = util.NowTime()
	data["end_time"] = util.StrTimeFormat("2006-01-02 15:04", param.EndTime)
	isOk := model.UpdateReserve(data, param.ReserveID)
	if isOk {
		fields := make([]string, 0)
		fields = append(fields, "pet_id", "service_product_id")
		queryParam := make(map[string]interface{})
		queryParam["reserve_id"] = param.ReserveID
		reserveRes := model.GetResByQuery(fields, queryParam).(map[string]interface{})
		openId := model.GetOpenIdByUserId(param.UserID)
		service.PushReserveStart(openId, reserveRes["pet_id"].(string), reserveRes["service_product_id"].(string), param.EndTime)
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "操作成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "操作失败，请刷新重试",
		}
	}
}

// EndReserve 结束预约订单
func EndReserve(reserveId, userId string) serializer.Response {
	data := make(map[string]interface{})
	data["status"] = 5 // 结单
	isOk := model.UpdateReserve(data, reserveId)
	if isOk {
		fields := make([]string, 0)
		fields = append(fields, "pet_id", "service_product_id")
		queryParam := make(map[string]interface{})
		queryParam["reserve_id"] = reserveId
		reserveRes := model.GetResByQuery(fields, queryParam).(map[string]interface{})
		openId := model.GetOpenIdByUserId(userId)
		service.PushReserveEnd(openId, reserveRes["pet_id"].(string), reserveRes["service_product_id"].(string), util.TimeFormat(util.NowTime()))
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "结单成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "结单失败",
		}
	}
}

// UpdateReserve 更新预约订单信息
func UpdateReserve(param *admin.UpdateReserveParam) serializer.Response {
	data := make(map[string]interface{})
	if len(param.ReserveID) > 0 && len(param.ServiceProductID) > 0 && len(param.ServiceItemID) > 0 && len(param.ReserveTime) > 0 && len(param.ToDoorService) > 0 && param.RealPayPrice > 0 {
		data["reserve_id"] = param.ReserveID
		data["service_product_id"] = param.ServiceProductID
		data["service_item_id"] = param.ServiceItemID
		data["reserve_time"] = param.ReserveTime
		data["to_door_service"] = param.ToDoorService
		data["real_pay_price"] = param.RealPayPrice
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "信息不合法",
		}
	}
	isOk := model.UpdateReserve(data, param.ReserveID)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "更新信息成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "更新信息失败",
		}
	}
}
