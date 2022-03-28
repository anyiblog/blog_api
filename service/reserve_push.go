package service

import (
	"anyiblog/model"
	"anyiblog/util"
	"fmt"
	"os"
)

// PushReserveStatus 推送预约服务状态通知
func PushReserveStatus(openId, PetId, ServiceProductID, ReserveTime, statusText, otherInfo string) {
	wxAccessToken := util.GetAccessToken()
	petInfo, _ := model.GetPetInfo(PetId)
	serviceProduct := model.GetServiceProductInfo(ServiceProductID)
	reserveData := util.ReserveStatusData{
		Thing2:  util.Value{Value: petInfo.PetName},
		Thing3:  util.Value{Value: serviceProduct.ServiceProductName},
		Time4:   util.Value{Value: ReserveTime},
		Phrase5: util.Value{Value: statusText},
		//"请等待店主接单，并在预订时间内保持您的电话通畅"
		Thing6: util.Value{Value: otherInfo},
	}
	util.WechatSendReserveStatusService(wxAccessToken, openId, reserveData)
}

// PushReserveStart 推送预约服务状态通知
func PushReserveStart(openId, PetId, ServiceProductID, EndTime string) {
	wxAccessToken := util.GetAccessToken()
	petInfo, _ := model.GetPetInfo(PetId)
	serviceProduct := model.GetServiceProductInfo(ServiceProductID)
	reserveData := util.ReserveStartData{
		Thing2: util.Value{Value: petInfo.PetName},
		Thing4: util.Value{Value: serviceProduct.ServiceProductName},
		Time3:  util.Value{Value: util.TimeFormat(util.NowTime())},
		Thing1: util.Value{Value: os.Getenv("ShopName")},
		Thing5: util.Value{Value: fmt.Sprintf("结束时间%s",EndTime)},
	}
	util.WechatSendReserveStartService(wxAccessToken, openId, reserveData)
}

// PushReserveEnd 推送预约服务状态通知
func PushReserveEnd(openId, PetId, ServiceProductID, EndTime string) {
	wxAccessToken := util.GetAccessToken()
	petInfo, _ := model.GetPetInfo(PetId)
	serviceProduct := model.GetServiceProductInfo(ServiceProductID)
	reserveData := util.ReserveEndData{
		Thing2: util.Value{Value: petInfo.PetName},
		Thing3: util.Value{Value: serviceProduct.ServiceProductName},
		Time5:  util.Value{Value: EndTime},
		Thing1: util.Value{Value: os.Getenv("ShopName")},
		Thing4: util.Value{Value: "小家伙已经洗好啦，请您尽快接它回家哦"},
	}
	util.WechatSendReserveEndService(wxAccessToken, openId, reserveData)
}
