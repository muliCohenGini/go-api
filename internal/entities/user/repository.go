package user

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/muliCohenGini/go-api/internal/db"
)

type UserRepository interface {
	GetUser(id string) (User, error)
	GetUsers() ([]User, error)
	SetUserOnboardingDetails(user User) error
	GetInterests() ([]Interest, error)
	GetNicknameValidation(nickname string) (bool, error)
	UpdateInterests(interest Interest, id string) (bool, error)
	GetIdentityNumberValidation(identityNumber string) (bool, error)
}

type Repository struct {
}

func NewUserRepository() UserRepository {
	return &Repository{}
}

func (r *Repository) GetUser(id string) (User, error) {
	var u User
	err := db.DB.QueryRow("SELECT * FROM \"Users\" WHERE id = $1", id).Scan(u.ScanFields()...)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *Repository) GetUsers() ([]User, error) {
	rows, err := db.DB.Query("SELECT * FROM \"Users\"")
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(u.ScanFields()...); err != nil {
			return []User{}, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return []User{}, err
	}
	return users, nil
}

func (r *Repository) SetUserOnboardingDetails(user User) error {
	query := `
		INSERT INTO "Users" (gender, "firstName", "lastName", "nickname", "dob", "identityNumber", "interests")
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	_, err := db.DB.Exec(query,
		user.Gender,
		user.FirstName,
		user.LastName,
		user.Nickname,
		user.DOB,
		user.IdentityNumber,
		user.Interests,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetNicknameValidation(nickname string) (bool, error) {
	var u User
	err := db.DB.QueryRow("SELECT nickname FROM \"Users\" WHERE nickname = $1", nickname).Scan(&u.Nickname)
	if err == sql.ErrNoRows {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, errors.New("user already exists")
}

func (r *Repository) GetInterests() ([]Interest, error) {
	rows, err := db.DB.Query("SELECT * FROM \"Interests\"")
	if err != nil {
		return []Interest{}, err
	}
	defer rows.Close()
	interests := []Interest{}
	for rows.Next() {
		var in Interest
		if err := rows.Scan(in.ScanFields()...); err != nil {
			return []Interest{}, err
		}
		interests = append(interests, in)
	}
	if err := rows.Err(); err != nil {
		return []Interest{}, err
	}
	return interests, nil
}

func (r *Repository) UpdateInterests(interest Interest, id string) (bool, error) {
	out, errr := json.Marshal(interest)
	if errr != nil {
		return false, errr
	}
	query := `UPDATE "Users" SET interests = $1 WHERE id = $2`
	_, err := db.DB.Query(query, out, id)
	if err != nil {
		return false, err
	}
	return true, err
}

func (r *Repository) GetIdentityNumberValidation(identityNumber string) (bool, error) {
	var u User
	err := db.DB.QueryRow(`SELECT * FROM "Users" WHERE "identityNumber" = $1`, identityNumber).Scan(&u.IdentityNumber)
	if err == sql.ErrNoRows {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, errors.New("identity number already exists")
}

func (u *User) ScanFields() []interface{} {
	return []interface{}{
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Nickname,
		&u.Gender,
		&u.DOB,
		&u.Carrier,
		&u.Email,
		&u.Phone,
		&u.IdentityNumber,
		&u.JoinDate,
		&u.Interests,
		&u.Preferences,
		&u.Rank,
		&u.ImageURL,
		&u.IsUnderage,
		&u.Points,
	}
}

func (i *Interest) ScanFields() []interface{} {
	return []interface{}{
		&i.ID,
		&i.Title,
		&i.ImageURL,
		&i.Feedback,
	}
}
