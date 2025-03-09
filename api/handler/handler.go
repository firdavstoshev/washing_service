package handler

import (
	"errors"
	"net/http"

	"github.com/firdavstoshev/washing_service/internal/domain"
	"github.com/firdavstoshev/washing_service/internal/dto"
	"github.com/firdavstoshev/washing_service/internal/service"
	"github.com/firdavstoshev/washing_service/internal/storage"
	"github.com/firdavstoshev/washing_service/pkg/errs"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	rep storage.IStorage
	svc service.IService
}

func NewHandler(rep storage.IStorage, svc service.IService) *Handler {
	return &Handler{rep: rep, svc: svc}
}

// @Summary Get all washing services
// @Summary Get all washing services
// @Description Get a list of all washing services
// @Tags services
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.ServiceDTO
// @Failure 500 {object} dto.ErrorResponse
// @Router /services [get]
func (h *Handler) GetServices(c *gin.Context) {
	services, err := h.rep.WashingService().GetWashingServices()
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	serviceDTOs := make([]dto.ServiceDTO, len(services))
	for i, svc := range services {
		serviceDTOs[i] = dto.ServiceDTO{
			ID:     svc.ID,
			Name:   svc.Name,
			TypeID: svc.TypeID,
			Type: dto.ServiceTypeDTO{
				ID:   svc.Type.ID,
				Name: svc.Type.Name,
				UnitType: dto.UnitTypeDTO{
					ID:   svc.Type.UnitTypeID,
					Name: svc.Type.UnitType.Name,
				},
			},
			UnitPrice: svc.UnitPrice,
		}
	}

	responseJSON(c, http.StatusOK, serviceDTOs)
}

// @Summary Create a new order
// @Description Create an order with provided customer and service data
// @Tags orders
// @Accept json
// @Produce json
// @Param order body dto.OrderPriceRequest true "Order data"
// @Success 201 {object} dto.CreateOrderResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	var req dto.OrderPriceRequest
	if err := c.BindJSON(&req); err != nil {
		errorJSON(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	// TODO: customerID необходимо получать из контекста запроса (токена)

	order := domain.NewOrder(
		req.CustomerID,
		req.IsChildItems,
		req.Express,
		req.WaitDays,
	)

	serviceItems := make([]domain.ServiceItem, len(req.Services))
	for i, item := range req.Services {
		serviceItems[i] = domain.ServiceItem{
			ServiceID: item.ServiceID,
			Amount:    item.Quantity,
		}
	}

	orderId, err := h.svc.Order().CreateOrder(order, &serviceItems)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrCustomerNotFound):
			errorJSON(c, http.StatusNotFound, "Customer not found")
		case errors.Is(err, errs.ErrWashingServiceNotFound):
			errorJSON(c, http.StatusNotFound, "Washing service not found")
		default:
			errorJSON(c, http.StatusInternalServerError, "Failed to calculate order price")
		}
		return
	}

	responseJSON(c, http.StatusCreated, dto.CreateOrderResponse{OrderID: orderId})
}

// @Summary Get order price
// @Description Get order price based on provided customer and service data
// @Tags orders
// @Accept json
// @Produce json
// @Param order body dto.OrderPriceRequest true "Order data"
// @Success 201 {object} dto.OrderPriceResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /order-price [post]
func (h *Handler) OrderPrice(c *gin.Context) {
	var req dto.OrderPriceRequest
	if err := c.BindJSON(&req); err != nil {
		errorJSON(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	order := domain.NewOrder(
		req.CustomerID,
		req.IsChildItems,
		req.Express,
		req.WaitDays,
	)

	serviceItems := make([]domain.ServiceItem, len(req.Services))
	for i, item := range req.Services {
		serviceItems[i] = domain.ServiceItem{
			ServiceID: item.ServiceID,
			Amount:    item.Quantity,
		}
	}

	orderPrice, err := h.svc.Order().OrderPrice(order, &serviceItems)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrCustomerNotFound):
			errorJSON(c, http.StatusNotFound, "Customer not found")
		case errors.Is(err, errs.ErrWashingServiceNotFound):
			errorJSON(c, http.StatusNotFound, "Washing service not found")
		default:
			errorJSON(c, http.StatusInternalServerError, "Failed to calculate order price")
		}
		return
	}

	responseJSON(c, http.StatusOK, dto.OrderPriceResponse{Price: orderPrice})
}
