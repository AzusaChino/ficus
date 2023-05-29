package v1

import "github.com/gofiber/fiber/v2"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (Article) Get(ctx *fiber.Ctx) error {
	return nil
}
func (Article) List(ctx *fiber.Ctx) error {
	return nil
}
func (Article) Create(ctx *fiber.Ctx) error {
	return nil
}
func (Article) Update(ctx *fiber.Ctx) error {
	return nil
}
func (Article) Delete(ctx *fiber.Ctx) error {
	return nil
}
