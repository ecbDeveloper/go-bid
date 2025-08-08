package services

import (
	"context"
	"errors"
	"time"

	"github.com/ecbDeveloper/go-bid/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrProductNotFound = errors.New("product not fount")

type ProductService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewProductService(pool *pgxpool.Pool) ProductService {
	return ProductService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (ps *ProductService) CreateProduct(
	ctx context.Context,
	sellerId uuid.UUID,
	productName,
	description string,
	baseprice float64,
	auctionEnd time.Time,
) (uuid.UUID, error) {
	args := sqlc.CreateProductParams{
		SellerID:    sellerId,
		ProductName: productName,
		Description: description,
		Baseprice:   baseprice,
		AuctionEnd:  auctionEnd,
	}

	id, err := ps.queries.CreateProduct(ctx, args)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (ps *ProductService) GetProduct(ctx context.Context, productId uuid.UUID) (sqlc.Product, error) {
	product, err := ps.queries.GetProductById(ctx, productId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sqlc.Product{}, ErrProductNotFound
		}

		return sqlc.Product{}, err
	}

	return product, nil

}
