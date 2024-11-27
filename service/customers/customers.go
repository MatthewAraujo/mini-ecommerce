package customers

import (
	"context"
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/repository"
	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	Service types.CostumersService
}

func NewHandler(Service types.CostumersService) *Handler {
	return &Handler{
		Service: Service,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/customers", h.getAllCustomers).Methods(http.MethodGet)
}

func (h *Handler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customerID := int32(1)
	orderItems := []repository.OrderItem{
		{
			ProductID: 1, // ID do primeiro produto
			Quantity:  2, // Quantidade solicitada
		},
		{
			ProductID: 2, // ID do segundo produto
			Quantity:  1, // Quantidade solicitada
		},
	}
	status, err := h.Service.Order(context.Background(), customerID, orderItems)
	if err != nil {
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "done")

}
