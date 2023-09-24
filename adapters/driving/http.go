package driving

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"hexarch/core/domain"
	"hexarch/core/ports"
	"hexarch/core/ports/inbound"
)

func SetupUserHTTPAdapter(usrSrv inbound.UserService, r fiber.Router) {
	r.Get("/:id<guid>", func(c *fiber.Ctx) error {
		usr, err := usrSrv.GetUserByID(c.Params("id"))
		if err != nil {
			if errors.Is(err, ports.ErrRecordNotFound) {
				return fiber.ErrNotFound
			}

			return fiber.ErrInternalServerError
		}

		return c.JSON(usr)
	})

	r.Get("/:username<minLen(2);maxLen(32)", func(c *fiber.Ctx) error {
		usr, err := usrSrv.GetUserByUsername(c.Params("username"))
		if err != nil {
			if errors.Is(err, ports.ErrRecordNotFound) {
				return fiber.ErrNotFound
			}

			return fiber.ErrInternalServerError
		}

		return c.JSON(usr)
	})

	r.Post("/", func(c *fiber.Ctx) error {
		var input domain.CreateUserInput
		if err := c.BodyParser(&input); err != nil {
			return fiber.ErrBadRequest
		}

		usr, err := usrSrv.CreateUser(input)
		if err != nil {
			if errors.Is(err, ports.ErrDuplicatedKey) {
				return fiber.ErrConflict
			}

			return fiber.ErrInternalServerError
		}

		return c.JSON(usr)
	})
}
