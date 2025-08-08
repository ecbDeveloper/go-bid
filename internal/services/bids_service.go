package services

import (
	"context"
	"errors"

	"github.com/ecbDeveloper/go-bid/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrBidIsTooLow = errors.New("you need to increase your bid to participate in the auction")
)

type BidService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewBidService(pool *pgxpool.Pool) BidService {
	return BidService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (bs *BidService) PlaceBid(
	ctx context.Context,
	productId,
	bidderId uuid.UUID,
	bidAmount float64,
) (sqlc.Bid, error) {
	product, err := bs.queries.GetProductById(ctx, productId)
	if err != nil {
		return sqlc.Bid{}, err
	}

	highestBid, err := bs.GetHighestBidByProductId(ctx, productId)
	if err != nil {
		return sqlc.Bid{}, err
	}

	if bidAmount <= product.Baseprice || bidAmount <= highestBid.BidAmount {
		return sqlc.Bid{}, ErrBidIsTooLow
	}

	args := sqlc.CreateBidParams{
		ProductID: productId,
		BidderID:  bidderId,
		BidAmount: bidAmount,
	}

	highestBid, err = bs.queries.CreateBid(ctx, args)
	if err != nil {
		return sqlc.Bid{}, err
	}

	return highestBid, nil
}

func (bs *BidService) GetAllBidsByProductId(ctx context.Context, productId uuid.UUID) ([]sqlc.Bid, error) {
	bids, err := bs.queries.GetBidsByProductId(ctx, productId)
	if err != nil {
		return nil, err
	}

	return bids, nil
}

func (bs *BidService) GetHighestBidByProductId(ctx context.Context, productId uuid.UUID) (sqlc.Bid, error) {
	highestBid, err := bs.queries.GetHighestBidByProductId(ctx, productId)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return sqlc.Bid{}, err
		}
	}

	return highestBid, nil
}
