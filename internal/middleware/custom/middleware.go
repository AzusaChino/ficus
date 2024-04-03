package custom

import (
	"time"

	"github.com/azusachino/ficus/pkg/support"
	"github.com/gofiber/fiber/v2"
)

func NewRateLimiter(capacity int, refill time.Duration) func(*fiber.Ctx) error {
	rl := support.NewRateLimiter(capacity, refill)
	return func(c *fiber.Ctx) error {
		if !rl.Acquire() {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "too many requests",
			})
		}
		return c.Next()
	}
}
