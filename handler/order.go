package handler

import (
	"encoding/json"
	"net/http"
	orderdto "server/dto/order"
	dto "server/dto/result"
	"server/model"
	"server/repositories"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
}

func HandlerOrder(orderRepository repositories.OrderRepository) *handlerOrder {
	return &handlerOrder{orderRepository}
}

func (h *handlerOrder) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userLogin := int(userInfo["id"].(float64))

	request := new(orderdto.RequestOrder)
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "cek dto Order"}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	tiket, err := h.OrderRepository.CheckTiket(request.Tiket)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Tiket yang dibeli tidak ada"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if tiket.Stok == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Mohon maaf Tiket telah habis"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if tiket.Stok < request.Qty {
		stok := strconv.Itoa(tiket.Stok)
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Mohon maaf sisa tiket : " + stok + " pcs"}
		json.NewEncoder(w).Encode(response)
		return
	}

	total := tiket.Harga * request.Qty

	item := model.Order{
		TiketID: request.Tiket,
		Qty:     request.Qty,
		Total:   total,
		Tanggal: time.Now(),
		UserID:  userLogin,
	}

	order, err := h.OrderRepository.CreateOrder(item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Gagal Order!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, _ := h.OrderRepository.GetOrder(order.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
