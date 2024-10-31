package service

import (
	"go-product-app/domain"
	"go-product-app/repository"
	"go-product-app/service/dto"
)

type IProductService interface {
	Add(productCreate dto.ProductCreate) error
	DeleteById(productId int64) error
	GetById(productId int64) domain.Product
	UpdatePrice(productId int64, newPrice float32) error
	GetAllProducts() []domain.Product
	GetAllProductByStore(storeName string) []domain.Product
}

type ProductService struct {
	productRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) Add(productCreate dto.ProductCreate) error {
	return nil
}
func (p *ProductService) DeleteById(productId int64) error {
	return nil

}

func (p *ProductService) GetById(productId int64) domain.Product {
	return domain.Product{}
}

func (p *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	return nil
}

func (p *ProductService) GetAllProducts() []domain.Product {
	return []domain.Product{}
}

func (p *ProductService) GetAllProductByStore(storeName string) []domain.Product {
	return []domain.Product{}
}
