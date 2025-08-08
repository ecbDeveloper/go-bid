package api

import (
	"context"
	"net/http"

	"github.com/ecbDeveloper/go-bid/internal/services"
	"github.com/ecbDeveloper/go-bid/internal/shared"
	"github.com/ecbDeveloper/go-bid/internal/usecase/product"
	"github.com/google/uuid"
)

func (api *Api) handleCreateProductAuction(w http.ResponseWriter, r *http.Request) {
	data, problems, err := shared.DecodeValidJson[product.CreateProductParams](r)
	if err != nil {
		shared.EncodeJson(w, http.StatusUnprocessableEntity, problems)
		return
	}

	userID, ok := api.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)
	if !ok {
		shared.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "unexpected error, try again later",
		})
		return
	}

	productId, err := api.ProductService.CreateProduct(
		r.Context(),
		userID,
		data.ProductName,
		data.Description,
		data.Baseprice,
		data.AuctionEnd,
	)
	if err != nil {
		shared.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to create product, try again later",
		})
	}

	ctx, _ := context.WithDeadline(context.Background(), data.AuctionEnd)
	auctionRoom := services.NewAuctionRoom(ctx, productId, api.BidService)

	go auctionRoom.Run()

	api.AuctionLobby.Lock()
	api.AuctionLobby.Rooms[productId] = auctionRoom
	api.AuctionLobby.Unlock()

	shared.EncodeJson(w, http.StatusOK, map[string]any{
		"message":    "Auction has started with success",
		"product_id": productId,
	})
}
