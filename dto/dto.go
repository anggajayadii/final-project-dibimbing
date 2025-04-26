package dto

type CreateAssetInput struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	PurchaseDate string `json:"purchase_date" binding:"required"`
	Location     string `json:"location" binding:"required"`
	Status       string `json:"status" binding:"required"`
	SerialNumber string `json:"serial_number" binding:"required"`
}
