package group

type groupService struct {
	Repository
}

func NewGroupService(repository Repository) Service {
	return &groupService{repository}
}

func (s *groupService) CreateGroup(userId int, groupRequest GroupRequest) (Group, error) {
	newGroup := Group{
		Name:           groupRequest.Name,
		Description:    groupRequest.Description,
		BalanceTarget:  groupRequest.BalanceTarget,
		DueDate:        groupRequest.DueDate,
		MemberCapacity: groupRequest.MemberCapacity,
		AdminCapacity:  groupRequest.AdminCapacity,
	}

	g, err := s.Repository.CreateGroup(newGroup)
	if err != nil {
		return Group{}, err
	}

	userGroupLink := UserGroupLinks{
		UserID:  userId,
		GroupID: int(g.ID),
		Role:    "superadmin",
	}
	_, err = s.Repository.CreateGroupLink(userGroupLink)
	if err != nil {
		return Group{}, err
	}

	return g, nil
}

func (s *groupService) UpdateGroup(id int, groupReq GroupRequest) (Group, error) {
	g, err := s.Repository.FindById(id)
	if err != nil {
		return Group{}, err
	}

	if groupReq.Name == "" {
		groupReq.Name = g.Name
	}
	if groupReq.Description == "" {
		groupReq.Description = g.Description
	}
	if groupReq.BalanceTarget == 0 {
		groupReq.BalanceTarget = g.BalanceTarget
	}
	if groupReq.DueDate.IsZero() {
		groupReq.DueDate = g.DueDate
	}
	if groupReq.MemberCapacity == 0 {
		groupReq.MemberCapacity = g.MemberCapacity
	}
	if groupReq.AdminCapacity == 0 {
		groupReq.AdminCapacity = g.AdminCapacity
	}

	g.Name = groupReq.Name
	g.Description = groupReq.Description
	g.BalanceTarget = groupReq.BalanceTarget
	g.DueDate = groupReq.DueDate
	g.MemberCapacity = groupReq.MemberCapacity
	g.AdminCapacity = groupReq.AdminCapacity

	updatedGroup, err := s.Repository.UpdateGroup(g)
	if err != nil {
		return Group{}, err
	}

	return updatedGroup, nil

}

func (s *groupService) FindById(id int) (Group, error) {
	group, err := s.Repository.FindById(id)
	if err != nil {
		return Group{}, err
	}
	return group, nil
}

func (s *groupService) DeleteGroup(id int) (Group, error) {
	group, err := s.Repository.FindById(id)

	if err != nil {
		return Group{}, err
	}

	userGroup, err := s.Repository.FindLinkByGroupId(int(group.ID))
	if err != nil {
		return Group{}, err
	}

	deletedGroup, err := s.Repository.DeleteGroup(group)
	if err != nil {
		return Group{}, err
	}

	_, err = s.Repository.DeleteGroupLink(userGroup)

	return deletedGroup, err
}
