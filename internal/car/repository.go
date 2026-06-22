package car

import (
	"context"

	"github.com/google/uuid"
	"github.com/whynullname/taxidrive/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, car *domain.Car) error
	GetAllCars(ctx context.Context) ([]domain.Car, error)
	GetCar(context.Context, uuid.UUID) (domain.Car, error)
	UpdateCar(context.Context, *domain.Car) error
	DeleteCar(context.Context, uuid.UUID) error
}
