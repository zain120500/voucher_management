package model

type Brand struct {
	ID       int       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	Vouchers []Voucher `gorm:"foreignKey:BrandID" json:"vouchers"`
}

type Voucher struct {
	ID          int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Title       string `gorm:"type:varchar(150);not null" json:"title"`
	CostInPoint int    `gorm:"type:smallint;not null" json:"cost_in_point"`
	BrandID     int    `json:"brand_id"`
	Brand       *Brand `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"brand"`
}

type Redemption struct {
	ID           int       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CustomerName string    `gorm:"type:varchar(100);not null" json:"customer_name"`
	TotalPoints  int       `gorm:"type:smallint;not null" json:"total_points"`
	Vouchers     []Voucher `gorm:"many2many:redemption_vouchers;" json:"vouchers"`
}

type RedemptionVoucher struct {
	RedemptionID int `gorm:"primaryKey;not null" json:"redemption_id"`
	VoucherID    int `gorm:"primaryKey;not null" json:"voucher_id"`
}
