package data

type UserInterface interface {
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetOne(id int) (*User, error)
	Update(u *User) error
	Delete() error
	DeleteByID(id int) error
	Insert(u *User) (int, error)
	ResetPassword(pass string) error
	PasswordMatches(pt string) (bool, error)
}

type PlanInterface interface {
	GetAll() ([]*Plan, error)
	GetOne(id int) (*Plan, error)
	SubscribeUserToPlan(user User, plan Plan) error
	AmountForDisplay() string
}
