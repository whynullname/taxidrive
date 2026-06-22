package car

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/whynullname/taxidrive/internal/domain"
)

type CarRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (c *CarRepository) CreateCar(ctx context.Context, car *domain.Car) error {
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

func (c *CarRepository) GetCar(ctx context.Context, id uuid.UUID) (domain.Car, error) {
	row := c.db.QueryRowContext(ctx, `SELECT brand, number_plate, status FROM cars WHERE id=$1`, id)

	output := domain.Car{}
	var brand, numberPlate, status string
	err := row.Scan(&brand, &numberPlate, &status)
	if err != nil {
		return output, err
	}

	output.ID = id
	output.Brand = brand
	output.NumberPlate = numberPlate
	output.Status = domain.CarStatus(status)

	return output, nil
}

func (c *CarRepository) UpdateCar(ctx context.Context, car *domain.Car) error {
	res, err := c.db.ExecContext(ctx, `UPDATE cars SET brand = $1, number_plate = $2 WHERE id=$3`,
		car.Brand,
		car.NumberPlate,
		car.ID)

	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
func (c *CarRepository) DeleteCar(ctx context.Context, id uuid.UUID) error {
	res, err := c.db.ExecContext(ctx, `DELETE FROM cars WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
