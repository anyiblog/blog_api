package model

import (
	"anyiblog/conf"
	"fmt"
	"time"
)

// WechatPay 微信支付记录表
type WechatPay struct {
	WechatPayID string    `gorm:"primary_key;column:wechat_pay_id;type:varchar(255);not null" json:"wechat_pay_id"` // 微信支付流水订单号
	OrderNo     string    `gorm:"column:order_no;type:char(36);not null" json:"order_no"`                           // 订单号
	PayAmount   float64   `gorm:"column:pay_amount;type:decimal(10,2);not null" json:"pay_amount"`                  // 微信回调后，微信实际支付总金额
	IsSubscribe string    `gorm:"column:is_subscribe;type:varchar(255);not null" json:"is_subscribe"`               // 支付后是否关注公众号
	UserOpenID  string    `gorm:"column:user_open_id;type:varchar(128);not null" json:"user_open_id"`               // 用户唯一ID
	PayTimeEnd  time.Time `gorm:"column:pay_time_end;type:timestamp" json:"pay_time_end"`                           // 支付结束时间
}

func CreateWechatPay(params WechatPay) {
	conf.DB.Create(params)
}

// QueryOrderNotPay 查询订单未支付  有记录 false  无记录代表未支付 true
func QueryOrderNotPay(WechatPayID, OrderNo string) bool {
	db := conf.DB.Model(WechatPay{}).Where("wechat_pay_id = ? AND order_no = ?", WechatPayID, OrderNo)
	fmt.Println(db.RowsAffected)
	if db.RowsAffected > 0 {
		return false
	} else {
		return true
	}
}

func GetPayIDByOrderNo(orderNo string) string {
	var WechatPayID struct {
		WechatPayID string
	}
	conf.DB.Model(WechatPay{}).Where("order_no = ?", orderNo).Find(&WechatPayID)
	return WechatPayID.WechatPayID
}
