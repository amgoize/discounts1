package services

import (
	"discounts/internal/repositories"
	"discounts/pkg/models"
	"errors"
	"log"
)

type CouponService struct {
	Repo *repositories.CouponRepository
}

// CreateCoupon
func (s *CouponService) CreateCoupon(coupon *models.Coupon) error {
	if coupon.Code == "" {
		return errors.New("coupon code is required")
	}
	log.Printf("333Decoded Coupon: %+v", coupon)
	if coupon.Discount <= 0 {
		return errors.New("coupon discount must be greater than zero")
	}
	return s.Repo.CreateCoupon(coupon)
}

// GetCoupon
func (s *CouponService) GetCoupon(id uint) (*models.Coupon, error) {
	return s.Repo.GetCoupon(id)
}

// UpdateCoupon
func (s *CouponService) UpdateCoupon(coupon *models.Coupon) error {
	if coupon.ID == 0 {
		return errors.New("coupon ID is required")
	}
	if coupon.Code == "" {
		return errors.New("coupon code is required")
	}
	if coupon.Discount <= 0 {
		return errors.New("coupon discount must be greater than zero")
	}
	return s.Repo.UpdateCoupon(coupon)
}

// DeleteCoupon
func (s *CouponService) DeleteCoupon(id uint) error {
	return s.Repo.DeleteCoupon(id)
}

// GetAllCoupons
func (s *CouponService) GetAllCoupons() ([]models.Coupon, error) {
	return s.Repo.GetAllCoupons()
}
