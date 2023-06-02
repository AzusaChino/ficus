package v1

import (
	"github.com/azusachino/ficus/global"
	"github.com/azusachino/ficus/internal/service"
	"github.com/azusachino/ficus/pkg/app"
	"github.com/azusachino/ficus/pkg/errcode"
	"github.com/gofiber/fiber/v2"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Count( ctx *fiber.Ctx) error {
    param := service.CountTagRequest{}
    ctx.ParamsParser(&param)
    svc := service.New(ctx.Context())
    res := app.NewResponse(ctx)
    c, err := svc.CountTag(&param)
    if err != nil {
        global.Logger.Errorf("svc.CountTag err: %v", err)
        res.ToErrorResponse(errcode.InvalidParams)
        return err
    }
    res.ToResponse(c)
    return nil
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
