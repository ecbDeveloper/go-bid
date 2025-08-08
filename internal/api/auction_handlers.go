package api

import (
	"errors"
	"net/http"

	"github.com/ecbDeveloper/go-bid/internal/services"
	"github.com/ecbDeveloper/go-bid/internal/shared"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// handleSubscribeUserToAuction atualiza a conexão para websocet, assim inscrevendo
// o usuário em uma auction room para ele poder fazer lances por um produto
func (api *Api) handleSubscribeUserToAuction(w http.ResponseWriter, r *http.Request) {
	rawProductId := chi.URLParam(r, "productId")

	productId, err := uuid.Parse(rawProductId)
	if err != nil {
		shared.EncodeJson(w, http.StatusBadRequest, map[string]string{
			"error": "invalid product id - must be a valid uuid",
		})
		return
	}

	_, err = api.ProductService.GetProduct(r.Context(), productId)
	if err != nil {
		if errors.Is(err, services.ErrProductNotFound) {
			shared.EncodeJson(w, http.StatusNotFound, map[string]string{
				"error": "product not found",
			})
			return
		}

		shared.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "unexpected internal server error",
		})
		return
	}

	userId, ok := api.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)
	if !ok {
		shared.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "unexpected internal server error",
		})
		return
	}

	api.AuctionLobby.Lock()
	room, ok := api.AuctionLobby.Rooms[productId]
	api.AuctionLobby.Unlock()

	if !ok {
		shared.EncodeJson(w, http.StatusBadRequest, map[string]string{
			"message": "the auction has ended",
		})
		return
	}

	conn, err := api.WsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		shared.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "could not upgrade connection to a websocket protocol",
		})
		return
	}

	client := services.NewClient(room, conn, userId)

	room.Register <- client
	go client.ReadEventLoop()
	go client.WriteEventLoop()
}
