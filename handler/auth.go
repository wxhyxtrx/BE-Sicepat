package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	authdto "server/dto/auth"
	dto "server/dto/result"
	"server/model"
	"server/pkg/bcrypt"
	jwtToken "server/pkg/jwt"
	"server/repositories"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	dataUser, _ := h.AuthRepository.Login(request.Email)

	for _, data := range dataUser {
		if data.Email == request.Email {
			if data.Status != "AKTIF" {
				today := time.Now().Day()
				timeOff := data.TimeOFF.Day()

				dateTime := today - timeOff

				if dateTime <= 7 {
					w.WriteHeader(http.StatusBadRequest)
					response := dto.ErrorResult{Code: http.StatusBadRequest, Message: data.Status}
					json.NewEncoder(w).Encode(response)
					return
				}

				if dateTime > 7 {
					user := model.User{
						Username: request.Username,
						Email:    request.Email,
						Password: password,
						Status:   "AKTIF",
					}
					data, err := h.AuthRepository.Register(user)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
						json.NewEncoder(w).Encode(response)
					}

					registrasiResponse := authdto.RegistResponse{
						Username: data.Username,
						Email:    data.Email,
					}

					w.WriteHeader(http.StatusOK)
					response := dto.SuccessResult{Code: http.StatusOK, Data: registrasiResponse}
					json.NewEncoder(w).Encode(response)
					return
				}
			}
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "E-mail already registered"}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	user := model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		Status:   "AKTIF",
	}
	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	registrasiResponse := authdto.RegistResponse{
		Username: data.Username,
		Email:    data.Email,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: registrasiResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := model.User{
		Email:    request.Email,
		Password: request.Password,
	}

	data, err := h.AuthRepository.CekUser(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Check your email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, data.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Check your email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}
	claims := jwt.MapClaims{}
	claims["id"] = data.ID
	claims["email"] = data.Email
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)

	if errGenerateToken != nil {
		fmt.Println(errGenerateToken)
		return
	}

	loginResponse := authdto.LoginResponse{
		Username: data.Username,
		Email:    data.Email,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	user, err := h.AuthRepository.UserLogin(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "User Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userData := authdto.AuthResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: userData}
	json.NewEncoder(w).Encode(response)
}
