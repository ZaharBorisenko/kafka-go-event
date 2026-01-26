package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"producer/command"
	"producer/services"
)

type AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithDrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(service services.AccountService) AccountController {
	return accountController{accountService: service}
}

func (a accountController) OpenAccount(c *fiber.Ctx) error {
	command := command.OpenAccountCommand{}

	if err := c.BodyParser(&command); err != nil {
		log.Println(err)
		return err
	}

	id, err := a.accountService.OpenAccount(command)

	if err != nil {
		log.Println(err)
		return err
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "open account success",
		"id":      id,
	})

}

func (a accountController) DepositFund(c *fiber.Ctx) error {
	command := command.DepositFundCommand{}

	if err := c.BodyParser(&command); err != nil {
		log.Println(err)
		return err
	}

	if err := a.accountService.DepositFund(command); err != nil {
		log.Println(err)
		return err
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "deposit fund success",
	})

}

func (a accountController) WithDrawFund(c *fiber.Ctx) error {
	command := command.WithDrawFundCommand{}

	if err := c.BodyParser(&command); err != nil {
		log.Println(err)
		return err
	}

	if err := a.accountService.WithDrawFund(command); err != nil {
		log.Println(err)
		return err
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "withdraw fund success",
	})
}

func (a accountController) CloseAccount(c *fiber.Ctx) error {
	command := command.CloseAccountCommand{}

	if err := c.BodyParser(&command); err != nil {
		log.Println(err)
		return err
	}

	if err := a.accountService.CloseAccount(command); err != nil {
		log.Println(err)
		return err
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "close account success",
	})

}
