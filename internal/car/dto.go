package car

import "github.com/whynullname/taxidrive/internal/domain"

type AddCarRequest struct {
	Id          string `json:"id"`
	Brand       string `json:"brand"`
	NumberPlate string `json:"number_plate"`
}

type GetCarsResponse struct {
	Cars []domain.Car `json:"cars"`
}
