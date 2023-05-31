package v1

import (
	"github.com/gofiber/fiber/v2"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(ctx *fiber.Ctx) error {
	return nil
}

func (t Tag) List(ctx *fiber.Ctx) error {
	return nil
}

func (t Tag) Create(ctx *fiber.Ctx) error {
	return nil
}

func (t Tag) Update(ctx *fiber.Ctx) error {
	return nil
}

func (t Tag) Delete(ctx *fiber.Ctx) error {
	return nil
}
