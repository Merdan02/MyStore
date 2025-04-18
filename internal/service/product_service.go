package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"mystore/internal/models"
	"mystore/internal/repository"
)

type ProductService interface {
	GetAllProduct() ([]*models.Product, error)
	GetById(id int) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id int) error
}

type productService struct {
	repo   repository.ProductRepo
	logger *zap.Logger
}

func NewProductService(repo repository.ProductRepo, logger *zap.Logger) ProductService {
	return &productService{
		repo:   repo,
		logger: logger,
	}
}

func (p *productService) GetAllProduct() ([]*models.Product, error) {
	products, err := p.repo.GetAllProduct()
	if err != nil {
		p.logger.Error("error of getting all products")
		return nil, fmt.Errorf("productService.GetAllProducts: %w", err)
	}
	return products, nil
}

func (p *productService) GetById(id int) (*models.Product, error) {
	if id == 0 || id <= 0 {
		p.logger.Warn("invalid product id", zap.Int("id", id))
		return nil, errors.New("invalid product id")
	}
	product, err := p.repo.GetById(id)
	if err != nil {
		p.logger.Error("error of getting product", zap.Int("id", id), zap.Error(err))
		return nil, errors.New("failed to get product")
	}
	return product, nil
}

func (p *productService) Update(product *models.Product) error {
	if product.Name == "" || product.Price <= 0 || product.Description == "" || product.Quantity <= 0 {
		p.logger.Warn("invalid product fields", zap.Any("product", product))
		return errors.New("invalid product fields")
	}

	if product.ID <= 0 {
		p.logger.Warn("invalid product id", zap.Int("id", product.ID))
		return errors.New("invalid product id")
	}

	err := p.repo.Update(product)
	if err != nil {
		p.logger.Error("error of updating product", zap.Int("id", product.ID), zap.Error(err))
		return errors.New("failed to update product")
	}
	return nil
}

func (p *productService) Create(product *models.Product) error {
	if product.Name == "" || product.Price <= 0 || product.Description == "" || product.Quantity <= 0 {
		p.logger.Warn("invalid product fields", zap.Any("product", product))
		return errors.New("invalid product fields")
	}
	err := p.repo.Create(product)
	if err != nil {
		p.logger.Error("error of creating user", zap.Error(err))
		return errors.New("failed to create user")
	}
	return nil
}
func (p *productService) Delete(id int) error {
	if id <= 0 {
		p.logger.Warn("invalid product id", zap.Int("id", id))
		return errors.New("invalid product id")
	}
	err := p.repo.Delete(id)
	if err != nil {
		p.logger.Error("error of deleting user", zap.Error(err))
		return errors.New("failed to delete user")
	}
	return nil
}
