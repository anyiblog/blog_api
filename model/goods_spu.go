package model

import (
	"anyiblog/conf"
	adminParams "anyiblog/serializer/params/admin"
	"anyiblog/util"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"time"
)

// GoodsSpu Spu表（产品表）
type GoodsSpu struct {
	SpuID             string    `gorm:"primary_key;column:spu_id;type:char(36);not null" json:"spu_id"`                   // 产品主键ID
	BrandID           string    `gorm:"column:brand_id;type:char(36)" json:"brand_id"`                                    // 品牌ID
	SpuImg            string    `gorm:"column:spu_img;type:char(36);not null" json:"spu_img"`                             // 产品封面图，默认轮播图第一项
	SpuPrice          float64   `gorm:"column:spu_price;type:decimal(18,2);not null" json:"spu_price"`                    // 产品价格（默认对应SKU的普通用户售价）
	ImgBanner         string    `gorm:"column:img_banner;type:json;not null" json:"img_banner"`                           // 产品轮播图（json）
	SpuName           string    `gorm:"column:spu_name;type:varchar(36);not null" json:"spu_name"`                        // 产品名称
	SpuRecommendTitle string    `gorm:"column:spu_recommend_title;type:varchar(50);not null" json:"spu_recommend_title"`  // 产品推荐标语
	SpuSales          int       `gorm:"column:spu_sales;type:int" json:"spu_sales"`                                       // 产品销量
	SpuSpec           string    `gorm:"column:spu_spec;type:longtext;not null" json:"spu_spec"`                           // 产品规格信息，json字符串
	SpuShippingMethod string    `gorm:"column:spu_shipping_method;type:varchar(255);not null" json:"spu_shipping_method"` // 产品配送方式
	SpuOtherInfo      string    `gorm:"column:spu_other_info;type:json;not null" json:"spu_other_info"`                   // 产品其他信息
	DetailContent     string    `gorm:"column:detail_content;type:longtext;not null" json:"detail_content"`               // 商品介绍，json图片列表
	SpuStatus         uint      `gorm:"column:spu_status;type:int unsigned;not null" json:"spu_status"`                   // 0上架，1下架
	CreatedAt         time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`                      // 创建时间
	UpdatedAt         time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                               // 更新时间
}

type ResSpu struct { // 数据表映射结构
	SpuID             string    `json:"SpuID"`
	BrandID           string    `json:"BrandName"`
	SpuPrice          string    `json:"SpuPrice"`
	SpuSales          string    `json:"SpuSales"`
	SpuImg            string    `json:"SpuImg"`
	SpuName           string    `json:"ProductName"`
	SpuShippingMethod string    `json:"ProductShippingMethod"`
	SpuStatus         uint      `json:"ProductStatus"`
	UpdatedAt         time.Time `json:"UpdatedAt"`
	SpuSpec           string    `json:"SpuSpec"`
}
type ResProduct struct {
	ResSpu
	Sku []GoodsSku `json:"Sku,omitempty"`
}

// ResProductList 返回商品列表
type ResProductList struct {
	Count       int64        `json:"Count"`
	ProductList []ResProduct `json:"ProductList"`
}

// ResProductDetail 返回商品详情
type ResProductDetail struct {
	GoodsSpu
	Sku []GoodsSku `json:"Sku,omitempty"`
}

// GetSpuInfoByQuery 根据查询字段名和查询参数获取结果
func GetSpuInfoByQuery(Fields interface{}, QueryParams map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	conf.DB.Model(GoodsSpu{}).Select(Fields).Where(QueryParams).First(&result)
	return result
}

// AddGoodsSpu 新增产品SPU
func AddGoodsSpu(params *adminParams.AddProductParam) bool {
	spuID := util.GenerateUUID()
	goodsSpu := GoodsSpu{
		SpuID:             spuID,
		BrandID:           params.BrandID,
		SpuImg:            gjson.Parse(params.BannerList).Get("0").String(), // 获取轮播图第一项
		SpuPrice:          params.SkuInfo[0].SkuUserPrice,
		ImgBanner:         params.BannerList,
		SpuName:           params.ProductName,
		SpuRecommendTitle: params.ProductRecommendTitle,
		SpuSpec:           params.SpuSpec,
		SpuShippingMethod: params.ShippingMethod,
		SpuOtherInfo:      params.OtherInfo,
		DetailContent:     params.DetailsInfo,
		SpuStatus:         0, // 默认上架商品
		CreatedAt:         util.NowTime(),
		UpdatedAt:         util.NowTime(),
	}
	if conf.DB.Create(goodsSpu).RowsAffected > 0 {
		if AddGoodsSku(params.SkuInfo, spuID) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// GetGoodsDetail 获取单个产品详情
func GetGoodsDetail(spuID string) (bool, ResProductDetail) {
	var spu GoodsSpu
	var pDetail ResProductDetail
	db := conf.DB.Model(GoodsSpu{}).Where("spu_id = ? ", spuID).Find(&spu)
	if db.RowsAffected > 0 {
		pDetail.GoodsSpu = spu
		getBrandNameStatus, brandId := GetBrandName(spu.BrandID)
		if getBrandNameStatus {
			pDetail.BrandID = brandId
		} else {
			pDetail.BrandID = ""
		}
		pDetail.SpuImg = GetImgUrl(spu.SpuImg)
		//pDetail.ImgBanner = GetImgUrl(spu.ImgBanner)

		jsonArr := gjson.Parse(spu.ImgBanner)
		tempImgList := make([]string, 0)
		for _, v := range jsonArr.Array() {
			tempImgList = append(tempImgList, GetImgUrl(v.String()))
		}
		marshal, err := json.Marshal(tempImgList)
		if err != nil {
			return false, ResProductDetail{}
		}
		pDetail.ImgBanner = string(marshal)
		pDetail.Sku = GetAllSku(spu.SpuID)
		return true, pDetail
	} else {
		return false, ResProductDetail{}
	}
}

// GetAllGoods 获取所有产品
func GetAllGoods(QueryParams map[string]interface{}, LikeParams map[string]interface{}, limit, page int, isGetSku bool) ResProductList {
	var spuList []ResSpu // spu信息列表
	var p ResProduct     // 单个产品信息
	pList := ResProductList{}
	//var pList ResProductList // 所有产品列表
	// 基本查询语句
	db := conf.DB.Model(GoodsSpu{}).Where(QueryParams).Order("created_at desc").Limit(limit).Offset((page - 1) * limit)
	countDB := conf.DB.Model(GoodsSpu{}).Where(QueryParams)
	// 拼接模糊查询语句
	for likeName := range LikeParams {
		// WHERE name LIKE '%jin%';
		db = db.Where(fmt.Sprintf(likeName+" like %q ", "%"+LikeParams[likeName].(string)+"%"))
		countDB = countDB.Where(fmt.Sprintf(likeName+" like %q ", "%"+LikeParams[likeName].(string)+"%"))
	}
	db.Find(&spuList)
	countDB.Count(&pList.Count)

	for _, spuItem := range spuList { // 合并产品数据信息，并追加到产品列表
		p.ResSpu = spuItem
		getBrandNameStatus, brandId := GetBrandName(spuItem.BrandID)
		if getBrandNameStatus {
			p.BrandID = brandId
		} else {
			p.BrandID = "暂未关联"
		}
		p.SpuImg = GetImgUrl(spuItem.SpuImg)
		if isGetSku {
			p.Sku = GetAllSku(spuItem.SpuID)
		}
		pList.ProductList = append(pList.ProductList, p)
	}
	return pList
}

// SetProductStatus 设置产品状态
func SetProductStatus(spuID string, status uint) bool {
	if conf.DB.Model(GoodsSpu{}).Where("spu_id = ?", spuID).Update("spu_status", status).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// SetProductSales 更新产品销量
func SetProductSales(spuID string, salesCount int) bool {
	var oldSales struct {
		SpuSales int
	}
	conf.DB.Model(GoodsSpu{}).Select("spu_sales").Where("spu_id = ?", spuID).Find(&oldSales)
	upSales := conf.DB.Model(GoodsSpu{}).Where("spu_id = ?", spuID).Update("spu_sales", oldSales.SpuSales+salesCount)
	if upSales.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// DeleteProduct 删除产品
func DeleteProduct(spuID string) bool {
	if conf.DB.Where("spu_id = ?", spuID).Delete(GoodsSpu{}).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
