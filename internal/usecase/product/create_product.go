package product

import (
	"context"
	"time"

	"github.com/ecbDeveloper/go-bid/internal/shared"
)

type CreateProductParams struct {
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
}

const minAuctionDuration = 2 * time.Hour

func (req CreateProductParams) Valid(ctx context.Context) shared.ErrorsValidator {
	var eval shared.ErrorsValidator

	eval.CheckField(shared.NotBlank(req.ProductName), "product_name", "this field cannot be empty")

	eval.CheckField(shared.NotBlank(req.Description), "description", "this field cannot be empty")
	eval.CheckField(shared.MinChars(req.Description, 10) &&
		shared.MaxChars(req.Description, 255),
		"description",
		"this filed must have a length between 10 and 255 chars",
	)

	eval.CheckField(req.Baseprice > 0, "baseprice", "this field must be greater then 0")

	eval.CheckField(
		req.AuctionEnd.Sub(time.Now()) <= minAuctionDuration,
		"auction_end",
		"the auction duration must be at least 2 hours",
	)

	return eval
}
