package car

import (
	"context"

	"github.com/whynullname/taxidrive/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, car *domain.Car) error
	GetAllCars(ctx context.Context) ([]domain.Car, error)
}
