package car

import (
	"context"
	"database/sql"

	"github.com/whynullname/taxidrive/internal/domain"
)

type CarRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (c *CarRepository) AddCar(ctx context.Context, car *domain.Car) error {
	_, err := c.db.ExecContext(ctx, `INSERT INTO cars (id, brand, number_plate, status) VALUES ($1, $2, $3, $4)`,
		car.ID,
		car.Brand,
		car.NumberPlate,
		car.Status,
	)

	return err
}

func (c *CarRepository) GetAllCars(ctx context.Context) ([]domain.Car, error) {
	rows, err := c.db.QueryContext(ctx, `SELECT id, brand, number_plate, status FROM cars`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	cars := make([]domain.Car, 0)

	for rows.Next() {
		var car domain.Car
		err = rows.Scan(&car.ID, &car.Brand, &car.NumberPlate, &car.Status)
		if err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}
