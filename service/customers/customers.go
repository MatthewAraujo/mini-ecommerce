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
	router.HandleFunc("/customers", h.getAllCustomers).Methods(http.MethodGet)
}

func (h *Handler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "done")
}
