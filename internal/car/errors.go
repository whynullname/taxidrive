package car

import "errors"

var ErrInternalWhileCreateCar = errors.New("internal error while create car")
var ErrCarNotFound = errors.New("car not found")
var ErrInternalErrorWhileGetCar = errors.New("internal error while get car")
var ErrInternalWhileGetAllCars = errors.New("internal error while get all cars")
