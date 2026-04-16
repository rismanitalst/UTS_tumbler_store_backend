package services

import (
	"errors"

	"github.com/rismanitalst/UTS_tumbler_store/models"
	"github.com/rismanitalst/UTS_tumbler_store/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{
		productRepo: repositories.NewProductRepository(),
	}
}

// GetAll ambil semua produk dengan pagination + filter category
func (s *ProductService) GetAll(page, limit int, category string) ([]models.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	return s.productRepo.FindAll(page, limit, category)
}

// GetByID ambil produk berdasarkan ID
func (s *ProductService) GetByID(id uint) (*models.Product, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}

	return s.productRepo.FindByID(id)
}

// Create membuat produk baru
func (s *ProductService) Create(req *models.CreateProductRequest) (*models.Product, error) {
	if req == nil {
		return nil, errors.New("request tidak boleh kosong")
	}

	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		ImageURL:    req.ImageURL,
		IsActive:    true,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

// Update produk (partial update)
func (s *ProductService) Update(id uint, req *models.UpdateProductRequest) (*models.Product, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}
	if req == nil {
		return nil, errors.New("request tidak boleh kosong")
	}

	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update hanya field yang dikirim (pointer nil = tidak diupdate)
	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Stock != nil {
		product.Stock = *req.Stock
	}
	if req.Category != nil {
		product.Category = *req.Category
	}
	if req.ImageURL != nil {
		product.ImageURL = *req.ImageURL
	}

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

// Delete produk (soft delete)
func (s *ProductService) Delete(id uint) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}

	return s.productRepo.Delete(id)
}
