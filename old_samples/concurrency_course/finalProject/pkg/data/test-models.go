package data

import (
	"database/sql"
	"fmt"
	"time"
)

func TestNew(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: &UserTest{},
		Plan: &PlanTest{},
	}
}

type UserTest struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    int
	IsAdmin   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      *Plan
}

func (usr *UserTest) GetAll() ([]*User, error) {
	return []*User{
		{
			ID:        1,
			Email:     "admin@example.com",
			FirstName: "Admin",
			LastName:  "Admin",
			Password:  "abc",
			Active:    1,
			IsAdmin:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

func (usr *UserTest) GetByEmail(email string) (*User, error) {
	return &User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "abc",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (usr *UserTest) GetOne(id int) (*User, error) {
	return usr.GetByEmail("")
}

func (usr *UserTest) Update(u *User) error {
	return nil
}

func (usr *UserTest) Delete() error {
	return nil
}

func (usr *UserTest) DeleteByID(id int) error {
	return nil
}

func (usr *UserTest) Insert(u *User) (int, error) {
	return 1, nil
}

func (usr *UserTest) ResetPassword(pass string) error {
	return nil
}

func (usr *UserTest) PasswordMatches(pt string) (bool, error) {
	return true, nil
}

type PlanTest struct {
	ID                  int
	PlanName            string
	PlanAmount          int
	PlanAmountFormatted string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (p *PlanTest) GetAll() ([]*Plan, error) {
	return []*Plan{
		{
			ID:         1,
			PlanName:   "Plan",
			PlanAmount: 20,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}, nil
}

func (p *PlanTest) GetOne(id int) (*Plan, error) {
	return &Plan{
		ID:         1,
		PlanName:   "Plan",
		PlanAmount: 20,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func (p *PlanTest) SubscribeUserToPlan(user User, plan Plan) error {
	return nil
}

func (p *PlanTest) AmountForDisplay() string {
	return fmt.Sprintf("$%.2f", float64(p.PlanAmount)/100)
}
