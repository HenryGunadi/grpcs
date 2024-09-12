package types

type OrderPayload struct {
	OrderID   int32 `json:"orderID"`
	ProductID int32 `json:"productID"`
	Quantity  int32 `json:"quantity"`
	Status    bool  `json:"status"`
}