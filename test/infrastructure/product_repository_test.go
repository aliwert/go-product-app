package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"go-product-app/common/postgressql"
	"go-product-app/domain"
	"go-product-app/repository"
	"os"
	"testing"
)

var productRepository repository.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgressql.GetConnectionPool(ctx, postgressql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		Dbname:                "productapp",
		Username:              "postgres",
		Password:              "585858",
		MaxConnections:        "5",
		MaxConnectionIdleTime: "30s",
	})
	productRepository = repository.NewProductRepository(dbPool)
	fmt.Println("Before all test")
	exitCode := m.Run()
	fmt.Println("After all test")
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)

}

func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)

}

func TestGetAllProducts(t *testing.T) {
	setup(ctx, dbPool)
	expectedProduct := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Iron",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Washing Machine",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Floor Lamp",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "QWE DECORATION",
		},
	}
	t.Run("Get All Products", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProduct, actualProducts)
	})

	clear(ctx, dbPool)
}

func TestGetAllProductsByStore(t *testing.T) {
	setup(ctx, dbPool)
	expectedProduct := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Iron",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Washing Machine",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}
	t.Run("Get All Products By Store", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expectedProduct, actualProducts)
	})

	clear(ctx, dbPool)
}

func TestAddProduct(t *testing.T) {
	expectedProduct := []domain.Product{
		{
			Id:       1,
			Name:     "Coffee Machine",
			Price:    5000.0,
			Discount: 10.0,
			Store:    "FGH HOME",
		},
	}
	newProduct := domain.Product{
		Name:     "Coffee Machine",
		Price:    5000.0,
		Discount: 10.0,
		Store:    "FGH HOME",
	}
	t.Run("Add Product", func(t *testing.T) {
		productRepository.AddProduct(newProduct)
		products := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(products))
		assert.Equal(t, expectedProduct, products)
	})

	clear(ctx, dbPool)
}

func TestGetById(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("Get Product By Id", func(t *testing.T) {
		productsId, _ := productRepository.GetProductById(2)
		assert.Equal(t, domain.Product{
			Id:       2,
			Name:     "Iron",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		}, productsId)

	})
	clear(ctx, dbPool)
}
