package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"go-product-app/domain"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(product domain.Product) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool: dbPool}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products")
	if err != nil {
		log.Error("Error while getting all products %v", err)
		return []domain.Product{}
	}

	return getProducts(productRows)
}
func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()

	getProductsByStoreNameSql := `SELECT * FROM products WHERE store = $1`

	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)
	if err != nil {
		log.Error("Error while getting products by store %v", err)
		return []domain.Product{}
	}
	return getProducts(productRows)

}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insert_sql := `INSERT INTO products (name, price, discount,store) VALUES ($1, $2, $3, $4)`
	addNewProduct, err := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		log.Error("Error while adding product %v", err)
		return err
	}
	log.Info(fmt.Printf("Product added to product store %v\n", addNewProduct))
	return nil
}

func getProducts(productRows pgx.Rows) []domain.Product {
	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		err := productRows.Scan(&id, &name, &price, &discount, &store)
		if err != nil {
			log.Error("Error while getting all products %v", err)
			return []domain.Product{}
		}
		products = append(products, domain.Product{Id: id, Name: name, Price: price, Discount: discount, Store: store})
	}
	return products
}
