package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"discounts/internal/services"
	"discounts/pkg/models"

	"log"

	"github.com/gorilla/mux"
)

type CouponHandler struct {
	Service *services.CouponService
}

// CreateCoupon
func (h *CouponHandler) CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon models.Coupon

	if err := json.NewDecoder(r.Body).Decode(&coupon); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Printf("111Decoded Coupon: %+v", coupon)

	if err := h.Service.CreateCoupon(&coupon); err != nil {
		http.Error(w, "Failed to create coupon", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(coupon)
}

// GetCoupon
func (h *CouponHandler) GetCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid coupon ID", http.StatusBadRequest)
		return
	}

	coupon, err := h.Service.GetCoupon(uint(id))
	if err != nil {
		http.Error(w, "Coupon not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coupon)
}

// UpdateCoupon
func (h *CouponHandler) UpdateCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid coupon ID", http.StatusBadRequest)
		return
	}

	var coupon models.Coupon
	if err := json.NewDecoder(r.Body).Decode(&coupon); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	coupon.ID = uint(id)
	if err := h.Service.UpdateCoupon(&coupon); err != nil {
		http.Error(w, "Failed to update coupon", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coupon)
}

// DeleteCoupon
func (h *CouponHandler) DeleteCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid coupon ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteCoupon(uint(id)); err != nil {
		http.Error(w, "Failed to delete coupon", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllCoupons
func (h *CouponHandler) GetAllCoupons(w http.ResponseWriter, r *http.Request) {
	coupons, err := h.Service.GetAllCoupons()
	if err != nil {
		http.Error(w, "Failed to get coupons", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coupons)
}
