package products

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MatthewAraujo/min-ecommerce/repository"
	"github.com/MatthewAraujo/min-ecommerce/service/auth"
	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/gorilla/mux"
)

var logger = utils.NewParentLogger("Rota api/v1/products")

type Handler struct {
	Service types.ProductService
	store   repository.Queries
}

func NewHandler(Service types.ProductService, store repository.Queries) *Handler {
	return &Handler{
		Service: Service,
		store:   store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", auth.WithJWTAuth(h.CreateProduct, h.store, "admin")).Methods(http.MethodPost)
	router.HandleFunc("/get-all-products", h.GetAllProducts).Methods(http.MethodGet)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path, "Creating product")

	var payload types.CreateProductPayload

	logger.Info("Parsing JSON")
	if err := utils.ParseJSON(r, &payload); err != nil {
		logger.LogError(r.URL.Path, err, "Erro ao parsear o JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := h.Service.CreateProduct(&payload)
	if err != nil {
		logger.LogError(r.URL.Path, err)
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, status, map[string]string{"response": "product created"})

}

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path, "Get All products")

	query := r.URL.Query()
	pageStr := query.Get("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		logger.LogError(r.URL.Path, err, "Invalid or missing 'page' parameter")
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid or missing 'page' parameter"))
		return
	}

	payload := types.GetAllProductsPayload{
		Page: page,
	}

	logger.Info("Parsing JSON")
	if err := utils.ParseJSON(r, &payload); err != nil {
		logger.LogError(r.URL.Path, err, "parsing json")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	products, status, err := h.Service.GetAllProducts(&payload)
	if err != nil {
		logger.LogError(r.URL.Path, err)
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, status, products)

}
