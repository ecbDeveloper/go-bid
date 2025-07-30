package api

import (
	"net/http"

	"github.com/ecbDeveloper/go-bid/internal/shared"
	"github.com/ecbDeveloper/go-bid/internal/usecase/product"
	"github.com/google/uuid"
)

func (api *Api) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
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

	id, err := api.ProductService.CreateProduct(
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

	shared.EncodeJson(w, http.StatusOK, map[string]any{
		"message":    "product created successfully",
		"product_id": id,
	})
}
