package car

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Handlers interface {
	AddCar(w http.ResponseWriter, r *http.Request)
	GetAllCars(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	carUseCase UseCase
}

func (h *handlers) AddCar(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "application/json") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := AddCarRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	car := Car{
		Id:          request.Id,
		Brand:       request.Brand,
		NumberPlate: request.Brand,
		Status:      "free",
	}
	err = h.carUseCase.AddCar(&car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handlers) GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.carUseCase.GetAllCars()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := GetCarsResponse{
		Cars: cars,
	}
	data, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
