package api

type SetAddressParams struct {
	UserName     string `json:"userName" binding:"required"`
	TelNumber    string `json:"telNumber" binding:"required"`
	ProvinceName string `json:"ProvinceName" binding:"required"` //国标收货地址第一级地址
	CityName     string `json:"cityName" binding:"required"`     // 国标收货地址第二级地址
	CountyName   string `json:"countyName" binding:"required"`   //国标收货地址第三级地址
	DetailInfo   string `json:"detailInfo" binding:"required"`   //详细收货地址信息
}
