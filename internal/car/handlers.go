package car

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/whynullname/taxidrive/internal/logger"
)

type Handlers interface {
	CreateCar(w http.ResponseWriter, r *http.Request)
	GetAllCars(w http.ResponseWriter, r *http.Request)
	GetCar(w http.ResponseWriter, r *http.Request)
	UpdateCar(w http.ResponseWriter, r *http.Request)
	DeleteCar(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	carUseCase UseCase
}

func NewCarsHandlers(carUseCase UseCase) Handlers {
	return &handlers{carUseCase: carUseCase}
}

func (h *handlers) CreateCar(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "application/json") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Instance.Errorf("error while read body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := CreateCarRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		logger.Instance.Errorf("error while unmarshal create car request: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createCarInput := CreateCarInput{
		Brand:       request.Brand,
		NumberPlate: request.Brand,
	}
	err = h.carUseCase.CreateCar(r.Context(), createCarInput)
	if err != nil {
		logger.Instance.Errorf("error while create car: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handlers) GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.carUseCase.GetAllCars(r.Context())
	if err != nil {
		logger.Instance.Errorf("error while get all cars: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := GetCarsResponse{
		Cars: cars,
	}
	data, err := json.Marshal(response)
	if err != nil {
		logger.Instance.Errorf("can't marshal get cars response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		logger.Instance.Errorln("can't write get cars response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handlers) GetCar(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Instance.Warnf("invalid id in get car: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	car, err := h.carUseCase.GetCar(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := &GetCarResponse{
		Brand:       car.Brand,
		NumberPlate: car.NumberPlate,
	}
	data, err := json.Marshal(response)
	if err != nil {
		logger.Instance.Errorf("can't marshal get car response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		logger.Instance.Errorln("can't write get car response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handlers) UpdateCar(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Instance.Warnf("invalid id in update car: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Instance.Errorf("can't read body in update car: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := UpdateCarRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		logger.Instance.Errorf("error while unmarshal update car request: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := UpdateCarInput{
		Id:          id,
		Brand:       request.Brand,
		NumberPlate: request.NumberPlate,
	}
	err = h.carUseCase.UpdateCar(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handlers) DeleteCar(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Instance.Warnf("invalid id in update car: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.carUseCase.DeleteCar(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
