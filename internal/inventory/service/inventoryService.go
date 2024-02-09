package service

import (
	"github.com/COPPED/inventory-service/pkg/models"
	"github.com/google/uuid"
)

func FetchInventory(id string) *models.Inventory {
	return &models.Inventory{
		Id:     id,
		Size:   "XXL",
		ItemId: uuid.NewString(),
	}
}

func FetchInventoryList() *[]models.Inventory {
	return &[]models.Inventory{
		{
			Id:     "1",
			Size:   "XXL",
			ItemId: uuid.NewString(),
		},
		{
			Id:     "12",
			Size:   "XXL",
			ItemId: uuid.NewString(),
		},
	}
}
