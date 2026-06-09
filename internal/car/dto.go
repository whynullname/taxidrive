package car

type AddCarRequest struct {
	Id          string `json:"id"`
	Brand       string `json:"brand"`
	NumberPlate string `json:"number_plate"`
}

type GetCarsResponse struct {
	Cars []Car `json:"cars"`
}
