package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Contact struct {
	ID         int
	Username   string
	GivenName  string
	FamilyName string
	Phone      []struct {
		TypeID      int
		CountryCode int
		Operator    int
		Number      int
	}
	Email     []string
	Birthdate time.Time
}

type Group struct {
	ID          int
	Title       string
	Description string
	Contacts    []int
}

// Service ...
type Service struct {
	contacts []Contact
	groups   []Group
}

// New ...
func New() *Service {
	return &Service{
		groups:   []Group{},
		contacts: []Contact{},
	}
}

// CreateContact ...
func (s *Service) CreateContact(ctx *fiber.Ctx) error {
	var contact Contact
	if err := ctx.BodyParser(&contact); err != nil {
		return err
	}

	contact.ID = len(s.contacts) + 1

	s.contacts = append(s.contacts, contact)

	return ctx.JSON(contact)
}

// ReadAllContacts ...
func (s *Service) ReadAllContacts(ctx *fiber.Ctx) error {
	return ctx.JSON(s.contacts)
}

// ReadContact ...
func (s *Service) ReadContact(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	for _, contact := range s.contacts {
		if contact.ID == id {
			return ctx.JSON(contact)
		}
	}

	return fiber.ErrNotFound
}

// UpdateContact ...
func (s *Service) UpdateContact(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	var updatedContact Contact
	if err := ctx.BodyParser(&updatedContact); err != nil {
		return err
	}

	for i, contact := range s.contacts {
		if contact.ID == id {
			s.contacts[i] = updatedContact
			return ctx.JSON(updatedContact)
		}
	}

	return fiber.ErrNotFound
}

// DeleteContact ...
func (s *Service) DeleteContact(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	for i, contact := range s.contacts {
		if contact.ID == id {
			s.contacts = append(s.contacts[:i], s.contacts[i+1:]...)
			return ctx.SendStatus(fiber.StatusNoContent)
		}
	}

	return fiber.ErrNotFound
}

func (s *Service) CreateGroup(ctx *fiber.Ctx) error {
	var group Group
	if err := ctx.BodyParser(&group); err != nil {
		return err
	}

	group.ID = len(s.groups) + 1

	s.groups = append(s.groups, group)

	return ctx.JSON(group)
}

func (s *Service) ReadAllGroups(ctx *fiber.Ctx) error {
	return ctx.JSON(s.groups)
}

func (s *Service) ReadGroup(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	for _, group := range s.groups {
		if group.ID == id {
			return ctx.JSON(group)
		}
	}

	return fiber.ErrNotFound
}

// UpdateGroup ...
func (s *Service) UpdateGroup(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	var updatedGroup Group
	if err := ctx.BodyParser(&updatedGroup); err != nil {
		return err
	}

	for i, group := range s.groups {
		if group.ID == id {
			s.groups[i] = updatedGroup
			return ctx.JSON(updatedGroup)
		}
	}

	return fiber.ErrNotFound
}

// DeleteGroup ...
func (s *Service) DeleteGroup(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	for i, group := range s.groups {
		if group.ID == id {
			s.groups = append(s.groups[:i], s.groups[i+1:]...)
			return ctx.SendStatus(fiber.StatusNoContent)
		}
	}

	return fiber.ErrNotFound
}
