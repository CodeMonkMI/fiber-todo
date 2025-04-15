package todo

import "github.com/gofiber/fiber/v2"

func getAll(h *fiber.Ctx) error {
	return h.JSON(fiber.Map{"message": "Success"})
}
func single(h *fiber.Ctx) error {
	return h.JSON(fiber.Map{"message": "Success"})
}
func create(h *fiber.Ctx) error {
	return h.JSON(fiber.Map{"message": "Success"})
}
func update(h *fiber.Ctx) error {
	return h.JSON(fiber.Map{"message": "Success"})
}
func remove(h *fiber.Ctx) error {
	return h.JSON(fiber.Map{"message": "Success"})
}
