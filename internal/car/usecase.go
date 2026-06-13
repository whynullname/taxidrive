package car

import "github.com/whynullname/taxidrive/internal/domain"

type UseCase interface {
	AddCar(car *domain.Car) error
	GetAllCars() ([]domain.Car, error)
}

type useCase struct{}

func (u *useCase) AddCar(car *domain.Car) error {
	return nil
}

func (u *useCase) GetAllCars() ([]domain.Car, error) {
	return nil, nil
}
