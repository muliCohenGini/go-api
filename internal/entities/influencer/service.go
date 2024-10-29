package influencer

import (
	"encoding/json"
)

type InfluencerService interface {
	GetInfluencers() ([]Influencer, error)
}

type Service struct {
	Repository InfluencerRepository
}

type Influencer struct {
	ID      int              `json:"id"`
	Name    *string          `json:"name"`
	Content *json.RawMessage `json:"content"`
}

type InfluencerRespons struct {
	Influencers []Influencer `json:"influencers"`
}

func NewInfluencerService(repo InfluencerRepository) InfluencerService {
	return &Service{Repository: repo}
}

func (s *Service) GetInfluencers() ([]Influencer, error) {
	return s.Repository.GetInfluencers()
}
