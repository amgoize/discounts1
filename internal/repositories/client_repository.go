package repositories

import (
	"discounts/pkg/models"

	"gorm.io/gorm"
)

type ClientRepository struct {
	DB *gorm.DB
}

// CreateClient
func (r *ClientRepository) CreateClient(client *models.Client) error {
	return r.DB.Create(client).Error
}

// GetClientWithCoupons
func (r *ClientRepository) GetClientWithCoupons(id uint) (*models.Client, error) {
	var client models.Client
	err := r.DB.Preload("Coupons").First(&client, id).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// UpdateClient обновляет информацию о клиенте
func (r *ClientRepository) UpdateClient(client *models.Client) error {
	return r.DB.Save(client).Error
}

// DeleteClient удаляет клиента по ID
func (r *ClientRepository) DeleteClient(id uint) error {
	return r.DB.Delete(&models.Client{}, id).Error
}

// GetAllClients возвращает всех клиентов
func (r *ClientRepository) GetAllClients() ([]models.Client, error) {
	var clients []models.Client
	err := r.DB.Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}
