package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Counter struct {
	value int
}

func (cn *Counter) GetValue(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprint(cn.value))
}

func (c *Counter) Increase() int {
	c.value++
    return c.value
}
func (c *Counter) Decrease() int {
	c.value--
    return c.value
}
