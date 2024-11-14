package data

var _ Repository = (*repository)(nil)

type repository struct {
	RepoDependencies
}

// CreateChore implements Repository.
func (r *repository) CreateChore(details ChoreDetails, schedule Schedule) (*ChoreDefinition, error) {
	panic("unimplemented")
}

// CreateMember implements Repository.
func (r *repository) CreateMember(member *Member) (*Member, error) {
	panic("unimplemented")
}

// DeactivateChore implements Repository.
func (r *repository) DeactivateChore(choreID string) error {
	panic("unimplemented")
}

// DeactivateMember implements Repository.
func (r *repository) DeactivateMember(memberID string) error {
	panic("unimplemented")
}

// GetChoreById implements Repository.
func (r *repository) GetChoreById(choreID string) (*ChoreDefinition, error) {
	panic("unimplemented")
}

// GetMemberById implements Repository.
func (r *repository) GetMemberById(memberID string) (*Member, error) {
	panic("unimplemented")
}

// ListChores implements Repository.
func (r *repository) ListChores(pagination PaginationOptions) ([]*ChoreDefinition, error) {
	panic("unimplemented")
}

// ListMembers implements Repository.
func (r *repository) ListMembers(filter MembersFilterOptions, pagination PaginationOptions) ([]*Member, error) {
	panic("unimplemented")
}

// UpdateChoreDetails implements Repository.
func (r *repository) UpdateChoreDetails(choreID string, details ChoreDetails) (*ChoreDefinition, error) {
	panic("unimplemented")
}

// UpdateChoreSchedule implements Repository.
func (r *repository) UpdateChoreSchedule(choreID string, schedule Schedule) (*ChoreDefinition, error) {
	panic("unimplemented")
}

// UpdateMemberInfo implements Repository.
func (r *repository) UpdateMemberInfo(memberID string, info MemberInfo) (*Member, error) {
	panic("unimplemented")
}
