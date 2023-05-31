package app

import (
	"github.com/azusachino/ficus/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func GetPageIndex(ctx *fiber.Ctx) int {
	return util.StrTo(ctx.Query("page_index", "1")).MustInt()
}

func GetPageSize(ctx *fiber.Ctx) int {
	return util.StrTo(ctx.Query("page_size", "10")).MustInt()
}

func GetPageOffset(page, pageSize int) int {
	if page <= 0 {
		page = 1
	}
	return (page - 1) * pageSize
}
