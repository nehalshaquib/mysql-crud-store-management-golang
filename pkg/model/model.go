package model

type Item struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Type      string `json:"type" db:"name"`
	Price     int    `json:"price" db:"name"`
	Available bool   `json:"available" db:"name"`
	Quantity  int    `json:"quantity" db:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
type SuccessResponse struct {
	Detail string `json:"detail"`
}
