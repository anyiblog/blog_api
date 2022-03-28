package admin

// GetReserveParam 获取预约服务
type GetReserveParam struct {
	ServiceProductID string `json:"ServiceProductID"`
	ToDoorService    string `json:"ToDoorService"`
	Status           string `json:"Status"`
	Limit            int    `json:"Limit"`
	Page             int    `json:"Page"`
}
// AddServiceProductParam 新增服务类别
type AddServiceProductParam struct {
	ServiceProductName string `json:"ServiceProductName"`
}
// AddServiceItemParam 新增服务项
type AddServiceItemParam struct {
	ServiceProductID string  `json:"ServiceProductID"`
	ServiceItemName  string  `json:"ServiceItemName"`
	ServiceItemPrice float64 `json:"ServiceItemPrice"`
}
// DeleteServiceItemParam 删除服务项
type DeleteServiceItemParam struct {
	ServiceItemID string `json:"ServiceItemID"`
}
// AcceptReserveParam 接受预约服务
type AcceptReserveParam struct {
	ReserveID string `json:"ReserveID"` // 预约ID
	UserID    string `json:"UserID"`    // 用户ID
}
// ActionReserveParam 开始服务预约订单
type ActionReserveParam struct {
	ReserveID  string `json:"ReserveID"`  // 预约ID
	UserID     string `json:"UserID"`     // 用户ID
	ActionUser string `json:"ActionUser"` // 操作人
	EndTime    string `json:"EndTime"`    // 预计服务结束时间
}
// UpdateReserveParam 更新预约服务
type UpdateReserveParam struct {
	ReserveID        string  `json:"ReserveID"`        // 预约ID
	ServiceProductID string  `json:"ServiceProductID"` // 服务ID
	ServiceItemID    string  `json:"ServiceItemID"`    // 服务项目id
	ReserveTime      string  `json:"ReserveTime"`      // 预约时间
	ToDoorService    string  `json:"ToDoorService"`    // 上门接送
	RealPayPrice     float64 `json:"RealPayPrice"`     // 实付价钱
}