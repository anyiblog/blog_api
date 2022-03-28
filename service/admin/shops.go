package admin

import (
	"anyiblog/model"
	"anyiblog/serializer"
	adminParams "anyiblog/serializer/params/admin"
)

func CreateShopService(params *adminParams.CreateShopParams) serializer.Response {
	shopRes, isOk := model.CreateShop(params)
	if isOk {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "创建店铺成功",
			Data: shopRes,
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "创建店铺失败",
		}
	}
}
