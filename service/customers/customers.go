package customers

import (
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.CostumersStore
}

func NewHandler(store types.CostumersStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/customers", h.getAllCustomers).Methods(http.MethodGet)
}

func (h *Handler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := h.store.GetAllCustomers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, customers)

}
