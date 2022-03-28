//业务逻辑公用函数库

package service

import (
	"anyiblog/cache"
	"anyiblog/serializer"
	"context"
	"errors"
	"github.com/nelsonken/cos-go-sdk-v5/cos"
	"io"
	"os"
	"strings"
)

// WechatOpenInfo 微信解密信息
type WechatOpenInfo struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"openid"`
}

var cacheRedis = cache.NewCache(0, "")

// CosUpload 简单 Cos上传文件
func CosUpload(content io.Reader, path string) (fileUrl string, errInfo error) {
	cc := cos.New(&cos.Option{
		AppID:     os.Getenv("APP_ID"),
		SecretID:  os.Getenv("SECRET_ID"),
		SecretKey: os.Getenv("SECRET_KEY"),
		Region:    os.Getenv("COS_REGION"),
	})
	errInfo = cc.Bucket(os.Getenv("COS_BUCKET_NAME")).UploadObject(context.Background(), path, content, &cos.AccessControl{})
	fileUrl = os.Getenv("COS_Url") + path
	return fileUrl, errInfo
}

// DeleteCosFile 删除Cos文件
func DeleteCosFile(fileName string) (errInfo error) {
	cc := cos.New(&cos.Option{
		AppID:     os.Getenv("APP_ID"),
		SecretID:  os.Getenv("SECRET_ID"),
		SecretKey: os.Getenv("SECRET_KEY"),
		Region:    os.Getenv("COS_REGION"),
	})
	errInfo = cc.Bucket(os.Getenv("COS_BUCKET_NAME")).DeleteObject(context.Background(), fileName)
	return errInfo
}

// CheckSmsRedis 检测短信验证码Redis值是否相等
func CheckSmsRedis(phone, SmsType, smsCode string) bool {
	c := cache.NewCache(serializer.RedisSms, "sms_")
	key := phone + "_" + SmsType
	value := c.Get(key)
	if value == smsCode {
		return true
	} else {
		return false
	}
}

// GetRedisWechatInfo 从缓存中获取微信解密信息
func GetRedisWechatInfo(token string) (WechatOpenInfo, error) {
	cacheInfo := cacheRedis.Get(token)
	var wechatOpenInfo WechatOpenInfo
	var err error
	if len(cacheInfo) != 0 {
		wechatInfos := strings.Split(cacheInfo, "[BCX]")
		wechatOpenInfo = WechatOpenInfo{
			SessionKey: wechatInfos[0],
			OpenId:     wechatInfos[1],
		}
	} else if len(cacheInfo) == 0 {
		err = errors.New("检索用户信息失败")
	}
	return wechatOpenInfo, err
}
