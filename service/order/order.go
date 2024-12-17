package order

import (
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/repository"
	"github.com/MatthewAraujo/min-ecommerce/service/auth"
	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/gorilla/mux"
)

var logger = utils.NewParentLogger("Rota api/v1/order")

type Handler struct {
	Service types.OrderService
	store   repository.Queries
}

func NewHandler(Service types.OrderService, store repository.Queries) *Handler {
	return &Handler{
		Service: Service,
		store:   store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", auth.WithJWTAuth(h.Order, h.store, "user")).Methods(http.MethodPost)
}

func (h *Handler) Order(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path, "Creating customer")

	var payload types.CreateOrderPayload

	logger.Info("Parsing JSON")
	if err := utils.ParseJSON(r, &payload); err != nil {
		logger.LogError(r.URL.Path, err, "Erro ao parsear o JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := h.Service.Order(&payload)
	if err != nil {
		logger.LogError(r.URL.Path, err)
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, status, map[string]string{"response": "order created"})

}
