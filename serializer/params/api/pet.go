package api

type AddMyPetParam struct {
	PetName          string  `binding:"required" json:"PetName"`          // 姓名
	PetAvatar        string  `binding:"required" json:"PetAvatar"`        // 头像
	PetVariety       string  `binding:"required" json:"PetVariety"`       // 品种
	PetBirthday      string  `binding:"required" json:"PetBirthday"`      // 生日
	PetAge           int     `binding:"required" json:"PetAge"`           // 年龄
	PetWeight        float64 `binding:"required" json:"PetWeight"`        // 体重
	PetGender        string  `binding:"required" json:"PetGender"`        // 性别
	PetSterilization int     `binding:"required" json:"PetSterilization"` // 0 绝育，1未绝育
	PetHair          int     `binding:"required" json:"PetHair"`          // 0 无毛 1 短毛 2长毛
}

type AddPetReserveParam struct {
	PetID            string `binding:"required" json:"PetID"`            // 宠物ID
	ServiceProductID string `binding:"required" json:"ServiceProductID"` // 服务ID
	ServiceItemID    string `binding:"required" json:"ServiceItemID"`    // 服务项目id
	ReserveTime      string `binding:"required" json:"ReserveTime"`      // 预约时间
	ToDoorService    string `binding:"required" json:"ToDoorService"`    // 上门接送
	//Status           int       // 订单状态（0待后台确认，1预约成功，2预约取消，3，商家取消，4服务完成已结单）
	//ActionUser       string    // 订单操作人
	//RealPayPrice     float64   // 实付价钱
	//OtherService     string    // 附加服务（改价用）
	OtherInfo string `json:"otherInfo"` // 备注信息
}

type SelectPetReserveParam struct {
	PetID string `binding:"required" json:"PetID"`
	Limit int    `binding:"required" json:"Limit"`
	Page  int    `binding:"required" json:"Page"`
}

type PushReserveParam struct {
	TemplateID string `binding:"required" json:"TemplateID"`
}
