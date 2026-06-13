package car

import (
	"context"

	"github.com/whynullname/taxidrive/internal/domain"
)

type UseCase interface {
	AddCar(ctx context.Context, car *domain.Car) error
	GetAllCars(ctx context.Context) ([]domain.Car, error)
}

type useCase struct {
	carRepository Repository
}

func NewUseCase(carRepository Repository) UseCase {
	return &useCase{carRepository: carRepository}
}

func (u *useCase) AddCar(ctx context.Context, car *domain.Car) error {
	err := u.carRepository.Create(ctx, car)
	if err != nil {
		//TODO: добавить проверку на разные sql error
		return err
	}

	return nil
}

func (u *useCase) GetAllCars(ctx context.Context) ([]domain.Car, error) {
	cars, err := u.carRepository.GetAllCars(ctx)
	if err != nil {
		//TODO: добавить проверку на разные sql error
		return nil, err
	}

	return cars, nil
}
