package model

import (
	"anyiblog/conf"
	"anyiblog/serializer/params/admin"
	"anyiblog/util"
)

// ServiceItem 服务项目
type ServiceItem struct {
	ServiceItemID    string  `gorm:"primary_key;column:service_item_id;type:char(36);not null" json:"service_item_id"`
	ServiceProductID string  `gorm:"column:service_product_id;type:char(36)" json:"service_product_id"`             // 服务ID
	ServiceItemName  string  `gorm:"column:service_item_name;type:varchar(255);not null" json:"service_item_name"`    // 服务名称
	ServiceItemPrice float64 `gorm:"column:service_item_price;type:decimal(10,2);not null" json:"service_item_price"` // 服务价格
}

func GetServiceItemInfo(serviceItemID string) ServiceItem {
	var resInfo ServiceItem
	conf.DB.Model(ServiceItem{}).Where("service_item_id = ?", serviceItemID).Take(&resInfo)
	return resInfo
}

func GetAllServiceItemInfo(serviceProductID string) []ServiceItem {
	var resInfos []ServiceItem
	conf.DB.Model(ServiceItem{}).Where("service_product_id = ?", serviceProductID).Find(&resInfos)
	return resInfos
}

func AddServiceItem(param *admin.AddServiceItemParam) (ServiceItem, bool) {
	serviceItemID := util.GenerateUUID()
	i := ServiceItem{
		ServiceItemID:    serviceItemID,
		ServiceProductID: param.ServiceProductID,
		ServiceItemName:  param.ServiceItemName,
		ServiceItemPrice: param.ServiceItemPrice,
	}
	if conf.DB.Create(i).RowsAffected > 0 {
		return i, true
	} else {
		return i, false
	}
}

func DeleteServiceItem(ServiceItemServiceID string) bool {
	i := conf.DB.Where("service_item_id = ?", ServiceItemServiceID).Delete(ServiceItem{})
	if i.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
