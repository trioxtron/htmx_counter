package api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Counter struct {
	value int
}

func (cn *Counter) GetValue(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprint(cn.value))
}

func (cn *Counter) SetValue(c *fiber.Ctx) error {
	value, err := strconv.Atoi(c.FormValue("value"))
	if err != nil {
		return c.SendString("value must be a number")
	}
	cn.value = value
	return c.SendString(fmt.Sprint(cn.value))
}

func (c *Counter) ReturnValue() int {
	return c.value
}

func (c *Counter) Increase() int {
	c.value++
	return c.value
}
func (c *Counter) Decrease() int {
	c.value--
	return c.value
}
