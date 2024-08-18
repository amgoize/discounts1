package models

type Client struct {
	ID      uint     `gorm:"primaryKey"`
	Name    string   `gorm:"not null"`
	Coupons []Coupon `gorm:"foreignKey:ClientID"`
}

func (Client) TableName() string {
	return "clients"
}
