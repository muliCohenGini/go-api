package event

import (
	"github.com/muliCohenGini/go-api/internal/db"
)

type EventRepository interface {
	GetEvents() ([]Event, error)
}

type Repository struct {
}

func NewEventRepository() EventRepository {
	return &Repository{}
}

func (r *Repository) GetEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT * FROM \"Events\"")
	if err != nil {
		return []Event{}, err
	}
	defer rows.Close()
	events := []Event{}
	for rows.Next() {
		var event Event
		if err := rows.Scan(event.ScanFields()...); err != nil {
			return []Event{}, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return []Event{}, err
	}
	return events, nil
}

func (e *Event) ScanFields() []interface{} {
	return []interface{}{
		&e.ID,
		&e.UserId,
		&e.Date,
		&e.Rank,
		&e.Location,
		&e.Content,
	}
}
