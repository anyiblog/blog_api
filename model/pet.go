package model

import (
	"anyiblog/conf"
	"anyiblog/serializer/params/api"
	"anyiblog/util"
	"fmt"
	"time"
)

// Pet 宠物信息表
type Pet struct {
	PetID            string    `gorm:"primary_key;column:pet_id;type:char(36);not null" json:"PetID"`
	UserID           string    `gorm:"column:user_id;type:char(36);not null" json:"-"`                     // 关联用户ID
	PetName          string    `gorm:"column:pet_name;type:varchar(255);not null" json:"PetName"`          // 姓名
	PetAvatar        string    `gorm:"column:pet_avatar;type:char(36)" json:"PetAvatar"`                   // 头像
	PetVariety       string    `gorm:"column:pet_variety;type:varchar(255);not null" json:"PetVariety"`    // 品种
	PetBirthday      time.Time `gorm:"column:pet_birthday;type:datetime;not null" json:"PetBirthday"`      // 生日
	PetAge           int       `gorm:"column:pet_age;type:int;not null" json:"PetAge"`                     // 年龄
	PetWeight        float64   `gorm:"column:pet_weight;type:double;not null" json:"PetWeight"`            // 体重
	PetGender        string    `gorm:"column:pet_gender;type:varchar(255);not null" json:"PetGender"`      // 性别
	PetSterilization int       `gorm:"column:pet_sterilization;type:int;not null" json:"PetSterilization"` // 0 绝育，1未绝育
	PetHair          int       `gorm:"column:pet_hair;type:int;not null" json:"PetHair"`                   // 0 无毛 1 短毛 2长毛
}

// CreatePet 创建宠物信息
func CreatePet(param *api.AddMyPetParam, userId string) (Pet, bool) {
	petId := util.GenerateUUID()
	pet := Pet{
		PetID:            petId,
		UserID:           userId,
		PetName:          param.PetName,
		PetAvatar:        param.PetAvatar,
		PetVariety:       param.PetVariety,
		PetBirthday:      util.StrTimeFormat(param.PetBirthday, "2006-01-02 15:04:05"),
		PetAge:           param.PetAge,
		PetWeight:        param.PetWeight,
		PetGender:        param.PetGender,
		PetSterilization: param.PetSterilization,
		PetHair:          param.PetHair,
	}
	if conf.DB.Create(pet).RowsAffected > 0 {
		pet.PetAvatar = GetImgUrl(pet.PetAvatar)
		return pet, true
	} else {
		return pet, false
	}
}

// GetPetList 获取某个用户下的所有宠物列表
func GetPetList(userId string) ([]Pet, bool) {
	var petList []Pet
	i := conf.DB.Model(Pet{}).Where("user_id = ?", userId).Find(&petList)
	for key := range petList {
		petList[key].PetAvatar = GetImgUrl(petList[key].PetAvatar)
	}
	fmt.Println(petList)
	if i.RowsAffected > 0 {
		return petList, true
	} else {
		return nil, false
	}
}

// GetPetInfo 获取宠物信息
func GetPetInfo(petId string) (Pet, bool) {
	var petInfo Pet
	i := conf.DB.Model(Pet{}).Where("pet_id = ?", petId).Find(&petInfo)
	if i.RowsAffected > 0 {
		petInfo.PetAvatar = GetImgUrl(petInfo.PetAvatar)
		return petInfo, true
	} else {
		return Pet{}, false
	}
}
