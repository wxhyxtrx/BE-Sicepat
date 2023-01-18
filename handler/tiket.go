package handler

import (
	"encoding/json"
	"net/http"
	dto "server/dto/result"
	tiketdto "server/dto/tiket"
	"server/model"
	"server/repositories"

	"github.com/go-playground/validator/v10"
)

type handlerTiket struct {
	TiketRepository repositories.TiketRepository
}

func HandlerTiket(tiketRepository repositories.TiketRepository) *handlerTiket {
	return &handlerTiket{tiketRepository}
}

func (h *handlerTiket) CreateTiket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(tiketdto.RequestTiket)
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cek dto request =>" + err.Error()}
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

	tiket := model.Tiket{
		Name:  request.Name,
		Stok:  request.Stok,
		Harga: request.Harga,
	}

	dataTiket, err := h.TiketRepository.CreateTiket(tiket)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respose := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Failed Create Tiket"}
		json.NewEncoder(w).Encode(respose)
		return
	}

	data, _ := h.TiketRepository.GetTiket(dataTiket.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
