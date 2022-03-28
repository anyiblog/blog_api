package admin

type GetAllBrandParams struct {
	Limit int    `binding:"required" json:"Limit"`
	Page  int    `binding:"required" json:"Page"`
}
