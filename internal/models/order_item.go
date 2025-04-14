package models

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
	Price     float64
}
