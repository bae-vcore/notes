package handler

import (
	"log"
	"notes/internal/entity"
	"notes/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type noteHandler struct {
	noteService services.NoteService
}

func NewNoteHandler(noteService services.NoteService) *noteHandler {
	return &noteHandler{noteService}
}

func (h *noteHandler) GetAllNote(c *fiber.Ctx) error {
	notes, err := h.noteService.GetAllNote()

	if err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "message": "failed to get all notes", "data": nil})
	}

	return c.JSON(fiber.Map{"success": true, "message": "successfully get all notes", "data": notes})
}

func (h *noteHandler) GetNoteById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	note, err := h.noteService.GetByID(id)

	if err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "message": "failed to get note with id", "data": nil})
		return nil
	}

	return c.JSON(fiber.Map{"success": true, "message": "successfully get note", "data": note})
}

func (h *noteHandler) CreateNote(c *fiber.Ctx) error {

	newNote := new(entity.Note)

	if err := c.BodyParser(newNote); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	if newNote.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "title is required",
		})

	}

	if newNote.Description == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "description is required",
		})
	}

	err := h.noteService.CreateNote(newNote)

	if err != nil {
		c.JSON(fiber.Map{"success": false, "message": "failed to create note", "data": nil})
		return nil
	}

	return c.Status(201).JSON(fiber.Map{"success": true, "message": "successfully created a new note"})
}

func (h *noteHandler) DeleteNote(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := h.noteService.DeleteNote(id)

	if err != nil {
		c.Status(500).JSON(fiber.Map{"success": false, "message": "failed to delete note with id"})
		return nil
	}

	return c.JSON(fiber.Map{"success": true, "message": "successfully delete note"})
}

func (h *noteHandler) UpdateNote(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	updateNote := new(entity.Note)
	updateNote.ID = id

	if err := c.BodyParser(updateNote); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	err := h.noteService.UpdateNote(updateNote)

	if err != nil {
		log.Printf(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to update note",
		})
	}

	return c.JSON(fiber.Map{"success": true, "message": "successfully update note", "data": updateNote})

}
