package customers

import (
	"fmt"
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/go-playground/validator/v10"
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
	router.HandleFunc("/register", h.CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/login", h.Login).Methods(http.MethodPost)
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginCustomerPayload

	logger.Info(r.URL.Path, "Starting login process")

	// Parse o JSON recebido
	if err := utils.ParseJSON(r, &payload); err != nil {
		logger.LogError(r.URL.Path, fmt.Errorf("error parsing JSON: %w", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Info("Payload parsed successfully")

	// Validação do payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		logger.LogError(r.URL.Path, fmt.Errorf("validation error: %s", errors))
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %s", errors))
		return
	}

	logger.Info("Payload validated successfully")

	// Chamando o serviço de login
	token, status, err := h.Service.Login(&payload)
	if err != nil {
		logger.LogError(r.URL.Path, fmt.Errorf("login failed: %w", err))
		utils.WriteError(w, status, err)
		return
	}

	logger.Info("Login successful, returning token")

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}
