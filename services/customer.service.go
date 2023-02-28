package services

import "dayang/models"

type CustomerService interface {
	GetOpenID(*string) (string, error)
	CreateCustomer(*models.Customer) error
	GetCustomerByOpenID(*string) []models.Customer
	GetCustomerByName(*string) []models.Customer
	GetAll(*string, int, int) ([]models.Customer, int64, error)
	UpdateCustomer(*models.Customer) error
	DeleteCustomer([]models.Customer) error
}
