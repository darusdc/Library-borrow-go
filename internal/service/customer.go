package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/darusdc/belajar-go/domain"
	"github.com/darusdc/belajar-go/dto"
	"github.com/google/uuid"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

// Show implements domain.CustomerService.
func (cr *customerService) Show(ctx context.Context, id string) (dto.CustomerData, error) {
	persisted, err := cr.customerRepository.FindById(ctx, id)

	if err != nil {
		return dto.CustomerData{}, err
	}

	if persisted.Id != "" {
		return dto.CustomerData{}, errors.New("customer data isn't exist")
	}

	return dto.CustomerData{
		Id:   persisted.Id,
		Code: persisted.Code,
		Name: persisted.Name,
	}, err
}

// Delete implements domain.CustomerService.
func (cr *customerService) Delete(ctx context.Context, id string) error {
	exist, err := cr.customerRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if exist.Id == "" {
		return errors.New("account isn't exist")
	}

	return cr.customerRepository.Delete(ctx, id)
}

// Update implements domain.CustomerService.
func (cr *customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := cr.customerRepository.FindById(ctx, req.ID)

	if err != nil {
		return err
	}

	if persisted.Id != "" {
		return errors.New("customer data isn't exist")
	}

	persisted.Code = req.Code
	persisted.Name = req.Name
	persisted.UpdateAt = sql.NullTime{Valid: true, Time: time.Now()}

	return cr.customerRepository.Update(ctx, &persisted)
}

// Create implements domain.CustomerService.
func (cr *customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		Id:        uuid.NewString(),
		Name:      req.Name,
		Code:      req.Code,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	return cr.customerRepository.Save(ctx, &customer)
}

// Index implements domain.CustomerService.
func (cr *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := cr.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData

	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			Id:   v.Id,
			Code: v.Code,
			Name: v.Name,
		})
	}

	return customerData, nil
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}
