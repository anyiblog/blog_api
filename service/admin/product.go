package admin

import (
	"anyiblog/model"
	"anyiblog/serializer"
	adminParams "anyiblog/serializer/params/admin"
	"github.com/tidwall/gjson"
)

// AddProductService 添加产品
func AddProductService(params *adminParams.AddProductParam) serializer.Response {
	if model.AddGoodsSpu(params) {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "新增商品成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemError,
			Msg:  "新增商品失败",
		}
	}
}

// GetAllProductService 获取所有产品
func GetAllProductService(params *adminParams.GetAllProductParam) serializer.Response {
	queryParams := make(map[string]interface{}) // 定值查询
	likeParams := make(map[string]interface{})  // 模糊查询
	if len(params.ProductName) > 0 {
		likeParams["spu_name"] = params.ProductName
	}
	if len(params.BrandName) > 0 { // 根据品牌名获取对应ID
		getBrandIdStatus, brandId := model.GetBrandID(params.BrandName)
		if getBrandIdStatus {
			queryParams["brand_id"] = brandId
		}
	}
	if params.Status >= 0 {
		queryParams["spu_status"] = params.Status
	}
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}
	return serializer.Response{
		Code: serializer.SystemOk,
		Msg:  "获取成功",
		Data: model.GetAllGoods(queryParams, likeParams, params.Limit, params.Page,true),
	}
}

// ProductOnService 上架产品
func ProductOnService(productIDArray string) serializer.Response {
	jsonArr := gjson.Parse(productIDArray)
	var counter []int
	for _, v := range jsonArr.Array() {
		if model.SetProductStatus(v.String(), 0) { // 0 上架
			counter = append(counter, 1) // 状态计数器
		}
	}
	if len(counter) == len(jsonArr.Array()) {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "操作成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "操作错误",
		}
	}
}

// ProductOffService 下架产品
func ProductOffService(productIDArray string) serializer.Response {
	jsonArr := gjson.Parse(productIDArray)
	var counter []int
	for _, v := range jsonArr.Array() {
		if model.SetProductStatus(v.String(), 1) { // 1 下架
			counter = append(counter, 1) // 状态计数器
		}
	}
	if len(counter) == len(jsonArr.Array()) {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "操作成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "操作错误",
		}
	}
}

// DeleteProductService 删除产品
func DeleteProductService(productIDArray string) serializer.Response {
	jsonArr := gjson.Parse(productIDArray)
	var counter []int
	for _, v := range jsonArr.Array() {
		if model.DeleteProduct(v.String()) { // 1 下架
			counter = append(counter, 1) // 状态计数器
		}
	}
	if len(counter) == len(jsonArr.Array()) {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "删除成功",
		}
	} else {
		return serializer.Response{
			Code: serializer.SystemOk,
			Msg:  "删除错误",
		}
	}
}
