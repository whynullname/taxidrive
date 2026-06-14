package car

import (
	"github.com/google/uuid"
	"github.com/whynullname/taxidrive/internal/domain"
)

type CreateCarRequest struct {
	Brand       string `json:"brand"`
	NumberPlate string `json:"number_plate"`
}

type UpdateCarRequest struct {
	Brand       string `json:"brand"`
	NumberPlate string `json:"number_plate"`
}

type GetCarsResponse struct {
	Cars []domain.Car `json:"cars"`
}

type GetCarResponse struct {
	Brand       string `json:"brand"`
	NumberPlate string `json:"number_plate"`
}

type CreateCarInput struct {
	Brand       string
	NumberPlate string
}

type UpdateCarInput struct {
	Id          uuid.UUID
	Brand       string
	NumberPlate string
}
