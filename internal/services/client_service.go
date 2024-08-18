package services

import (
	"discounts/internal/repositories"
	"discounts/pkg/models"
	"errors"
)

type ClientService struct {
	Repo *repositories.ClientRepository
}

// CreateClient
func (s *ClientService) CreateClient(client *models.Client) error {
	if client.Name == "" {
		return errors.New("client name is required")
	}
	return s.Repo.CreateClient(client)
}

// GetClientWithCoupons
func (s *ClientService) GetClientWithCoupons(id uint) (*models.Client, error) {
	return s.Repo.GetClientWithCoupons(id)
}

// UpdateClient
func (s *ClientService) UpdateClient(client *models.Client) error {
	if client.ID == 0 {
		return errors.New("client ID is required")
	}
	return s.Repo.UpdateClient(client)
}

// DeleteClient
func (s *ClientService) DeleteClient(id uint) error {
	return s.Repo.DeleteClient(id)
}

// GetAllClients
func (s *ClientService) GetAllClients() ([]models.Client, error) {
	return s.Repo.GetAllClients()
}
