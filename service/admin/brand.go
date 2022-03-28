package admin

import (
	"anyiblog/model"
	"anyiblog/serializer"
	"anyiblog/serializer/params/admin"
)

// GetAllBrandService 获取所有品牌
func GetAllBrandService(params *admin.GetAllBrandParams) serializer.Response {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}
	return serializer.Response{
		Code: serializer.SystemOk,
		Msg:  "获取成功",
		Data: model.GetAllBrand(params.Limit, params.Page),
	}
}
