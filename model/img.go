package model

import (
	"anyiblog/conf"
	"anyiblog/util"
	"time"
)

// Img 图片表
type Img struct {
	ImgID     string    `gorm:"primary_key;column:img_id;type:char(36);not null" json:"img_id"` // 图片主键ID
	ImgURL    string    `gorm:"column:img_url;type:varchar(255);not null" json:"img_url"`       // 图片地址
	ImgTag    int       `gorm:"column:img_tag;type:int;not null" json:"img_tag"`                // 图片标签
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`    // 创建时间
}

// ResTagInfo 标签信息
type ResTagInfo struct {
	TagId   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

// ResImgList 图片列表
type ResImgList struct {
	Count   int64   `json:"count"`
	ImgList []Img `json:"img_list"`
}

// GetImgUrl 获取图片URL
func GetImgUrl(imgId string) string {
	var resImg struct {
		ImgID  string
		ImgURL string
	}
	var i int64
	i = conf.DB.Model(Img{}).Select("img_id,img_url").Where("img_id = ? ", imgId).Scan(&resImg).RowsAffected
	if i > 0 {
		return resImg.ImgURL
	}
	return ""
}

// GetImgListForTag 根据tag获取图片列表
func GetImgListForTag(tag, limit, page int) ResImgList {
	if page <= 0 {
		return ResImgList{
			Count:   0,
			ImgList: nil,
		}
	} else {
		var resImgList ResImgList
		conf.DB.Debug().Model(Img{}).Where("img_tag = ? ", tag).Order("created_at desc").Limit(limit).Offset((page - 1) * limit).Find(&resImgList.ImgList)
		conf.DB.Debug().Model(Img{}).Where("img_tag = ? ", tag).Count(&resImgList.Count)
		return resImgList
	}
}

// ImgTagUpdate 根据oldTagId更新图片newTagId
func ImgTagUpdate(oldTagId, newTagId int, imgId string) bool {
	if conf.DB.Debug().Model(Img{}).Where("img_id = ? and img_tag = ? ", imgId, oldTagId).Update("img_tag", newTagId).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// CreateImgByUrl 图片创建
func CreateImgByUrl(imgUrl, imgId string) bool {
	i := Img{
		ImgID:     imgId,
		ImgURL:    imgUrl,
		CreatedAt: util.NowTime(),
	}
	if conf.DB.Create(i).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// DeleteImg 图片删除
func DeleteImg(imgUrl string) bool {
	if conf.DB.Where("img_url = ?", imgUrl).Delete(&Img{}).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// CreateImgForUrl 从已有网络url的图片创建，不上传COS，返回ImgID
func CreateImgForUrl(imgUrl string) string {
	imgId := util.GenerateUUID()
	i := Img{
		ImgID:     imgId,
		ImgURL:    imgUrl,
		CreatedAt: util.NowTime(),
	}
	conf.DB.Create(i)
	var resImg = struct {
		ImgID  string
		ImgURL string
	}{}
	conf.DB.Model(Img{}).Select("img_id,img_url").Where("img_id = ? ", imgId).Scan(&resImg)
	return resImg.ImgID
}
