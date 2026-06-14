package car

import (
	"context"

	"github.com/google/uuid"
	"github.com/whynullname/taxidrive/internal/domain"
	"github.com/whynullname/taxidrive/internal/logger"
)

type UseCase interface {
	CreateCar(context.Context, CreateCarInput) error
	GetAllCars(context.Context) ([]domain.Car, error)
	GetCar(context.Context, uuid.UUID) (domain.Car, error)
	UpdateCar(context.Context, UpdateCarInput) error
	DeleteCar(context.Context, uuid.UUID) error
}

type useCase struct {
	carRepository Repository
}

func NewUseCase(carRepository Repository) UseCase {
	return &useCase{carRepository: carRepository}
}

func (u *useCase) CreateCar(ctx context.Context, createCarInput CreateCarInput) error {
	carId, err := uuid.NewV6()
	if err != nil {
		logger.Instance.Errorln(err)
		return err
	}

	car := &domain.Car{
		ID:          carId,
		Brand:       createCarInput.Brand,
		NumberPlate: createCarInput.NumberPlate,
		Status:      domain.Free,
	}

	err = u.carRepository.Create(ctx, car)
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

func (u *useCase) GetCar(ctx context.Context, id uuid.UUID) (domain.Car, error) {
	return domain.Car{}, nil
}

func (u *useCase) UpdateCar(ctx context.Context, updateCarInput UpdateCarInput) error {
	return nil
}
func (u *useCase) DeleteCar(ctx context.Context, id uuid.UUID) error {
	return nil
}
