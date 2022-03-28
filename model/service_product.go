package model

import (
	"anyiblog/conf"
	"anyiblog/util"
)

// ServiceProduct [...]
type ServiceProduct struct {
	ServiceProductID   string `gorm:"primary_key;column:service_product_id;type:char(36);not null" json:"service_product_id"`
	ServiceProductName string `gorm:"column:service_product_name;type:varchar(255);not null" json:"service_product_name"` // 服务类目名称
}

// GetServiceProductInfo 获取某条服务的信息
func GetServiceProductInfo(serviceProductID string) ServiceProduct {
	var resInfo ServiceProduct
	conf.DB.Model(ServiceProduct{}).Where("service_product_id = ?", serviceProductID).Take(&resInfo)
	return resInfo
}

// GetAllServiceProduct 获取所有服务
func GetAllServiceProduct() []ServiceProduct {
	var resInfo []ServiceProduct
	conf.DB.Model(ServiceProduct{}).Find(&resInfo)
	return resInfo
}

// AddServiceProduct 添加服务类别
func AddServiceProduct(ServiceProductName string) (ServiceProduct, bool) {
	serviceProductID := util.GenerateUUID()
	i := ServiceProduct{
		ServiceProductID:   serviceProductID,
		ServiceProductName: ServiceProductName,
	}
	if conf.DB.Create(i).RowsAffected > 0 {
		return i, true
	} else {
		return i, false
	}
}
