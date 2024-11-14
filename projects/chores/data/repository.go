package data

import "chores_app/configs"

// PaginationOptions can be used to control the pagination
// of a collection of resources
type PaginationOptions struct {
	Page  int // default: 1
	Limit int // default: 10
}

type ChoresRepository interface {
	ListChores(pagination PaginationOptions) ([]*ChoreDefinition, error)
	CreateChore(details ChoreDetails, schedule Schedule) (*ChoreDefinition, error)
	GetChoreById(choreID string) (*ChoreDefinition, error)
	UpdateChoreDetails(choreID string, details ChoreDetails) (*ChoreDefinition, error)
	UpdateChoreSchedule(choreID string, schedule Schedule) (*ChoreDefinition, error)
	DeactivateChore(choreID string) error
}

type MembersFilterOptions struct {
	Role   Role
	Active bool
}

type MembersRepository interface {
	ListMembers(filter MembersFilterOptions, pagination PaginationOptions) ([]*Member, error)
	CreateMember(member *Member) (*Member, error)
	GetMemberById(memberID string) (*Member, error)
	UpdateMemberInfo(memberID string, info MemberInfo) (*Member, error)
	DeactivateMember(memberID string) error
}

type Repository interface {
	ChoresRepository
	MembersRepository
}

type RepoDependencies struct {
	Config *configs.RepositoryConfigs
}

func NewRepository(deps RepoDependencies) Repository {
	repo := new(repository)
	repo.RepoDependencies = deps

	return repo
}
