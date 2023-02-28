package beans

type OrderBean struct {
	Id         string  `json:"id" example:"1" validate:"required"`
	CustomerId string  `json:"customer_id" example:"2" validate:"required"`
	Quantity   int     `json:"quantity" example:"10" validate:"required"`
	Price      float32 `json:"price" example:"2.5" validate:"required"`
}
