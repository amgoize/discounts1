// main.go
package main

import (
	"discounts/internal/handlers"
	"discounts/internal/repositories"
	"discounts/internal/services"
	"discounts/pkg/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := database.SetupDatabase()

	// Инициализация
	clientRepo := &repositories.ClientRepository{DB: db}
	couponRepo := &repositories.CouponRepository{DB: db}

	clientService := &services.ClientService{Repo: clientRepo}
	couponService := &services.CouponService{Repo: couponRepo}

	clientHandler := &handlers.ClientHandler{Service: clientService}
	couponHandler := &handlers.CouponHandler{Service: couponService}

	r := mux.NewRouter()
	r.HandleFunc("/clients", clientHandler.CreateClient).Methods("POST")
	r.HandleFunc("/clients/{id}", clientHandler.GetClientWithCoupons).Methods("GET")
	r.HandleFunc("/clients/{id}", clientHandler.UpdateClient).Methods("PUT")
	r.HandleFunc("/clients/{id}", clientHandler.DeleteClient).Methods("DELETE")
	r.HandleFunc("/clients", clientHandler.GetAllClients).Methods("GET")
	r.HandleFunc("/coupons", couponHandler.CreateCoupon).Methods("POST")
	r.HandleFunc("/coupons/{id}", couponHandler.GetCoupon).Methods("GET")
	r.HandleFunc("/coupons/{id}", couponHandler.UpdateCoupon).Methods("PUT")
	r.HandleFunc("/coupons/{id}", couponHandler.DeleteCoupon).Methods("DELETE")
	r.HandleFunc("/coupons", couponHandler.GetAllCoupons).Methods("GET")

	// Запуск сервера
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
