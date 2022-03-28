package model

import (
	"anyiblog/conf"
	"anyiblog/serializer/params/api"
	"anyiblog/util"
	"fmt"
	"time"
)

// PetReserve 预约服务
type PetReserve struct {
	ReserveID        string    `gorm:"primary_key;column:reserve_id;type:char(36);not null" json:"reserve_id"`     // 主键ID
	PetID            string    `gorm:"column:pet_id;type:char(36);not null" json:"pet_id"`                         // 宠物ID
	ServiceProductID string    `gorm:"column:service_product_id;type:char(36);not null" json:"service_product_id"` // 服务ID
	ServiceItemID    string    `gorm:"column:service_item_id;type:char(36);not null" json:"service_item_id"`       // 服务项目id
	UserID           string    `gorm:"column:user_id;type:char(36);not null" json:"user_id"`                       // 用户ID
	ReserveTime      time.Time `gorm:"column:reserve_time;type:timestamp;not null" json:"reserve_time"`            // 预约时间
	ToDoorService    string    `gorm:"column:to_door_service;type:varchar(255);not null" json:"to_door_service"`   // 上门接送
	//订单状态（0待确认，1预约成功，2用户取消，3商家取消，4服务中，5已结单）
	Status       int       `gorm:"column:status;type:int;not null" json:"status"`
	ActionUser   string    `gorm:"column:action_user;type:char(10);not null" json:"action_user"`            // 订单操作人
	RealPayPrice float64   `gorm:"column:real_pay_price;type:decimal(10,2);not null" json:"real_pay_price"` // 实付价钱
	OtherService string    `gorm:"column:other_service;type:longtext" json:"other_service"`                 // 附加服务（改价用）
	OtherInfo    string    `gorm:"column:other_info;type:varchar(255)" json:"other_info"`                   // 备注信息
	StartTime    time.Time `gorm:"column:start_time;type:timestamp" json:"start_time"`                      // 开始时间
	EndTime      time.Time `gorm:"column:end_time;type:timestamp" json:"end_time"`                          // 结束时间
}

type OtherService struct {
	Name  string
	Price float64
}

type ReserveUserInfo struct {
	UserID string `json:"UserID"`
	Phone  string `json:"Phone"`
}

// ResPetReserveItem 预约信息
type ResPetReserveItem struct {
	UserInfo       ReserveUserInfo `json:"UserInfo"`
	ReserveID      string          `json:"ReserveID"`
	ServiceProduct string          `json:"ServiceProduct"`
	ServiceItem    string          `json:"ServiceItem"`
	PetInfo        Pet             `json:"PetInfo"`
	ReserveTime    string          `json:"ReserveTime"`
	ToDoorService  string          `json:"ToDoorService"`
	Status         int             `json:"Status"`
	ActionUser     string          `json:"ActionUser"`
	RealPayPrice   float64         `json:"RealPayPrice"`
	OtherService   string          `json:"OtherService"`
	OtherInfo      string          `json:"OtherInfo"`
}

// ResPetReserveList 预约服务列表
type ResPetReserveList struct {
	Count          int64               `json:"Count"`
	PetReserveList []ResPetReserveItem `json:"PetReserveList"`
}

// AddPetReserve 添加预约服务
func AddPetReserve(param *api.AddPetReserveParam, userId string) bool {
	reserveID := util.GenerateUUID()
	fmt.Println(param.ReserveTime)
	fmt.Println(util.StrTimeFormat(param.ReserveTime, "2006-01-02 15:04:05"))
	petReserve := PetReserve{
		ReserveID:        reserveID,
		PetID:            param.PetID,
		ServiceProductID: param.ServiceProductID,
		ServiceItemID:    param.ServiceItemID,
		UserID:           userId,
		ReserveTime:      util.StrTimeFormat(param.ReserveTime, "2006-01-02 15:04:05"),
		ToDoorService:    param.ToDoorService,
		Status:           0,
		OtherInfo:        param.OtherInfo,
		//ActionUser
		//RealPayPrice
		//OtherService
	}
	if conf.DB.Create(petReserve).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// GetResByQuery 根据查询字段名和查询参数获取结果
func GetResByQuery(Fields interface{}, QueryParams map[string]interface{}) interface{} {
	result := map[string]interface{}{}
	conf.DB.Model(PetReserve{}).Select(Fields).Where(QueryParams).First(&result)
	return result
}

// UpdateReserve 根据字段 更新预约记录
func UpdateReserve(QueryParams map[string]interface{}, ReserveID string) bool {
	i := conf.DB.Model(PetReserve{}).Where("reserve_id = ?", ReserveID).Updates(QueryParams)
	if i.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// SelectReserveByQuery 根据查询名 查询预约记录
func SelectReserveByQuery(QueryParams map[string]interface{}, limit, page int) (ResPetReserveList, bool) {
	var dataList []PetReserve
	var resPetReserveList ResPetReserveList
	i := conf.DB.Model(PetReserve{}).Where(QueryParams).Order("reserve_time asc").Limit(limit).Offset((page - 1) * limit).Find(&dataList).Count(&resPetReserveList.Count)
	if i.RowsAffected > 0 {
		for _, item := range dataList {
			petInfo, _ := GetPetInfo(item.PetID)
			serviceItemName := GetServiceItemInfo(item.ServiceItemID).ServiceItemName
			serviceProductName := GetServiceProductInfo(item.ServiceProductID).ServiceProductName
			itemRealPayPrice := GetServiceItemInfo(item.ServiceItemID).ServiceItemPrice
			userPhone := GetUserPhoneByUserId(item.UserID)
			resPetReserveList.PetReserveList = append(resPetReserveList.PetReserveList, ResPetReserveItem{
				UserInfo: ReserveUserInfo{
					UserID: item.UserID,
					Phone:  userPhone,
				},
				ReserveID:      item.ReserveID,
				ServiceProduct: serviceProductName,
				ServiceItem:    serviceItemName,
				PetInfo:        petInfo,
				ReserveTime:    item.ReserveTime.Format("2006-01-02 15:04:05"),
				ToDoorService:  item.ToDoorService,
				Status:         item.Status,
				ActionUser:     item.ActionUser,
				RealPayPrice:   itemRealPayPrice,
				OtherService:   item.OtherService,
				OtherInfo:      item.OtherInfo,
			})
		}
		return resPetReserveList, true
	} else {
		return ResPetReserveList{}, false
	}
}
