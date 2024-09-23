package repository

import (
	"database/sql"
	"fmt"
	"log"
	"notes/internal/entity"
)

type NoteRepository interface {
	GetAllNote() ([]entity.Note, error)
	GetByID(id int) (entity.Note, error)
	CreateNote(note *entity.Note) error
	DeleteNote(id int) error
	UpdateNote(note *entity.Note) error
}

type noteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) *noteRepository {
	return &noteRepository{db}
}

func (r *noteRepository) GetAllNote() ([]entity.Note, error) {
	var notes []entity.Note

	rows, err := r.db.Query("SELECT * FROM notes")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var note entity.Note

		if err := rows.Scan(&note.ID, &note.Title, &note.Description); err != nil {
			panic(err)
		}

		notes = append(notes, note)
	}

	return notes, err
}

func (r *noteRepository) GetByID(id int) (entity.Note, error) {
	var note entity.Note

	row := r.db.QueryRow("SELECT * FROM notes WHERE id = ? ", id)

	err := row.Scan(&note.ID, &note.Title, &note.Description)

	if err != nil {
		log.Printf("error, note with id : %d is not found", id)
	}
	return note, err
}

func (r *noteRepository) CreateNote(note *entity.Note) error {
	_, err := r.db.Exec("INSERT INTO notes (id,title,description) VALUES (?,?,?)", note.ID, note.Title, note.Description)

	if err != nil {
		return fmt.Errorf("create note : %v", err)
	}

	return err
}

func (r *noteRepository) DeleteNote(id int) error {
	_, err := r.db.Exec("DELETE FROM notes WHERE id = ?", id)

	if err != nil {
		return fmt.Errorf("failed to delete note : %v", err)
	}

	return err
}

func (r *noteRepository) UpdateNote(note *entity.Note) error {
	_, err := r.db.Exec("UPDATE notes n SET n.title = ? , n.description = ? WHERE n.id = ?", note.Title, note.Description, note.ID)

	if err != nil {
		return fmt.Errorf("failed to update note : %v", err)
	}

	return err
}
