package benefit

import (
	"encoding/json"
)

type BenefitService interface {
	GetBenefits(category string) ([]Benefit, error)
}

type Service struct {
	Repository BenefitRepository
}

type Benefit struct {
	ID             int              `json:"id"`
	Category       *string          `json:"category"`
	Information    *json.RawMessage `json:"information"`
	Monthly        *string          `json:"monthly"`
	ExpirationDate *string          `json:"expirationDate"`
}

type BenefitRespons struct {
	Benefits []Benefit `json:"benefits"`
}

func NewBenefitService(repo BenefitRepository) BenefitService {
	return &Service{Repository: repo}
}

func (s *Service) GetBenefits(category string) ([]Benefit, error) {
	return s.Repository.GetBenefits(category)
}
