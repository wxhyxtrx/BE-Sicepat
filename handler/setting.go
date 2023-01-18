package handler

import (
	"encoding/json"
	"net/http"
	dto "server/dto/result"
	settingdto "server/dto/setting"
	"server/model"
	"server/repositories"
	"strings"

	"github.com/go-playground/validator/v10"
)

type handlerSetting struct {
	SettingRepository repositories.SettingRepository
}

func HandlerSetting(settingRepository repositories.SettingRepository) *handlerSetting {
	return &handlerSetting{settingRepository}
}

func (h *handlerSetting) CreateSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	request := new(settingdto.RequestSetting)
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

	teks := strings.ToUpper(request.Teks)

	setting := model.Setting{
		Teks: teks,
	}

	datachat, err := h.SettingRepository.CreateSetting(setting)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respose := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Failed Create Chat"}
		json.NewEncoder(w).Encode(respose)
		return
	}

	data, _ := h.SettingRepository.GetSetting(datachat.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
