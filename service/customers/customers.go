package customers

import (
	"net/http"

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
	router.HandleFunc("/customers", h.CreateCustomer).Methods(http.MethodGet)
}

func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateCustomerPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ct, status, err := h.Service.CreateCustomer(&payload)
	if err != nil {
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, status, ct)

}
