package model

type Brand struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Vouchers []Voucher `gorm:"foreignKey:BrandID" json:"vouchers"`
}

type Voucher struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	CostInPoint int    `json:"cost_in_point"`
	BrandID     uint   `json:"brand_id"`
	Brand       Brand  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"brand"`
}

type Redemption struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `json:"customer_name"`
	TotalPoints  int       `json:"total_points"`
	Vouchers     []Voucher `gorm:"many2many:redemption_vouchers;" json:"vouchers"`
}
