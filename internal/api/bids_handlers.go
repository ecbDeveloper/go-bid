package api

import (
	"net/http"

	"github.com/ecbDeveloper/go-bid/internal/shared"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api *Api) handleGetHighestBidByProductId(w http.ResponseWriter, r *http.Request) {
	rawProductId := chi.URLParam(r, "productId")

	productId, err := uuid.Parse(rawProductId)
	if err != nil {
		shared.EncodeJson(w, http.StatusBadRequest, map[string]string{
			"error": "invalid product id - must be a valid uuid",
		})
		return
	}

	highestBid, err := api.BidService.GetHighestBidByProductId(r.Context(), productId)
	if err != nil {
		shared.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "unexpected internal server error",
		})
	}

	shared.EncodeJson(w, http.StatusOK, highestBid)
}
