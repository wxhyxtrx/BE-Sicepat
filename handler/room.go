package handler

import (
	"encoding/json"
	"net/http"
	dto "server/dto/result"
	roomchatdto "server/dto/roomchat"
	"server/model"
	"server/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerRoom struct {
	RoomRepository repositories.RoomRepository
}

func HandlerRoom(roomRepository repositories.RoomRepository) *handlerRoom {
	return &handlerRoom{roomRepository}
}

func (h *handlerRoom) CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userLogin := int(userInfo["id"].(float64))

	request := new(roomchatdto.RequestRoomChat)
	err := json.NewDecoder(r.Body).Decode(&request)
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

	dataRoom := model.Room{
		Name:    request.Nameroom,
		AdminID: userLogin,
	}

	room, err := h.RoomRepository.CreateRoom(dataRoom)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respose := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Failed Create Room"}
		json.NewEncoder(w).Encode(respose)
		return
	}

	data, _ := h.RoomRepository.GetRoom(room.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerRoom) GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	idRoom, _ := strconv.Atoi(mux.Vars(r)["id"])
	room, err := h.RoomRepository.GetRoom(idRoom)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: room}
	json.NewEncoder(w).Encode(response)
}
