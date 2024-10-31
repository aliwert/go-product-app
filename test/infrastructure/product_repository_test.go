package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-product-app/common/postgressql"
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

	fmt.Println("Testing all products")

	clear(ctx, dbPool)
}
