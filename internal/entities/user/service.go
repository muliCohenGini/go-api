package user

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID             int              `json:"id"`
	FirstName      *string          `json:"firstName"`
	LastName       *string          `json:"lastName"`
	Nickname       *string          `json:"nickname" validate:"required,min=3,max=15"`
	Gender         *string          `json:"gender"`
	DOB            *string          `json:"dob"`
	Carrier        *string          `json:"carrier"`
	Email          *string          `json:"email"`
	Phone          *string          `json:"phone"`
	IdentityNumber *string          `json:"identityNumber" validate:"required,min=6,max=9"`
	JoinDate       *string          `json:"joinDate"`
	Interests      *json.RawMessage `json:"interests"`
	Preferences    *json.RawMessage `json:"preferences"`
	Rank           *string          `json:"rank"`
	ImageURL       *string          `json:"imageUrl"`
	IsUnderage     *bool            `json:"isUnderage"`
	Points         *int             `json:"points"`
}

type Interest struct {
	ID       *int             `json:"id,omitempty"`
	Title    *string          `json:"title"`
	ImageURL *string          `json:"imageurl"`
	Feedback *json.RawMessage `json:"feedback"`
}

type UserService interface {
	GetUser(id string) (User, error)
	GetUsers() ([]User, error)
	SetUserOnboardingDetails(user User) error
	GetNicknameValidation(nickname string) (bool, error)
	GetInterests() ([]Interest, error)
	UpdateInterests(interest Interest, id string) (bool, error)
	GetIdentityNumberValidation(identityNumber string) (bool, error)
}

type Service struct {
	Repository UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &Service{Repository: repo}
}

func (s *Service) GetUser(id string) (User, error) {
	return s.Repository.GetUser(id)
}

func (s *Service) GetUsers() ([]User, error) {
	return s.Repository.GetUsers()
}

func (s *Service) SetUserOnboardingDetails(user User) error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		err := err.(validator.ValidationErrors)
		fmt.Printf("Validation error: %s", err)
		return err
	}

	return s.Repository.SetUserOnboardingDetails(user)
}

func (s *Service) GetNicknameValidation(nickname string) (bool, error) {
	validate := validator.New()
	err := validate.Var(nickname, "min=2,max=15")
	if err != nil {
		return false, errors.New("nickname not valid")
	}
	return s.Repository.GetNicknameValidation(nickname)
}

func (s *Service) GetInterests() ([]Interest, error) {
	return s.Repository.GetInterests()
}

func (s *Service) UpdateInterests(interest Interest, id string) (bool, error) {
	return s.Repository.UpdateInterests(interest, id)
}

func (s *Service) GetIdentityNumberValidation(identityNumber string) (bool, error) {
	validate := validator.New()
	err := validate.Var(identityNumber, "len=9,numeric")
	if err != nil {
		return false, errors.New("identity number not valid")
	}
	return s.Repository.GetIdentityNumberValidation(identityNumber)
}
