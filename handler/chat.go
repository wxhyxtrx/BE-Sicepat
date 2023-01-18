package handler

import (
	"encoding/json"
	"net/http"
	chatdto "server/dto/chat"
	dto "server/dto/result"
	"server/model"
	"server/repositories"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerChat struct {
	ChatRepository repositories.ChatRepository
}

func HandlerChat(chatRepository repositories.ChatRepository) *handlerChat {
	return &handlerChat{chatRepository}
}

func (h *handlerChat) CreateChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userLogin := int(userInfo["id"].(float64))

	request := new(chatdto.RequestChat)
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

	cekMessage := strings.Split(request.Message, " ")

	settings, err := h.ChatRepository.CheckSetting()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "setting not found!" + err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, teks := range cekMessage {
		teks = strings.ToUpper(teks)
		for _, chek := range settings {
			if strings.Contains(teks, chek.Teks) {
				dataUser, _ := h.ChatRepository.User(userLogin)

				if dataUser.Status == "AKTIF" {
					dataUser.Status = "1X PERINGATAN"
					data, _ := h.ChatRepository.SettingUser(dataUser)

					w.WriteHeader(http.StatusBadRequest)
					respose := dto.ErrorResult{Code: http.StatusBadRequest, Message: data.Status}
					json.NewEncoder(w).Encode(respose)
					return
				}
				if dataUser.Status == "1X PERINGATAN" {
					dataUser.Status = "2X PERINGATAN"
					data, _ := h.ChatRepository.SettingUser(dataUser)

					w.WriteHeader(http.StatusBadRequest)
					respose := dto.ErrorResult{Code: http.StatusBadRequest, Message: data.Status}
					json.NewEncoder(w).Encode(respose)
					return
				}
				if dataUser.Status == "2X PERINGATAN" {
					dataUser.Status = "3X PERINGATAN"
					data, _ := h.ChatRepository.SettingUser(dataUser)

					w.WriteHeader(http.StatusBadRequest)
					respose := dto.ErrorResult{Code: http.StatusBadRequest, Message: data.Status}
					json.NewEncoder(w).Encode(respose)
					return
				}
				if dataUser.Status == "3X PERINGATAN" {
					dataUser.Status = "Akun telah di matikan selama 7 Hari"
					dataUser.TimeOFF = time.Now()

					data, _ := h.ChatRepository.SettingUser(dataUser)

					w.WriteHeader(http.StatusBadRequest)
					respose := dto.ErrorResult{Code: http.StatusBadRequest, Message: data.Status}
					json.NewEncoder(w).Encode(respose)
					return
				}
				if dataUser.Status == "Akun telah di matikan selama 7 Hari" {
					w.WriteHeader(http.StatusBadRequest)
					respose := dto.ErrorResult{Code: http.StatusBadRequest, Message: dataUser.Status}
					json.NewEncoder(w).Encode(respose)
					return
				}
			}
		}
	}

	chat := model.Chat{
		Message:  request.Message,
		FromUser: userLogin,
		RoomID:   request.RoomID,
	}

	datachat, err := h.ChatRepository.CreateChat(chat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respose := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Failed Create Chat"}
		json.NewEncoder(w).Encode(respose)
		return
	}

	data, _ := h.ChatRepository.GetChat(datachat.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
