package api

import (
	"anyiblog/model"
	"anyiblog/serializer"
	"anyiblog/serializer/params/api"
)

// SetAddressService 添加地址
func SetAddressService(params api.SetAddressParams, userId string) serializer.Response {
	address := model.CreateAddress(params, userId)
	return serializer.Response{
		Code: serializer.SystemOk,
		Msg:  "设置成功",
		Data: address,
	}
}

// GetDefaultAddressService 获取默认地址
func GetDefaultAddressService(userId string) serializer.Response {
	address := model.GetDefaultAddress(userId)
	return serializer.Response{
		Code: serializer.SystemOk,
		Msg:  "获取成功",
		Data: address,
	}
}

// SetDefaultAddressService 设置默认地址
func SetDefaultAddressService(userId, addressId string) serializer.Response {
	isOk := model.SetDefaultAddress(userId, addressId)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "设置成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "设置失败",
		}
	}
}
