package car

type UseCase interface {
	AddCar(car *Car) error
	GetAllCars() ([]Car, error)
}

type useCase struct{}

func (u *useCase) AddCar(car *Car) error {
	return nil
}

func (u *useCase) GetAllCars() ([]Car, error) {
	return nil, nil
}
