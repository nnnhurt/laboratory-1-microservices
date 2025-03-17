package main

import (
	"log"
	"s0709-22/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	api := fiber.New()

	svc := service.New()

	api.Get("/api/v1/contact/:id", svc.ReadContact)
	api.Post("/api/v1/contact", svc.CreateContact)
	api.Put("/api/v1/contact/:id", svc.UpdateContact)
	api.Delete("/api/v1/contact/:id", svc.DeleteContact)
	api.Get("/api/v1/group/:id", svc.ReadGroup)
	api.Post("/api/v1/group", svc.CreateGroup)
	api.Put("/api/v1/group/:id", svc.UpdateGroup)
	api.Delete("/api/v1/group/:id", svc.DeleteGroup)
	api.Get("/api/v1/contact/all", svc.ReadAllContacts)
	api.Get("/api/v1/group/all", svc.ReadAllGroups)

	if err := api.Listen(":7080"); err != nil {
		log.Fatalln(err)
	}
}
