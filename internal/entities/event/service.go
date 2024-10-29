package event

import (
	"encoding/json"
)

type EventService interface {
	GetEvents() ([]Event, error)
}

type Service struct {
	Repository EventRepository
}

type Event struct {
	ID       int              `json:"id"`
	UserId   *string          `json:"userId"`
	Date     *string          `json:"date"`
	Rank     *string          `json:"rank"`
	Location *string          `json:"location"`
	Content  *json.RawMessage `json:"content"`
}

type EventRespons struct {
	Events []Event `json:"events"`
}

func NewEventsService(repo EventRepository) EventService {
	return &Service{Repository: repo}
}

func (s *Service) GetEvents() ([]Event, error) {
	return s.Repository.GetEvents()
}
