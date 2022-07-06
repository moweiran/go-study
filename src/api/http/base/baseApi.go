package base

import "github.com/gofiber/fiber/v2"

type BaseApi struct {
}

func (base *BaseApi) DataOrError(c *fiber.Ctx, response interface{}, err error) error {
	if err != nil {
		return base.Error(c, err)
	}

	return base.Data(c, response, "success")
}

func (base *BaseApi) Error(c *fiber.Ctx, err error) error {
	return c.JSON(map[string]interface{}{
		"code": "999",
	})
}

func (base *BaseApi) Data(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(map[string]interface{}{
		"code":    "200",
		"data":    data,
		"message": message,
	})
}
