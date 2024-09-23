package services

import (
	"notes/internal/entity"
	"notes/internal/repository"
)

type NoteService interface {
	GetAllNote() ([]entity.Note, error)
	GetByID(id int) (entity.Note, error)
	CreateNote(note *entity.Note) error
	DeleteNote(id int) error
	UpdateNote(note *entity.Note) error
}

type noteService struct {
	repository repository.NoteRepository
}

func NewNoteService(repository repository.NoteRepository) *noteService {
	return &noteService{repository}
}

func (s *noteService) GetAllNote() ([]entity.Note, error) {
	notes, err := s.repository.GetAllNote()

	return notes, err
}

func (s *noteService) GetByID(id int) (entity.Note, error) {
	note, err := s.repository.GetByID(id)

	return note, err
}

func (s *noteService) CreateNote(note *entity.Note) error {
	err := s.repository.CreateNote(note)

	return err
}

func (s *noteService) DeleteNote(id int) error {
	err := s.repository.DeleteNote(id)

	return err
}

func (s *noteService) UpdateNote(note *entity.Note) error {
	err := s.repository.UpdateNote(note)

	return err
}
