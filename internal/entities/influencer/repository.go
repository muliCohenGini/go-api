package influencer

import (
	"github.com/muliCohenGini/go-api/internal/db"
)

type InfluencerRepository interface {
	GetInfluencers() ([]Influencer, error)
}

type Repository struct {
}

func NewInfluencerRepository() InfluencerRepository {
	return &Repository{}
}

func (r *Repository) GetInfluencers() ([]Influencer, error) {
	rows, err := db.DB.Query("SELECT * FROM \"Influencers\"")
	if err != nil {
		return []Influencer{}, err
	}
	defer rows.Close()
	Influencers := []Influencer{}
	for rows.Next() {
		var inf Influencer
		if err := rows.Scan(inf.ScanFields()...); err != nil {
			return []Influencer{}, err
		}
		Influencers = append(Influencers, inf)
	}
	if err := rows.Err(); err != nil {
		return []Influencer{}, err
	}
	return Influencers, nil
}

func (i *Influencer) ScanFields() []interface{} {
	return []interface{}{
		&i.ID,
		&i.Name,
		&i.Content,
	}
}
