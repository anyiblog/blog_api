package api

type QueryProductParam struct {
	ProductName string `json:"ProductName"`
	Limit       int    `json:"Limit"`
	Page        int    `json:"Page"`
}

type product struct {
	SpuID    string `json:"SpuID" binding:"required"`
	SkuID    string `json:"SkuID" binding:"required"`
	BuyCount int    `json:"BuyCount" binding:"required"`
}

// BuyProductParam JSON数组
type BuyProductParam struct {
	AddressID   string    `json:"AddressID" binding:"required"`
	OtherInfo   string    `json:"OtherInfo"`
	ProductList []product `json:"ProductList" binding:"required"`
}
