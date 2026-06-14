package domain

import "github.com/google/uuid"

type CarStatus string

const (
	Free       CarStatus = "free"
	Rented     CarStatus = "rented"
	Repair     CarStatus = "repair"
	Inspection CarStatus = "inspection"
)

type Car struct {
	ID          uuid.UUID
	Brand       string
	NumberPlate string
	Status      CarStatus
}
