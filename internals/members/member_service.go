package members

type memberService struct {
	Repository
}

func NewMemberService(repo Repository) *memberService {
	return &memberService{repo}
}

func (s *memberService) CreateMember(memberReq MemberRequest, userID int) (Members, error) {
	newMember := Members{
		UserID: userID,
		Name:   memberReq.Name,
	}

	member, err := s.Repository.CreateMember(newMember)
	if err != nil {
		return Members{}, err
	}

	return member, nil
}
