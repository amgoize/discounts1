package repositories

import (
	"discounts/pkg/models"

	"log"

	"gorm.io/gorm"
)

type CouponRepository struct {
	DB *gorm.DB
}

// CreateCoupon
func (r *CouponRepository) CreateCoupon(coupon *models.Coupon) error {
	log.Printf("222Decoded Coupon: %+v", coupon)
	return r.DB.Create(coupon).Error
}

// GetCoupon
func (r *CouponRepository) GetCoupon(id uint) (*models.Coupon, error) {
	var coupon models.Coupon
	err := r.DB.First(&coupon, id).Error
	if err != nil {
		return nil, err
	}
	return &coupon, nil
}

// UpdateCoupon
func (r *CouponRepository) UpdateCoupon(coupon *models.Coupon) error {
	return r.DB.Save(coupon).Error
}

// DeleteCoupon
func (r *CouponRepository) DeleteCoupon(id uint) error {
	return r.DB.Delete(&models.Coupon{}, id).Error
}

// GetAllCoupons
func (r *CouponRepository) GetAllCoupons() ([]models.Coupon, error) {
	var coupons []models.Coupon
	err := r.DB.Find(&coupons).Error
	if err != nil {
		return nil, err
	}
	return coupons, nil
}
