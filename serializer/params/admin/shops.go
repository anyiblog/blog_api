package admin

type CreateShopParams struct {
	ShopName         string `json:"shop_name"`          // 门店名称
	ShopLogo         string `json:"shop_logo"`          // 图片id
	ShopAddress      string `json:"shop_address"`       // 门店地址
	ShopContactName  string `json:"shop_contact_name"`  // 门店联系人
	ShopContactPhone string `json:"shop_contact_phone"` // 门店联系电话
	ShopTime         string `json:"shop_time"`          // 营业时间
	ShopOtherInfo    string `json:"shop_other_info"`    // 备注信息
}
