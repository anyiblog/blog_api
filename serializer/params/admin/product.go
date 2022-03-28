package admin

type SkuInfo struct {
	SkuImg       string  `json:"Image"`
	SkuVipPrice  float64 `json:"VipPrice"`
	SkuUserPrice float64 `json:"UserPrice"`
	SkuPrice     float64 `json:"Price"`
	SkuStock     int     `json:"Stock"`
	SkuWeight    float64 `json:"Weight"`
	SkuSpec      string  `json:"sku"`
}
type AddProductParam struct {
	BrandID               string    `json:"BrandID"`
	BannerList            string    `json:"BannerList"`
	ProductName           string    `json:"ProductName"`
	ProductRecommendTitle string    `json:"ProductRecommendTitle"`
	SpuSpec               string    `json:"SpuSpec"`
	ShippingMethod        string    `json:"ShippingMethod"`
	OtherInfo             string    `json:"OtherInfo"`
	DetailsInfo           string    `json:"DetailsInfo"`
	SkuInfo               []SkuInfo `json:"SkuInfo"`
}

type GetAllProductParam struct {
	ProductName string `json:"ProductName"`
	BrandName   string `json:"BrandName"`
	Status      int    `json:"ProductStatus"`
	Limit       int    `json:"Limit"`
	Page        int    `json:"Page"`
}

type SetProductStatusParam struct {
	ProductIDArray string `json:"ProductIDArray"`
}
