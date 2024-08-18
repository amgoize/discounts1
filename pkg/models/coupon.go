package models

type Coupon struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Code     string  `gorm:"not null;unique" json:"code"`
	Discount float64 `gorm:"not null" json:"discount"`
	ClientID uint    `gorm:"not null" json:"client_id"`
}

func (Coupon) TableName() string {
	return "coupons"
}
