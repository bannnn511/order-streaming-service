package domain

import (
	"bytes"
	"encoding/gob"
)

type OrderBean struct {
	Id         string  `json:"id" example:"1" validate:"required"`
	CustomerId string  `json:"customer_id" example:"2" validate:"required"`
	Quantity   int     `json:"quantity" example:"10" validate:"required"`
	Price      float32 `json:"price" example:"2.5" validate:"required"`
}

func (ob OrderBean) ToByte() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(ob); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
