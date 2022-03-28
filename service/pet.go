package service

import (
	"anyiblog/model"
	"anyiblog/serializer"
	"anyiblog/serializer/params/api"
)

func AddMyPetService(param *api.AddMyPetParam, userId string) serializer.Response {
	petRes, isOk := model.CreatePet(param, userId)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "添加宠物成功",
			Data: petRes,
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "添加宠物失败",
		}
	}
}

// MyPetList 获取我的宠物列表
func MyPetList(userId string) serializer.Response {
	petList, isOk := model.GetPetList(userId)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "获取宠物信息成功",
			Data: petList,
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "暂无宠物信息",
		}
	}
}

// AddPetReserve 添加宠物预约
func AddPetReserve(param *api.AddPetReserveParam, userId, openId string) serializer.Response {
	isOk := model.AddPetReserve(param, userId)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "添加预约成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "添加预约失败",
		}
	}
}

// SelectPetReserve 查询宠物预约记录
func SelectPetReserve(petId, userId string, limit, page int) serializer.Response {
	petReserveRes, isOK := model.SelectReserveByQuery(map[string]interface{}{"pet_id": petId, "user_id": userId}, limit, page)
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
