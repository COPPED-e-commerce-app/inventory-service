package rest

import (
	"github.com/COPPED/inventory-service/internal/inventory/service"
	"github.com/gofiber/fiber/v2"
)

func FetchInventory(c *fiber.Ctx) error {
	id := c.Params("id")
	inventory := service.FetchInventory(id)
	return c.JSON(inventory)
}

func FetchInventoryList(c *fiber.Ctx) error {
	inventoryList := service.FetchInventoryList()
	return c.JSON(inventoryList)
}
