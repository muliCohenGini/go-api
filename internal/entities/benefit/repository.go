package benefit

import (
	"github.com/muliCohenGini/go-api/internal/db"
)

type BenefitRepository interface {
	GetBenefits(category string) ([]Benefit, error)
}

type Repository struct {
}

func NewBenefitRepository() BenefitRepository {
	return &Repository{}
}

func (r *Repository) GetBenefits(category string) ([]Benefit, error) {
	rows, err := db.DB.Query("SELECT * FROM \"Benefits\" WHERE category = $1", category)
	if err != nil {
		return []Benefit{}, err
	}
	defer rows.Close()
	benefits := []Benefit{}
	for rows.Next() {
		var benefit Benefit
		if err := rows.Scan(benefit.ScanFields()...); err != nil {
			return []Benefit{}, err
		}
		benefits = append(benefits, benefit)
	}
	if err := rows.Err(); err != nil {
		return []Benefit{}, err
	}
	return benefits, nil
}

func (b *Benefit) ScanFields() []interface{} {
	return []interface{}{
		&b.ID,
		&b.Category,
		&b.Information,
		&b.Monthly,
		&b.ExpirationDate,
	}
}
