package friend

type FriendService interface {
	addFriendRequest(friend Friend) (bool, error)
}

type Service struct {
	Repository FriendRepository
}

type Friend struct {
	ID          int    `json:"id"`
	SenderId    int    `json:"senderId"`
	RecipientId int    `json:"recipientId"`
	Community   string `json:"community"`
	Status      string `json:"status"`
}

type ResponseFriendRespons struct {
	Success bool  `json:"success"`
	Message error `json:"message"`
}

func NewFriendsService(repo FriendRepository) FriendService {
	return &Service{Repository: repo}
}

func (s *Service) addFriendRequest(friend Friend) (bool, error) {
	return s.Repository.addFriendRequest(friend)
}
