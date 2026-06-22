package car

import (
	"context"
	"database/sql"
	"errors"

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
		return ErrInternalWhileCreateCar
	}

	car := &domain.Car{
		ID:          carId,
		Brand:       createCarInput.Brand,
		NumberPlate: createCarInput.NumberPlate,
		Status:      domain.Free,
	}

	err = u.carRepository.Create(ctx, car)
	if err != nil {
		logger.Instance.Errorln(err)
		return ErrInternalWhileCreateCar
	}

	return nil
}

func (u *useCase) GetAllCars(ctx context.Context) ([]domain.Car, error) {
	cars, err := u.carRepository.GetAllCars(ctx)
	if err != nil {
		logger.Instance.Errorln(err)
		return nil, ErrInternalWhileGetAllCars
	}

	return cars, nil
}

func (u *useCase) GetCar(ctx context.Context, id uuid.UUID) (domain.Car, error) {
	car, err := u.carRepository.GetCar(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Car{}, ErrCarNotFound
		}
		logger.Instance.Errorln(err)
		return domain.Car{}, ErrInternalErrorWhileGetCar
	}
	return car, nil
}

func (u *useCase) UpdateCar(ctx context.Context, updateCarInput UpdateCarInput) error {
	car, err := u.GetCar(ctx, updateCarInput.Id)
	if err != nil {
		logger.Instance.Errorf("can't get car in db: %v\n", err)
		return err
	}

	car.Brand = updateCarInput.Brand
	car.NumberPlate = updateCarInput.NumberPlate
	err = u.carRepository.UpdateCar(ctx, &car)
	if err != nil {
		logger.Instance.Errorln(err)
		return err
	}

	return nil
}
func (u *useCase) DeleteCar(ctx context.Context, id uuid.UUID) error {
	err := u.DeleteCar(ctx, id)
	if err != nil {
		logger.Instance.Errorf("error while delete car %v\n", err)
		return err
	}

	return nil
}
