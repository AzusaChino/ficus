package app

import (
	"net/http"

	"github.com/azusachino/ficus/pkg/errcode"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Ctx *fiber.Ctx
}

type Pagination struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
}

func NewResponse(ctx *fiber.Ctx) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) error {
	return r.Ctx.Status(http.StatusOK).JSON(data)
}

func (r *Response) ToResponsePage(list interface{}, total int) {
	r.Ctx.Status(http.StatusOK).JSON(fiber.Map{
		"list": list,
		"result": Pagination{
			PageIndex: GetPageIndex(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			Total:     total,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := fiber.Map{
		"code": err.Code(),
		"msg":  err.Msg(),
	}
	if err.Details() != nil {
		response["details"] = err.Details()
	}
	r.Ctx.Status(err.StatusCode()).JSON(response)
}
