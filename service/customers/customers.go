package customers

import (
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/gorilla/mux"
)

var logger = utils.NewParentLogger("Rota api/v1/customer")

type Handler struct {
	Service types.CostumersService
}

func NewHandler(Service types.CostumersService) *Handler {
	return &Handler{
		Service: Service,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/customers", h.CreateCustomer).Methods(http.MethodPost)
}

func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path, "Creating customer")

	var payload types.CreateCustomerPayload

	logger.Info("Parsing JSON")
	if err := utils.ParseJSON(r, &payload); err != nil {
		logger.LogError(r.URL.Path, err, "Erro ao parsear o JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := h.Service.CreateCustomer(&payload)
	if err != nil {
		logger.LogError(r.URL.Path, err)
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, status, map[string]string{"response": "user created"})

}
